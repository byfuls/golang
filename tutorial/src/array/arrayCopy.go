package main

import "fmt"

func main() {
	// --1
	//arr := [5]int{1, 2, 3, 4, 5}
	//clone := [5]int{}

	//for i := 0; i < len(arr); i++ {
	//	clone[i] = arr[len(arr)-1-i]
	//}

	//fmt.Println(clone)

	// --2
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println(arr)
}
