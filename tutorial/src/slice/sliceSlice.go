package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8]
	fmt.Println(b)

	c := a[4:]
	fmt.Println(c)

	d := a[:4]
	fmt.Println(d)

	/********************************************/
	/* THE SAME MEMORY POINTER */
	/********************************************/

	b[0] = 1
	b[1] = 2
	fmt.Println(a)
}
