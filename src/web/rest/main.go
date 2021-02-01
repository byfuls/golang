package main

import (
	"net/http"

	"web/rest/myapp"
)

func main() {
	http.ListenAndServe("127.0.0.1:3000", myapp.NewHandler())
}
