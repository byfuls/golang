package main

import (
	"fmt"
	"strings"
	"strconv"
	"reflect"
)

func main() {
	s := "31"
	numberStr := strings.Replace(s, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)

	tmpType := reflect.TypeOf(numberStr)
	fmt.Println(numberStr, tmpType)

	n, err := strconv.ParseUint(numberStr, 16, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
