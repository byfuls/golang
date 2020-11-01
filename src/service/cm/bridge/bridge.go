package bridge

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"time"

	"service/cm/handling"
	"service/protocol"
)

//func regist() {
//	/* TODO TODO TODO */
//	fmt.Println("[regist] start")
//
//	return []byte("[regist] TEST")
//}

/***********************************/
func registCM() ([]byte, uint32) {
	p := protocol.Packet{}
	TEST_KEY := "TEST"
	buf, len := p.Make("CONNECTION", "REGIST", 0, uint8(len(TEST_KEY)), []byte(TEST_KEY), 0, nil)
	if 0 == len {
		fmt.Println("[registCM] make regist protocol error")
		return nil, 0
	}
	return buf, len
}

func proxyReceiver(proxy *net.TCPConn, pxToDv chan handling.Message) {
	fmt.Println("[proxyReceiver] start")

	buf := make([]byte, 1024)
	for {
		rsize, err := proxy.Read(buf)
		if err != nil || 0 >= rsize {
			/* TODO TODO TODO */
			fmt.Println("[proxyReceiver] proxy socket disconnected: ", err)
			return
		} else {
			fmt.Printf("[proxyReceiver] receive buf: \n%s\n", hex.Dump(buf[:rsize]))
			pxToDv <- handling.Message{
				From: handling.PROXY,
				Buf:  buf[:rsize],
			}
		}
	}
}

func proxyWriter(proxy *net.TCPConn, dvToPx chan handling.Message) {
	fmt.Println("[proxyWriter] start")

	if buf, _ := registCM(); buf != nil {
		if wsize, err := proxy.Write(buf); err != nil {
			fmt.Println("[proxyWriter] write to Regist Proxy error: ", err)
			return
		} else {
			fmt.Println("[proxyWriter] write to Regist Proxy success: ", wsize)
		}
	}

	for {
		message := <-dvToPx
		if wsize, err := proxy.Write(message.Buf); err != nil {
			fmt.Println("[proxyWriter] write to Proxy error: ", err)
			return
		} else {
			fmt.Println("[proxyWriter] write to Proxy success: ", wsize)
		}
	}
}

func Bridge(pxToDv chan handling.Message, dvToPx chan handling.Message, ip string, port int) {
	fmt.Println("[bridge] start, remote addr: ", ip, port)

	addr := net.TCPAddr{Port: port, IP: net.ParseIP(ip)}
	proxy, err := net.DialTCP("tcp", nil, &addr)
	if err != nil {
		fmt.Println("[bridge] make bridge socket error: ", err)
		os.Exit(1)
	}

	go proxyReceiver(proxy, pxToDv)
	go proxyWriter(proxy, dvToPx)

	for {
		time.Sleep(1 * time.Second)
	}
}
