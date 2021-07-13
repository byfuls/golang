package main

import (
	"container/list"
	"fmt"
)

type structData struct {
	name string
	age  int
}

func main() {
	l := list.New()

	s1 := new(structData)
	s1.name = "Jason"
	s1.age = 18
	fmt.Printf("%p %p\n", &s1, s1)

	l.PushBack(s1)
	//e4 := l.PushBack(s1)
	//e1 := l.PushFront(1)

	//l.InsertBefore(3, e4)
	//l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		newData := e.Value.(*structData)
		fmt.Println(newData)
		fmt.Printf("%p %p\n", &newData, newData)
		fmt.Println(newData.name)
	}
}
