package main

import (
	"fmt"
	"strings"
)

func dynamicArgumentsFunc(basic string, v ...string) {
	fmt.Println("basic: ", basic)
	fmt.Println("arguments: ", v)

	fmt.Println("arguments count: ", len(v))

	justString := strings.Join(v, " ")
	fmt.Println("space: ", justString)

	secondString := "\"" + strings.Join(v, "\", \"") + "\""
	fmt.Println("comma: ", secondString)
}

func main() {
	dynamicArgumentsFunc("First")
	dynamicArgumentsFunc("First", "1", "2")
	dynamicArgumentsFunc("First", "-ip", "127.0.0.1", "-port", "1234")
}
