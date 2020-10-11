package main

import "fmt"

func main() {
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := append(a, 3)

	fmt.Printf("%p %p\n", a, b)

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 4
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}
