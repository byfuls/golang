package main

import (
	"bytes"
	"os"
	"fmt"
	"encoding/binary"
	"encoding/hex"
)

func main() {
	b := new(bytes.Buffer)

	b.Write([]byte("HI"))
	b.Write([]byte("123"))
	if err := binary.Write(b, binary.BigEndian, byte(123)); err != nil {
		fmt.Println("error: ", err)
	}

	var test uint32
	test = 321
	if err := binary.Write(b, binary.BigEndian, test); err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("%s\n", hex.Dump(b.Bytes()))
	fmt.Printf("%s\n", b.Bytes())
	b.WriteTo(os.Stdout)
}
