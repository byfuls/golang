package main

import (
	"fmt"
)

const (
	E0 = iota
	E1 = iota
	E2 = iota
)

func main() {
	fmt.Println(E0, E1, E2)
	fmt.Printf("%T\n", E0) // show data type
}
