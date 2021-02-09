package router

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"

	"program/cm/channelMonitor/model"
)

var (
	store                = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	rd    *render.Render = render.New()
)

type RouterHandler struct {
	http.Handler
	database model.DatabaseHandler
}

func getSessionID(r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}

	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

func (rt *RouterHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/monitor.html", http.StatusTemporaryRedirect)
}

func (rt *RouterHandler) Close() {
	rt.database.Close()
}

func MakeHandler(databaseFilePath string) *RouterHandler {
	m := mux.NewRouter()

	n := negroni.New(negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(checkLogined),
		negroni.NewStatic(http.Dir(os.Getenv("COMMON_SRC")+"/cm/channelMonitor/public")))
	n.UseHandler(m)

	handler := &RouterHandler{
		Handler:  n,
		database: model.NewDatabaseHandlerUser(databaseFilePath),
	}

	m.HandleFunc("/monitor", handler.monitorHandler).Methods("GET")
	m.HandleFunc("/login", handler.loginHandler)
	m.HandleFunc("/", handler.indexHandler)

	return handler
}
