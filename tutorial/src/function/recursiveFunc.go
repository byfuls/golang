package main

import "fmt"

func main() {
	s := sum(10, 0)
	fmt.Println(s)
}

func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}
