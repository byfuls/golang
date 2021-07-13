package main

import (
	"fmt"
	"net"
	//"bufio"
	"os"
)

var ADDRESS = "127.0.0.1:1234"

func echo(client net.Conn) {
	// --1
	//buf := bufio.NewReader(client)
	//for {
	//	rbuf, err := buf.ReadBytes(byte('\n'))
	//	if err != nil {
	//		fmt.Println("read err: ", err)
	//		return
	//	}
	//	fmt.Println(rbuf)
	//	client.Write(rbuf)
	//}

	// --2
	buf := make([]byte, 1024)
	for {
		size, err := client.Read(buf)
		if err != nil || size == 0 {
			fmt.Println("read err: ", err)
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
