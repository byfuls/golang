package main

import (
	"net/http"

	"web/basic/myapp"
)

func main() {
	http.ListenAndServe("127.0.0.1:3000", myapp.NewHttpHandler())
}
