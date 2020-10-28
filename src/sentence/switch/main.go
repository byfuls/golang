package main

import "fmt"

func main() {
	var val = [...]string{"STR", "TTT", "TEST1", "TEST2"}

	for i, v := range val {
		fmt.Printf("[input] idx: %d, value: %s\n", i, v)
		switch v {
		case "STR":
			fmt.Println("STR")
		case "TTT":
			fmt.Println("TTT")
		case "TEST1", "TEST2":
			fmt.Println("TEST1 or TEST2")
		default:
			fmt.Println("ERROR")
		}
	}
}
