package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"module/logging"
	"service/sim/clientHandling"
	"service/sim/simProtHandling"
)

type clientPacket struct {
	addr *net.UDPAddr
	buf  []byte
}

var g_conn *net.UDPConn

const (
	PORT_UDP_CP = 20000
)

func writer(hTow chan clientPacket) {
	logging.TraceLn("writer start")

	for {
		clientPacket := <-hTow

		logging.DebugLn(clientPacket)
		logging.DebugF("[writer] Client address: %v\n", clientPacket.addr)
		logging.DebugF("[writer] Received bytes [%s] from socket\n", string(clientPacket.buf))
		logging.DebugF("%s\n", hex.Dump(clientPacket.buf))

		if wsize, err := g_conn.WriteTo(clientPacket.buf, clientPacket.addr); err != nil {
			logging.ErrorLn("[writer] write error: ", err)
		} else {
			logging.TraceLn("[writer] write success ", wsize)
		}
	}
}

func handler(rToh chan clientPacket, hTow chan clientPacket, no int) {
	logging.DebugF("handler %d ready\n", no)

	for {
		client := <-rToh

		logging.DebugLn(client)
		logging.DebugF("[handler] Client address: %v\n", client.addr)
		logging.DebugF("[handler] Received bytes [%s] from socket\n", string(client.buf))

		ret, pdata := simProtHandling.Parsing(client.buf)
		pdata.Addr = *client.addr
		if ret == false {
			logging.ErrorLn("[handler] parsing err")
			continue
		}

		switch pdata.Head.Command {
		case "AS90": // from Device
			if err := simProtHandling.Recv_AS90(&pdata); err == false {
				logging.ErrorLn("[handler] Recv_AS90 error")
				continue
			}

			if buf, len := simProtHandling.Resp_AA00(&pdata); 0 >= len {
				logging.ErrorLn("[handler] Resp_AA00 error")
				continue
			} else {
				hTow <- clientPacket{
					addr: &pdata.Addr,
					buf:  buf,
				}
			}
		case "AS03": // from Device
			if err := simProtHandling.Recv_AS03(&pdata); err == false {
				logging.ErrorLn("[handler] Recv_AS03 error")
				continue
			}

			if buf, len := simProtHandling.Resp_AA00(&pdata); 0 >= len {
				logging.ErrorLn("[handler] Resp_AA00 error")
				continue
			} else {
				hTow <- clientPacket{
					addr: &pdata.Addr,
					buf:  buf,
				}
			}

			if buf, len := simProtHandling.Send_AS03(&pdata); 0 >= len {
				logging.ErrorLn("[handler] Send_AS03 error")
				continue
			} else {
				hTow <- clientPacket{
					addr: &pdata.Addr,
					buf:  buf,
				}
			}
		case "AA00": // from Device & from Gateway
			if err := simProtHandling.Recv_AA00(&pdata); err == false {
				logging.ErrorLn("[handler] Recv_AA00 error")
				continue
			}
		case "AC03": // from Gateway
			if err := simProtHandling.Recv_AC03(&pdata); err == false {
				logging.ErrorLn("[handler] Recv_AC03 error")
				continue
			}

			if buf, len := simProtHandling.Resp_AA00(&pdata); 0 >= len {
				logging.ErrorLn("[handler] Resp_AA00 error")
				continue
			} else {
				hTow <- clientPacket{
					addr: &pdata.Addr,
					buf:  buf,
				}
			}

			if buf, len := simProtHandling.Send_AC03(&pdata); 0 >= len {
				logging.ErrorLn("[handler] Send_AC03 error")
				continue
			} else {
				hTow <- clientPacket{
					addr: client.addr,
					buf:  buf,
				}
			}
		case "MA06": // from MP
			if err := simProtHandling.Recv_MA06(&pdata); err == false {
				logging.ErrorLn("[handler] Recv_MA06 error")
				continue
			}
		case "AS07": // from Device
			if err := simProtHandling.Recv_AS07(&pdata); err == false {
				logging.ErrorLn("[handler] Recv_AS07 error")
				continue
			}
			if buf, len := simProtHandling.Send_LUR(&pdata); 0 >= len {
				logging.ErrorLn("[handler] Send_LUR error")
				continue
			} else {
				addr := net.UDPAddr{Port: PORT_UDP_CP, IP: net.ParseIP("127.0.0.1")}
				hTow <- clientPacket{
					addr: &addr,
					buf:  buf,
				}
			}
			if buf, len := simProtHandling.Resp_AS07(&pdata); 0 >= len {
				logging.ErrorLn("[handler] Resp_AS07 error")
				continue
			} else {
				hTow <- clientPacket{
					addr: client.addr,
					buf:  buf,
				}
			}
		default:
			logging.ErrorLn("[handler] unknown command: ", pdata.Head.Command)
		}
	}
}

func receiver(rToh chan clientPacket, ip string, port int) {
	logging.DebugLn("receiver start addr: ", ip, port)

	addr := net.UDPAddr{Port: port, IP: net.ParseIP(ip)}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		logging.ErrorLn("udp listen error: ", err)
		os.Exit(1)
	} else {
		g_conn = conn
	}
	buf := make([]byte, 1024)

	for {
		logging.DebugLn("Accept a new packet")

		rsize, client, err := g_conn.ReadFromUDP(buf)
		if err != nil {
			logging.ErrorF("net.ReadFromUDP() error: %s\n", err)
			os.Exit(1)
		} else {
			logging.DebugF("Client address: %v\n", client)

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

	if !logging.Init("/home/byfuls/golang/log", "sim") {
		fmt.Println("logging init error")
		return
	}
	logging.DebugLn("sim start")

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
