package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))

	x := bytes.Index([]byte("abcdefg"), []byte("c"))
	fmt.Println(x)
}
