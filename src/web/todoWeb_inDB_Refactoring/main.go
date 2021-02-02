package main

import (
	"net/http"
	"web/todoWeb_inDB_Refactoring_withOAuth/app"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	err := http.ListenAndServe("localhost:3000", n)
	if err != nil {
		panic(err)
	}
}
