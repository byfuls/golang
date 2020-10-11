package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Println(p.name)
}

func main() {
	var p Person
	p1 := Person{"홍길동", 15}
	p2 := Person{name: "심청이", age: 21}
	p3 := Person{name: "Jason"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "Peter"
	p.age = 24

	fmt.Println(p, p1, p2, p3, p4)

	p.PrintName()
	p1.PrintName()
}
