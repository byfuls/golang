package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"service/sim/clientHandling"
	"service/sim/simProtHandling"
)

type clientPacket struct {
	addr *net.UDPAddr
	buf  []byte
}

func writer(hTow chan clientPacket) {
	fmt.Println("writer start")

	for {
		clientPacket := <-hTow

		fmt.Println(clientPacket)
		fmt.Printf("[writer] Client address: %v\n", clientPacket.addr)
		fmt.Printf("[writer] Received bytes [%s] from socket\n", string(clientPacket.buf))
	}
}

func handler(rToh chan clientPacket, hTow chan clientPacket, no int) {
	fmt.Printf("handler %d ready\n", no)

	for {
		client := <-rToh

		fmt.Println(client)
		fmt.Printf("[handler] Client address: %v\n", client.addr)
		fmt.Printf("[handler] Received bytes [%s] from socket\n", string(client.buf))

		ret, pdata := simProtHandling.Parsing(client.buf)
		pdata.Addr = *client.addr
		if ret == false {
			fmt.Printf("[handler] parsing err")
			continue
		}

		switch pdata.Head.Command {
		case "AS07":
			if err := simProtHandling.Recv_AS07(&pdata); err == false {
				fmt.Println("[handler] Recv_AS07 error")
				continue
			}
			if buf, len := simProtHandling.Resp_AS07(&pdata); 0 >= len {
				fmt.Println("[handler] Resp_AS07 error")
				continue
			} else {
				/* TODO TODO TODO */
				hTow <- clientPacket{
					addr: client.addr,
					buf:  buf,
				}
			}
		default:
			fmt.Println("[handler] unknown command: ", pdata.Head.Command)
		}
	}
}

func receiver(rToh chan clientPacket, ip string, port int) {
	fmt.Println("receiver start addr: ", ip, port)

	addr := net.UDPAddr{Port: port, IP: net.ParseIP(ip)}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("udp listen error: ", err)
		os.Exit(1)
	}
	buf := make([]byte, 1024)

	for {
		fmt.Println("Accept a new packet")

		rsize, client, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("net.ReadFromUDP() error: %s\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("Received %s bytes [%s] from socket\n", rsize, string(buf[:rsize]))
			fmt.Printf("Client address: %v\n", client)

			rToh <- clientPacket{
				addr: client,
				buf:  buf[:rsize],
			}
		}
	}
}

func main() {
	ip := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])
	handlerCount := 5

	rToh := make(chan clientPacket)
	hTow := make(chan clientPacket)

	go receiver(rToh, ip, port)
	go writer(hTow)

	clientHandling.Init()
	for i := 0; i < handlerCount; i++ {
		go handler(rToh, hTow, i)
	}

	for {
		time.Sleep(1 * time.Second)
		//fmt.Println("main loop")
	}
}
