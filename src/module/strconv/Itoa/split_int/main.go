package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	i := 51010
	s := strconv.Itoa(i)

	fmt.Println(s)
	fmt.Printf("[0-3]=%s (data type: %s)\n", s[0:3], reflect.TypeOf(s).String())
	fmt.Printf("[3-5]=%s (data type: %s)\n", s[3:], reflect.TypeOf(s).String())

	i_1, _ := strconv.Atoi(s[0:3])
	i_2, _ := strconv.Atoi(s[3:])
	fmt.Printf("ret 1: %d (data type: %s)\n", i_1, reflect.TypeOf(i_1).String())
	fmt.Printf("ret 2: %d (data type: %s)\n", i_2, reflect.TypeOf(i_2).String())
}
