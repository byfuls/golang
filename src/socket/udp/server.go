package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	ip := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])

	addr := net.UDPAddr{Port: port, IP: net.ParseIP(ip)}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	buf := make([]byte, 2048)

	for {
		fmt.Println("Accept a new packet")

		rsize, client, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("net.ReadFromUDP() error: %s\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("Read %d bytes from socket\n", rsize)
			fmt.Printf("Bytes: %q\n", string(buf[:rsize]))
		}
		fmt.Printf("Remote address: %v\n", client)

		wsize, err := conn.WriteToUDP(buf[0:rsize], client)
		if err != nil {
			fmt.Printf("net.WriteTo() error: %s\n", err)
		} else {
			fmt.Printf("Wrote %d bytes to socket\n", wsize)
		}
	}
}
