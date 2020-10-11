package main

import "fmt"

func main() {
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}

	b = append(b, 3)

	fmt.Printf("%p %p\n", a, b)
}
