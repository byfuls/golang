package main

import "fmt"

func main() {
	one := make([]byte, 2)
	two := make([]byte, 2)
	one[0] = 0x00
	one[1] = 0x01
	two[0] = 0x02
	two[1] = 0x03
	fmt.Println(append(one[:], two[:]...))

	three := []byte{0, 1}
	four := []byte{2, 3}
	five := append(three, four...)
	fmt.Println(five)
}
