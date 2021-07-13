package main

import "fmt"

type Student struct {
	name  string
	class int
	grade Grade
}

type Grade struct {
	name  string
	grade string
}

func (s Student) ViewGrade() {
	fmt.Println(s.grade)
}

/* general function
func ViewGrade(s Student) {
	fmt.Println(s.grade)
}
*/

func (s Student) InputGrade(name string, grade string) {
	s.grade.name = name
	s.grade.grade = grade
}

func InputGrade(s Student, name string, grade string) {
	s.grade.name = name
	s.grade.grade = grade
}

func main() {
	var s Student
	s.name = "Talyor"
	s.class = 1

	s.grade.name = "math"
	s.grade.grade = "A+"

	s.ViewGrade()
	/* ViewGrade(s) */

	s.InputGrade("Computing", "A+")
	s.ViewGrade()
}
