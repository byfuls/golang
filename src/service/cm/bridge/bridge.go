package bridge

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"service/cm/handling"
	"service/protocol"
)

type proxyStatus struct {
	connFlag bool
	proxy    *net.TCPConn
}

/***********************************/
func registCM() ([]byte, uint32) {
	p := protocol.Packet{}
	TEST_KEY := "51010" // TEST TEST TEST
	buf, len := p.Make("CONNECTION", "REGIST", 0, uint8(len(TEST_KEY)), []byte(TEST_KEY), 0, nil)
	if 0 == len {
		fmt.Println("[registCM] make regist protocol error")
		return nil, 0
	}
	return buf, len
}

func proxyReceiver(pxy *proxyStatus, pxToDv chan handling.Message) {
	fmt.Println("[proxyReceiver] start")

	buf := make([]byte, 1024)
	for {
		rsize, err := pxy.proxy.Read(buf)
		if err != nil || 0 >= rsize {
			fmt.Println("[proxyReceiver] proxy socket disconnected: ", err, "back to connect")
			pxy.connFlag = false
			pxy.proxy = nil
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

func proxyWriter(pxy *proxyStatus, dvToPx chan handling.Message) {
	fmt.Println("[proxyWriter] start")

	if buf, _ := registCM(); buf != nil {
		if wsize, err := pxy.proxy.Write(buf); err != nil {
			fmt.Println("[proxyWriter] write to Regist Proxy error: ", err, "back to connect")
			pxy.connFlag = false
			pxy.proxy = nil
			return
		} else {
			fmt.Println("[proxyWriter] write to Regist Proxy success: ", wsize)
		}
	}

	for {
		message := <-dvToPx
		if wsize, err := pxy.proxy.Write(message.Buf); err != nil {
			fmt.Println("[proxyWriter] write to Proxy error: ", err)
			pxy.connFlag = false
			pxy.proxy = nil
			return
		} else {
			fmt.Println("[proxyWriter] write to Proxy success: ", wsize)
		}
	}
}

func Bridge(pxToDv chan handling.Message, dvToPx chan handling.Message, ip string, port int) {
	fmt.Println("[bridge] start, remote addr: ", ip, port)

	pxy := proxyStatus{}

	for {
		if !pxy.connFlag {
			addr := net.TCPAddr{Port: port, IP: net.ParseIP(ip)}
			proxy, err := net.DialTCP("tcp", nil, &addr)
			if err != nil {
				fmt.Println("[bridge] make bridge socket error: ", err, " retry ...")
				//os.Exit(1)
			} else {
				pxy.connFlag = true
				pxy.proxy = proxy

				go proxyReceiver(&pxy, pxToDv)
				go proxyWriter(&pxy, dvToPx)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
