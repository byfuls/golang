package router

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
	"time"
)

type Login struct {
	Result    bool   `json:"result"`
	ResultMsg string `json:"result_message"`
	Id        string `json:"id"`
}

func checkLogined(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if strings.Contains(r.URL.Path, "/login") {
		next(w, r)
		return
	}

	sessionID := getSessionID(r)
	log.Println("[checkLogined] sessionID: ", sessionID)
	if sessionID != "" {
		next(w, r)
		return
	}

	http.Redirect(w, r, "/login.html", http.StatusTemporaryRedirect)
}

func generateStateStateCookie(w *http.ResponseWriter) string {
	expiration := time.Now().Add(1 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{
		Name:    "sessionstate",
		Value:   state,
		Expires: expiration,
	}
	http.SetCookie((*w), cookie)
	return state
}

func loginResponse(w *http.ResponseWriter, result bool, resultMsg string, id string, httpStatus int) {
	responseMsg := Login{
		Result:    result,
		ResultMsg: resultMsg,
		Id:        id,
	}
	log.Println("[loginResponse] ", responseMsg)
	rd.JSON((*w), httpStatus, responseMsg)
}

func (rt *RouterHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	password := r.FormValue("password")

	/*_____ Check ID, Password in database _____*/
	var result = false
	list := rt.database.ReadUserId(id)
	log.Println("[loginHandler] list: ", list, len(list))
	if list != nil && (len(list) == 1) {
		log.Println("[loginHandler] ", list[0])
		log.Println("[loginHandler] Id       : ", list[0].Id)
		log.Println("[loginHandler] Password : ", list[0].Password)
		log.Println("[loginHandler] SessionID: ", list[0].SessionId)
		log.Println("[loginHandler] CreatedAt: ", list[0].CreatedAt)
		log.Println("[loginHandler] UpdatedAt: ", list[0].UpdatedAt)
		if id == list[0].Id && password == list[0].Password {
			result = true
		}
	}
	log.Println("[loginHandler] result: ", result)

	/*_____ Response _____*/
	if result {
		sessionId := generateStateStateCookie(&w)
		log.Println("[loginHandler] session-id: ", sessionId)
		if !rt.database.UpdateUser(id, sessionId) {
			log.Println("[loginHandler] updateUser fail, ", id, sessionId)
			loginResponse(&w, false, "login failed", id, http.StatusInternalServerError)
			return
		} else {
			log.Println("[loginHandler] updateUser success, ", id, sessionId)

			session, err := store.Get(r, "session")
			if err != nil {
				log.Println("[loginHandler] session store error: ", err.Error())
				loginResponse(&w, false, "session store error", id, http.StatusInternalServerError)
				return
			}
			session.Values["id"] = list[0].Id
			//session.Values["updatedAt"] = list[0].UpdatedAt
			err = session.Save(r, w)
			if err != nil {
				log.Println("[loginHandler] session save error: ", err.Error())
				loginResponse(&w, false, "session save error", id, http.StatusInternalServerError)
				return
			}

			loginResponse(&w, result, "success", id, http.StatusOK)
		}
	} else {
		loginResponse(&w, result, "not exist", id, http.StatusInternalServerError)
	}
}
