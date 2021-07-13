package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	cuz := now.Format("2006-01-02 15:04:05")
	fmt.Println(cuz)

	cuz0 := now.Format("20060102 150405")
	fmt.Println(cuz0)

	cuz1 := now.Format("20060102")
	fmt.Println(cuz1)

	cuz2 := now.Format("15:04:05")
	fmt.Println(cuz2)
}
