package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
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

func receiveFromCM(rToh chan bypass, cm *net.TCPConn) {
	fmt.Printf("[receiveFromCM] start addr: %v\n", cm)

	buf := make([]byte, 1024)
	for {
		rsize, err := cm.Read(buf)
		if err != nil || 0 >= rsize {
			fmt.Println("[receiveFromCM] receive error: ", err)
			return
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

		go receiveFromCM(rToh, cm)
		go writeToCM(&channel)
	}
}

func handler(rToh chan bypass, hTow chan bypass, no int) {
	fmt.Printf("[handler] %d ready\n", no)

	for {
		bypass := <-rToh

		fmt.Printf("[handler] receive channel: %s\n", hex.Dump(bypass.buf))
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
