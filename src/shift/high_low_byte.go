package main

import (
	"fmt"
)

func main() {
	val := 0x0102
	fmt.Println(val)

	high := val>>8 & 0xff
	low := val & 0xff
	fmt.Println("high: ", high)
	fmt.Println("low : ", low)

	////////////////////////////
	fmt.Println("////////")
	////////////////////////////

	var test uint8 = 0xff
	fmt.Printf("%08b\n", test)

	test<<=1
	fmt.Printf("%08b\n", test)

	test<<=1
	fmt.Printf("%08b\n", test)
}
