package main

import (
	"fmt"
	"strconv"
	"reflect"
	//"encoding/hex"
)

func main() {
	var i int
	i = 1234

	tmp := strconv.Itoa(i)
	tmpType := reflect.TypeOf(tmp)
	fmt.Println(tmp)
	fmt.Println(tmpType)

	tmp1 := strconv.FormatInt(3, 16)
	tmp1Type := reflect.TypeOf(tmp1)
	fmt.Println(tmp1)
	fmt.Println(tmp1Type)

	k := 255
	h := fmt.Sprintf("%x", k)
	fmt.Printf("Hex conv of '%d' is '%s'\n", k, h)
	h = fmt.Sprintf("%X", k)
	fmt.Printf("HEX conv of '%d' is '%s'\n", k, h)

	q := 1
	q1 := fmt.Sprintf("%x", q)
	fmt.Printf("%T, %s\n", q1, q1)
	t := byte(string(q)[0])
	fmt.Printf("%T, %d\n", t, t)

	fmt.Println("==", string(q)[0])
}
