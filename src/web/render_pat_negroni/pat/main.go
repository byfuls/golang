package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/pat"
)

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "byfuls", Email: "byfuls@gmail.com"}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func addUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("hello").ParseFiles("templates/hello.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	tmpl.ExecuteTemplate(w, "hello.tmpl", "byfuls")
}

func main() {
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserInfoHandler)
	mux.Get("/hello", helloHandler)

	http.ListenAndServe("127.0.0.1:3000", mux)
}
