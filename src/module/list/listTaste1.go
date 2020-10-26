package main

import (
	"container/list"
	"fmt"
)

type structData struct {
	name string
	age int
}

type interfaceTest interface {
	name string
}

func main() {
	l := list.New()

	s1 := structData{"Jason", 18}
	fmt.Printf("%p %p\n", &s1, s1)

	e4 := l.PushBack(s1)
	e1 := l.PushFront(1)

	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}


	var s1Point *list.Element
	s1Point = e4

	//fmt.Printf("%p %p\n", &s1Point, s1Point)

	fmt.Println("11111111111111")
	fmt.Println(s1Point)
	fmt.Println(s1Point.Value)
	fmt.Println("11111111111111")

	var s1Data interfaceTest{}
	s1Data = s1Point.Value
	fmt.Println(s1Data)

}
