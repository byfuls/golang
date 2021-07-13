package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintGrade() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputGrade(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student = Student{name: "byfuls", age: 23, class: "Math", grade: "A+"}

	s.InputGrade("Computing", "A+")
	s.PrintGrade()
}
