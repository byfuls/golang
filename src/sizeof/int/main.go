package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	var b int32

	a = 5
	b = 3

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}
