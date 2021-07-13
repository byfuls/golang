package main

import "fmt"

func main() {
	var a int
	var p *int

	var b int

	p = &a
	a = 3

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	b = 2
	p = &b
	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)
}
