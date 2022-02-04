package main

import (
	"fmt"
)

func a() []string {
	fmt.Println("a-1")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("a, recover")
		}
	}()
	fmt.Println("a-2, before call b()")
	return b()
	// fmt.Println("a-3, after call b()")
	// return rslt
}

func b() []string {
	// rslt := []string{"a1", "b2", "c3"}
	rslt := []string{"a"}
	return rslt

	fmt.Println("b-1")
	panic("b panic")
	fmt.Println("b-2")

	// rslt = "normal"
	return rslt
}

func main() {
	fmt.Println("main start")
	rslt := a()
	fmt.Printf("main end, result:%v,%v\n", len(rslt), rslt)
}
