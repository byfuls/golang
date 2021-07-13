package main

import "fmt"

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 5; i++ {
		var back int
		//a, back = RemoveBack(a)
		a, back = RemoveFront(a)
		fmt.Printf("%d, ", back)
	}

	fmt.Println()
	fmt.Println(a)
}
