package main

import "fmt"

func main() {
	var a int
	a = 1
	Increase(&a)

	fmt.Println(a)
	fmt.Println(&a)
}

func Increase(x *int) {
	(*x)++
}
