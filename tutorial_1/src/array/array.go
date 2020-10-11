package main

import "fmt"

func main() {
	// --1
	var a [10]int
	fmt.Println(a)
	fmt.Println(len(a))

	// --2
	b := [10]int{1, 2, 3}
	fmt.Println(b)
	fmt.Println(len(b))

	// --3
	c := [10]int{}
	fmt.Println(c)
	fmt.Println(len(c))
}
