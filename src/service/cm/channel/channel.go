package channel

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"

	"service/cm/handling"

	"service/proxy/channelManager"
)

func channelReceiver(chToDv chan handling.Message, ch *net.TCPConn, key string) {
	fmt.Printf("[channelReceiver] start addr: %v [%s]\n", ch, key)

	buf := make([]byte, 1024)
	for {
		rsize, err := ch.Read(buf)
		if err != nil || 0 >= rsize {
			if ret := channelManager.Del(key); ret {
				fmt.Println("[channelReceiver] receive error: ", err, "delete channel success")
				return
			} else {
				fmt.Println("[channelReceiver] receive error: ", err, "delete channel fail")
				os.Exit(1)
			}
		} else {
			fmt.Printf("[channelReceiver] receive buf: \n%s\n", hex.Dump(buf[:rsize]))
			chToDv <- handling.Message{
				From: handling.CH,
				Buf:  buf[:rsize],
			}
		}
	}
}

func channelWriter(CHMgr *handling.CHManage) {
	fmt.Println("[channelWriter] start")

	for {
		if message := <-CHMgr.ToCH; CHMgr.Socket != nil {
			fmt.Println("[channelWriter] ", CHMgr)

			if wsize, err := CHMgr.Socket.Write(message.Buf); err != nil {
				fmt.Println("[channelWriter] write to CH error: ", err)
			} else {
				fmt.Println("[channelWriter] write to CH success: ", wsize)
			}
		} else {
			fmt.Println("[channelWriter] could not use socket")
		}
	}
}

func Accepter(chToDv chan handling.Message, ip string, port int) {
	fmt.Println("[accepter] start, server addr: ", ip, port)

	addr := net.TCPAddr{Port: port, IP: net.ParseIP(ip)}
	conn, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Println("[accepter] tcp listen error: ", err)
		os.Exit(1)
	}

	for {
		ch, err := conn.AcceptTCP()
		if err != nil {
			fmt.Println("[accepter] accept error: ", err)
			continue
		}
		channel := handling.CHManage{
			Socket: ch,
			ToCH:   make(chan handling.Message),
		}

		TEST_KEY := "TEST"
		if ret := channelManager.Put(TEST_KEY, channel); ret {
			go channelReceiver(chToDv, ch, TEST_KEY)
			go channelWriter(&channel)
		} else {
			fmt.Println("[accepter] channel manager put error")
		}
	}
}
