package main

import (
	"fmt"
)

// [PANIC]
//func test() int {
//	arr := []int{}
//	element := arr[5]
//	return element
//}
//
//func main() {
//	test()
//	fmt.Println("smooth exit")
//}

func r() bool {
	if r := recover(); r != nil {
		fmt.Println("Recovered!!", r)
		//debug.PrintStack()
	}
	return false
}

func a() bool {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
	return true
}

func main() {
	fmt.Println(a())
	fmt.Println("normally returned from main")
}
