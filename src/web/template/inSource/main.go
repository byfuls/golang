package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user := User{Name: "byfuls", Email: "byfuls@gmail.com", Age: 23}
	user2 := User{Name: "aaa", Email: "aaa@gmail.com", Age: 21}
	tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge: {{.Age}}\n")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, user)
	tmpl.Execute(os.Stdout, user2)
}
