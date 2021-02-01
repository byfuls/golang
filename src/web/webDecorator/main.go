package main

import (
	"log"
	"net/http"
	"time"

	"web/rest/myapp"
	"web/webDecorator/decoHandler"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[logger1] started")
	h.ServeHTTP(w, r)
	log.Println("[logger1] completed time: ", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[logger2] started")
	h.ServeHTTP(w, r)
	log.Println("[logger2] completed time: ", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decoHandler.NewDecoHandler(h, logger)
	h = decoHandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe("127.0.0.1:3000", mux)
}
