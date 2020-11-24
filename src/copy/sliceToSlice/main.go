package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	a := make([]byte, 4)
	b := make([]byte, 8)

	a = []byte("abcd")
	b = []byte("12345678")

	fmt.Printf("[a] \n%s\n", hex.Dump(a))
	fmt.Printf("[b] \n%s\n", hex.Dump(b))

	copy(a, b)

	fmt.Printf("[a] \n%s\n", hex.Dump(a))
	fmt.Printf("[b] \n%s\n", hex.Dump(b))
}
