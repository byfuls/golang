package main

import (
	"fmt"
	"strconv"
	"reflect"
	//"encoding/hex"
)

func main() {
	var s string
	s = "31"

	tmp, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("convert error")
		return
	}
	tmpType := reflect.TypeOf(tmp)
	fmt.Printf("str str[%s] dec(%d)\n", s, tmp)
	fmt.Printf("converted type: %s\n", tmpType)
}
