package main

import (
	"fmt"
	"net"
	//"bufio"
	"os"
	"time"
	"log"
)

var ADDRESS = "127.0.0.1:1234"
const TIMEOUT = 10

func echo(client net.Conn) {
	defer func() {
		client.Close()
	}()

	log.Println("start")

	buf := make([]byte, 1024)

	timeoutDuration := time.Duration(TIMEOUT) * time.Second
	client.SetDeadline(time.Now().Add(timeoutDuration))

	for {
		size, err := client.Read(buf)
		if err != nil || size == 0 {
			log.Println("read error: ", err)
			return
		}
		//fmt.Println(buf)
		client.Write(buf)
	}
}

func main() {
	sock, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		fmt.Println("server listen err: ", err)
		os.Exit(1)
	}

	for {
		client, err := sock.Accept()
		if err != nil {
			fmt.Println("server accept err: ", err)
			continue
		}

		go echo(client)
	}
}
