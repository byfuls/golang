package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	arr := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	size := len(arr)
	p := uintptr(unsafe.Pointer(&arr))

	var data []byte

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = p + 3
	sh.Len = size - 3
	sh.Cap = size

	fmt.Println(data)
	fmt.Println(arr)

	runtime.KeepAlive(arr)

	p1 := unsafe.Pointer(&arr)
	fmt.Println((*[10]byte)(p1)[1:])
}
