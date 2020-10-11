package main

import "fmt"

// OOP => E.R (Entity Relactionship)

type Student struct {
	name  string
	age   int
	grade int
}

//func SetName(t *Student, newName string) {
func (t *Student) SetName(newName string) {
	t.name = newName
}

func (t *Student) SetAge(age int) {
	t.age = age
}

func main() {
	//a := Student{"aaa", 20, 10}

	// --1
	//var b *Student
	//b = &a

	//b.age = 30

	//fmt.Println(a)
	//fmt.Println(b)

	// --2
	////SetName(&a, "bbb")
	//a.SetName("bbb")
	//// a : instance (생명주기) , 주체 주어
	////	instance = structure pointer type
	//fmt.Println(a)

	// --3
	var b *Student
	b = &Student{"aaa", 20, 10}
	b.SetName("bbb")
	b.SetAge(30)
	fmt.Println(b)
}
