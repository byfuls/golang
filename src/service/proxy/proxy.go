package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"service/proxy/channelManager"
)

const (
	DEFAULT_UDP_IP   = "127.0.0.1"
	DEFAULT_UDP_PORT = "10000"
	DEFAULT_TCP_IP   = "127.0.0.1"
	DEFAULT_TCP_PORT = "10001"
	HANDLER_COUNT    = 1
)

const (
	CP = iota
	CM = iota
)

type bypass struct {
	from int
	buf  []byte
}

type CMManage struct {
	socket *net.TCPConn
	toCM   chan bypass
}

var g_conn *net.UDPConn

func writeToCM(CMMgr *CMManage) {
	fmt.Println("[writeToCM] start")

	for {
		if bypass := <-CMMgr.toCM; CMMgr.socket != nil {
			fmt.Println("[writeToCM] ", CMMgr)

			if wsize, err := CMMgr.socket.Write(bypass.buf); err != nil {
				fmt.Println("[writeToCM] write to CM error: ", err)
			} else {
				fmt.Println("[writeToCM] write to CM success: ", wsize)
			}
		} else {
			fmt.Println("[writeToCM] could not use socket")
		}
	}
}

func receiveFromCM(rToh chan bypass, cm *net.TCPConn, key string) {
	fmt.Printf("[receiveFromCM] start addr: %v\n", cm)

	buf := make([]byte, 1024)
	for {
		rsize, err := cm.Read(buf)
		if err != nil || 0 >= rsize {
			if ret := channelManager.Del(key); ret {
				fmt.Println("[receiveFromCM] receive error: ", err, "delete channel success")
				return
			} else {
				fmt.Println("[receiveFromCM] receive error: ", err, "delete channel fail")
				os.Exit(1)
			}
		} else {
			fmt.Printf("[receiveFromCM] receive buf: %s\n", hex.Dump(buf[:rsize]))
			rToh <- bypass{
				from: CM,
				buf:  buf[:rsize],
			}
		}
	}
}

func acceptFromCM(rToh chan bypass, ip string, port int) {
	fmt.Println("[acceptFromCM] start")

	addr := net.TCPAddr{Port: port, IP: net.ParseIP(ip)}
	conn, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Println("[acceptFromCM] tcp listen error: ", err)
		os.Exit(1)
	}

	for {
		cm, err := conn.AcceptTCP()
		if err != nil {
			fmt.Println("[acceptFromCM] accept error: ", err)
		}

		channel := CMManage{
			socket: cm,
			toCM:   make(chan bypass),
		}

		fmt.Println(channel)
		TEST_KEY := "TEST"
		if ret := channelManager.Put(TEST_KEY, channel); ret {
			go receiveFromCM(rToh, cm, TEST_KEY)
			go writeToCM(&channel)
		} else {
			fmt.Println("[acceptFromCM] channel manager put error")
		}
	}
}

func handler(rToh chan bypass, hTow chan bypass, no int) {
	fmt.Printf("[handler] %d ready\n", no)

	for {
		packet := <-rToh
		switch packet.from {
		case CP:
			fmt.Println("[handler] (CP) receive data")
			fmt.Printf("[handler] receive channel: %s\n", hex.Dump(packet.buf))

			if tmp, ret := channelManager.Get("TEST"); ret {
				cm := tmp.(CMManage)
				cm.toCM <- bypass{
					from: CP,
					buf:  packet.buf,
				}
			} else {
				fmt.Println("[handler] (CP) not found channel in map")
				continue
			}
		case CM:
			fmt.Println("[handler] (CM) receive data")
			fmt.Printf("[handler] receive channel: %s\n", hex.Dump(packet.buf))
		}
	}
}

func receiveFromCP(rToh chan bypass, ip string, port int) {
	fmt.Println("[receiveFromCP] start addr: ", ip, port)

	addr := net.UDPAddr{Port: port, IP: net.ParseIP(ip)}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("[receiveFromCP] udp listen error: ", err)
		os.Exit(1)
	} else {
		g_conn = conn
	}
	buf := make([]byte, 1024)

	for {
		fmt.Println("[receiveFromCP] Accept a new packet")

		rsize, client, err := g_conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("[receiveFromCP] read error: ", err)
			os.Exit(1)
		} else {
			fmt.Printf("[receiveFromCP] client address: %v\n", client)
			fmt.Printf("[receiveFromCP] %s\n", hex.Dump(buf[:rsize]))

			rToh <- bypass{
				from: CP,
				buf:  buf[:rsize],
			}
		}
	}
}

func main() {
	fmt.Println("[proxy] start ...")

	var udp_ip string
	var udp_port int
	var tcp_ip string
	var tcp_port int
	if 1 >= len(os.Args) {
		udp_ip = DEFAULT_UDP_IP
		udp_port, _ = strconv.Atoi(DEFAULT_UDP_PORT)
		tcp_ip = DEFAULT_TCP_IP
		tcp_port, _ = strconv.Atoi(DEFAULT_TCP_PORT)
	} else {
		udp_ip = os.Args[1]
		udp_port, _ = strconv.Atoi(os.Args[2])
	}
	fmt.Println(udp_ip, udp_port, tcp_ip, tcp_port)

	channelManager.Init()

	rToh := make(chan bypass)
	hTow := make(chan bypass)

	go receiveFromCP(rToh, udp_ip, udp_port)
	//go writeToCP()

	go acceptFromCM(rToh, tcp_ip, tcp_port)
	for i := 0; i < HANDLER_COUNT; i++ {
		go handler(rToh, hTow, i)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}
