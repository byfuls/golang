package main

import (
	"net/http"
	"web/login_session_perClientID/app"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()

	err := http.ListenAndServe("localhost:3000", m)
	if err != nil {
		panic(err)
	}
}
