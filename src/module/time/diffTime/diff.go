package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	//t2 := t1.Add(time.Second * 341)
	t2 := t1.Add(time.Second * 17)

	fmt.Println(t1)
	fmt.Println(t2)

	diff := t2.Sub(t1)
	fmt.Println(diff)
	fmt.Println(int64(diff/time.Second))
}
