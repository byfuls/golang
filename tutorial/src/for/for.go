package main

import "fmt"

func main() {
	// --1
	//i := 0
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}
	//fmt.Println("최종 i 값은: ", i)

	// --2
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	////fmt.Println("최종 i 값은: ", i)

	// --3
	//var i int
	//for {
	//	i++
	//	fmt.Println(i)
	//}

	// --4
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}

	// --5
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	// --6
	var i int
	for { // == for true {}
		i++
		if i == 5 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}
