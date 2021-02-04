package main

import (
	"net/http"

	"web/app_monitor/router"
)

func main() {
	m := router.MakeHandler("./database/user.db")
	defer m.Close()

	err := http.ListenAndServe("localhost:2219", m)
	if err != nil {
		panic(err)
	}
}
