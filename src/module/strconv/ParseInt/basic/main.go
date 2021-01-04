package main

import (
	"fmt"
	"strconv"
	"reflect"
	//"encoding/hex"
)

func main() {
	var h string
	h = "31"

	tmp, err := strconv.ParseInt(h, 16, 64)
	if err != nil {
		fmt.Println("convert error")
		return
	}
	tmpType := reflect.TypeOf(tmp)
	fmt.Printf("str str[%s] dec(%d)\n", h, tmp)
	fmt.Printf("converted type: %s\n", tmpType)
}
