package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

func main() {
	receivedPacket := new(bytes.Buffer)

	protocols := [][]byte{
		[]byte("MA06"),
		[]byte("510101182682080"),
		{0x00, 0x00, 0x00, 0x00, 0x00},
		{0x00, 0x00, 0x00, 0x00},
		{0x00, 0x00, 0x00, 0x00},
		[]byte("|"),
		[]byte("510101182682080"),
		[]byte("|"),
		{0x00, 0x00, 0x00, 0x00},
		[]byte("|"),
		[]byte("gatewayId"),
		[]byte("|"),
	}

	for _, b := range protocols {
		fmt.Println(b)
		if err := binary.Write(receivedPacket, binary.BigEndian, b); err != nil {
			fmt.Println("err")
			return
		}
	}

	fmt.Printf("%s\n", hex.Dump(receivedPacket.Bytes()))

	tmp := receivedPacket.Bytes()
	for found0 := bytes.Index(tmp, []byte("|")); found0 != -1; found0 = bytes.Index(tmp, []byte("|")) {
		fmt.Println(found0)
		fmt.Printf("%s\n", tmp[:found0])
		tmp = tmp[found0+1:]
	}

}
