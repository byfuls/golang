package main

import (
	"fmt"
	"bytes"
)

func main() {
	buf1 := []byte{0x31, 0x32, 0x33, 0x00}

	fmt.Println(buf1)
	fmt.Println(bytes.Trim(buf1, "\x00"))
}
