package main

import (
	"net/http"
	"web/todoWeb_inDB_beforeRefactoring/app"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	err := http.ListenAndServe("localhost:3000", n)
	if err != nil {
		panic(err)
	}
}
