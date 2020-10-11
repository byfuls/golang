package main

import "fmt"

func main() {
	// method 1.
	a := 4
	b := 2

	//// method 2.
	//var c int
	//c = 5

	//// method 3.
	//var d = 6
	//var e = 3.14

	//// method 4.
	//var f int = 8

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Println("result = ", a^b)

	c := 21
	v := c / 10
	l := c % 10
	fmt.Printf("10: %v 1: %v\n", v, l)
}
