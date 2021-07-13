package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	b := make([]int, 0, 8)
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))

	b = append(b, 1)
	fmt.Println(b)
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))
}
