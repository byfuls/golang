package main

import "fmt"

func main() {
	var a []int
	b := []int{1, 2, 3, 4}
	c := make([]int, 3)
	d := make([]int, 0, 8)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
