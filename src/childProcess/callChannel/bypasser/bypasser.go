package bypasser

//package main

import (
	"fmt"
	"net"
	"time"
)

type bypassServer struct {
	udpAddr   net.UDPAddr
	udpSock   *net.UDPConn
	maxBufLen int

	guestsCount int
	lastInTime  time.Time

	timeout chan bool
}

var server *bypassServer

func showGuests(guests [2]*net.UDPAddr) {
	//fmt.Println("-----------------------------------")
	//for i := 0; i < 2; i++ {
	//	fmt.Printf("(%d) [%s]\n", i, guests[i].String())
	//}
	//fmt.Println("-----------------------------------")
}

func (by *bypassServer) operator() {
	var guests [2]*net.UDPAddr
	by.guestsCount = 0
	for {
		buf := make([]byte, by.maxBufLen)

		_, guest, err := by.udpSock.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("[operator] read error: ", err)
			return
		} else {
			//fmt.Printf("[operator] read[%s] len(%d) \n%s\n", guest.String(), rsize, hex.Dump(buf))

			by.lastInTime = time.Now()
			if 2 > by.guestsCount {
				//fmt.Println("guest[0]: ", len(guests[0].String()))
				if guests[0] == nil {
					guests[0] = guest
					showGuests(guests)
				} else if guests[1] == nil {
					guests[1] = guest
					showGuests(guests)

					_, err := by.udpSock.WriteTo(buf, guests[0])
					if err != nil {
						fmt.Printf("[operator] write [%s => %s] error: %s\n",
							guest.String(), guests[1].String(), err)
						return
					} else {
						//fmt.Printf("[operator] write [%s => %s] success: %d",
						//	guest.String(), guests[1].String(), wsize)
					}
				}
				by.guestsCount++
			} else {
				if guests[0].String() == guest.String() {
					_, err := by.udpSock.WriteTo(buf, guests[1])
					if err != nil {
						fmt.Printf("[operator] write [%s => %s] error: %s\n",
							guest.String(), guests[1].String(), err)
						return
					} else {
						//fmt.Printf("[operator] write [%s => %s] success: %d\n",
						//	guest.String(), guests[1].String(), wsize)
					}
				} else if guests[1].String() == guest.String() {
					_, err := by.udpSock.WriteTo(buf, guests[0])
					if err != nil {
						fmt.Printf("[operator] write [%s => %s] error: %s\n",
							guest.String(), guests[1].String(), err)
						return
					} else {
						//fmt.Printf("[operator] write [%s => %s] success: %d\n",
						//	guest.String(), guests[1].String(), wsize)
					}
				}
			}
		}
	}
}

func (by *bypassServer) watcher(timeout int64) {
	for {
		if by.guestsCount > 0 {
			now := time.Now()
			diff := int64(now.Sub(by.lastInTime) / time.Second)
			if diff > timeout {
				fmt.Printf("[watcher] timeout(%d)\n", diff)
				by.timeout <- true
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func Timeout() bool {
	timeout := <-server.timeout
	if timeout {
		return true
	}
	return false
}

func Operator(timeout int64) {
	go server.operator()
	go server.watcher(timeout)
}

func Done() {
	server.udpSock.Close()
	server.udpSock = nil
	server.timeout = nil
	server = nil
}

func Init(ip string, port int, maxBufLen int) error {
	server = new(bypassServer)
	server.udpAddr = net.UDPAddr{Port: port, IP: net.ParseIP(ip)}

	var err error
	server.udpSock, err = net.ListenUDP("udp", &server.udpAddr)
	if err != nil {
		fmt.Println("[Init] error: ", err)
		return err
	}
	server.maxBufLen = maxBufLen
	server.timeout = make(chan bool)

	return nil
}
