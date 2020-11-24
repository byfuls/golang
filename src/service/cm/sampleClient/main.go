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
	PORT       = 10100
	IP         = "127.0.0.1"
	REGIST_KEY = "51010"
)

func main() {
	addr := net.TCPAddr{Port: PORT, IP: net.ParseIP(IP)}
	connCM, err := net.DialTCP("tcp", nil, &addr)
	if err != nil {
		fmt.Println("[sampleClient] make bridge socket error: ", err)
		os.Exit(1)
	} else {

		go func() {
			rbuf := make([]byte, 1024)
			for {
				rsize, err := connCM.Read(rbuf)
				if err != nil || 0 >= rsize {
					fmt.Println("[sampleClient] proxy socket disconnected: ", err)
					return
				} else {
					fmt.Printf("[sampleClient] receive buf: \n%s\n", hex.Dump(rbuf[:rsize]))

					p := protocol.Packet{}
					rbuf, len := p.Make("SMS", "RESULT", 0, uint8(len(REGIST_KEY)), []byte(REGIST_KEY), 2, []byte("OK"))
					if 0 == len {
						fmt.Println("[sampleClient-response] make regist protocol error")
						os.Exit(1)
					}
					if wsize, err := connCM.Write(rbuf); err != nil {
						fmt.Println("[sampleClient-response] write to response to CM error: ", err)
						return
					} else {
						fmt.Println("[sampleClient-response] write to response to CM success: ", wsize)
					}
				}
			}
		}()

		p := protocol.Packet{}
		wbuf := make([]byte, 1024)
		wbuf, len := p.Make("CONNECTION", "REGIST", 0, uint8(len(REGIST_KEY)), []byte(REGIST_KEY), 0, nil)
		if 0 == len {
			fmt.Println("[registCM] make regist protocol error")
			os.Exit(1)
		}
		if wsize, err := connCM.Write(wbuf); err != nil {
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
