package main

import "fmt"

func main() {
	s := "abcdefg"
	var a [20]byte

	copy(a[:], s)
	fmt.Println("s: ", []byte(s), "a: ", a)

	var b [2]byte

	copy(b[:], s)
	fmt.Println("s: ", []byte(s), "b: ", b)
}
