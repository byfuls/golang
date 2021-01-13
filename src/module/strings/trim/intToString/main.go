package main

import (
	_ "bytes"
	"fmt"
	_ "strconv"
	"strings"
)

func main() {
	A := []int{1, 2, 3}
	delim := ""

	fmt.Println("s:", fmt.Sprint(A))

	st := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(A)), delim), "[ ]")
	fmt.Println("st:", st)
	fmt.Println("len:", len(st))
}
