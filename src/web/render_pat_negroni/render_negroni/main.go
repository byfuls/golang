package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "byfuls", Email: "byfuls@gmail.com"}

	rd.JSON(w, http.StatusOK, user)
	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(user)
	//fmt.Fprint(w, string(data))
}

func addUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		//w.WriteHeader(http.StatusBadRequest)
		//fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()

	rd.JSON(w, http.StatusOK, user)
	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(user)
	//fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//tmpl, err := template.New("hello").ParseFiles("templates/hello.tmpl")
	//if err != nil {
	//	rd.Text(w, http.StatusBadRequest, err.Error())
	//	//w.WriteHeader(http.StatusBadRequest)
	//	//fmt.Fprint(w, err)
	//	return
	//}
	//tmpl.ExecuteTemplate(w, "hello.tmpl", "byfuls")

	user := User{Name: "byfuls", Email: "byfuls@gmail.com"}
	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	rd = render.New(render.Options{
		Directory:  "templates_rename",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserInfoHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	// = mux.Handle("/", http.FileServer(http.Dir("public")))
	// include file server, log, panic, ... in negroni

	http.ListenAndServe("127.0.0.1:3000", n)
}
