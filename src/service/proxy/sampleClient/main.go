package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"time"

	"service/protocol"
)

const (
	PORT = 10000
	IP   = "127.0.0.1"

	CP_PORT = 10010
	CP_IP = "127.0.0.1"
)

func main() {
	addr := net.UDPAddr{Port: PORT, IP: net.ParseIP(IP)}
	connCP, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		fmt.Println("[sampleClient] make bridge socket error: ", err)
		os.Exit(1)
	} else {

		// cp main receiver
		go func() {
			addr := net.UDPAddr{Port: CP_PORT, IP: net.ParseIP(CP_IP)}
			conn, err := net.ListenUDP("udp", &addr)
			if err != nil {
				fmt.Println("[CP main receiver] udp listen error: ", err)
				os.Exit(1)
			}

			rbuf := make([]byte, 1024)
			for {
				rsize, err := conn.Read(rbuf)
				if err != nil || 0 >= rsize {
					fmt.Println("[sampleClient] proxy socket disconnected: ", err)
					return
				} else {
					fmt.Printf("[sampleClient] receive buf: \n%s\n", hex.Dump(rbuf[:rsize]))
				}
			}
		}()

		// receiver 
		go func() {
			rbuf := make([]byte, 1024)
			for {
				rsize, err := connCP.Read(rbuf)
				if err != nil || 0 >= rsize {
					fmt.Println("[sampleClient] proxy socket disconnected: ", err)
					return
				} else {
					fmt.Printf("[sampleClient] receive buf: \n%s\n", hex.Dump(rbuf[:rsize]))
				}
			}
		}()

		// writer
		p := protocol.Packet{}
		REGIST_KEY := "51010"
		wbuf := make([]byte, 1024)
		wbuf, len := p.Make("SMS", "SEND", 0, uint8(len(REGIST_KEY)), []byte(REGIST_KEY), 5, []byte("HELLO"))
		if 0 == len {
			fmt.Println("[registCM] make regist protocol error")
			os.Exit(1)
		}
		if wsize, err := connCP.Write(wbuf); err != nil {
			fmt.Println("[registCM] write to Regist CM error: ", err)
			return
		} else {
			fmt.Println("[registCM] write to Regist CM success: ", wsize)
		}

		for {
			time.Sleep(1 * time.Second)
		}
	}
}
