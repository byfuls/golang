package main

import (
	"fmt"
	"encoding/binary"
	"encoding/hex"
	"bytes"
)

func main() {

	buf := new(bytes.Buffer)

	protocols := [][]byte{
		{0x01},
		{0x02},
		{0x03},
		[]byte("HIHI"),
	}	

	for _, b := range protocols {
		if err := binary.Write(buf, binary.BigEndian, b); err != nil {
			fmt.Println("write error")
		}
	}

	fmt.Printf("%s\n", hex.Dump(buf.Bytes()))
}
