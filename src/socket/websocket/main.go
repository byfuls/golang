package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type MonitInfo struct {
	Dest string `json:"dest"`

	SerialNo         string    `json:"serialNo"`
	UserImsi         string    `json:"userImsi"`
	Status           []byte    `json:"status"`
	UsedTime         time.Time `json:"usedTime"`
	LastReceivedTime time.Time `json:"lastReceivedTime"`
	WatchDogOn       bool      `json:"watchDogOn"`
}

type User struct {
	Id        string    `json:"id"`
	Password  string    `json:"password"`
	LoginTime time.Time `json:"login_time"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(*http.Request) bool { return true },
}

func monitor(w http.ResponseWriter, r *http.Request) {
	var err error

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket err: ", err)
		return
	} else {
		log.Println("websocket ok: ", conn)
	}

	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error: ", err)
			break
		} else {
			log.Println("mt: ", mt)
			log.Printf("read data: \n%s\n", hex.Dump(message))
		}

		var monit MonitInfo
		_ = json.Unmarshal(message, &monit)
		for {

			monit.SerialNo = "serial-no"
			monit.UserImsi = "user-imsi"
			monit.Status = []byte("\x01\x02")
			monit.UsedTime = time.Now()
			monit.LastReceivedTime = time.Now()
			monit.WatchDogOn = false

			data, _ := json.Marshal(monit)
			err = conn.WriteMessage(mt, data)
			if err != nil {
				log.Println("write error: ", err)
				break
			}

			time.Sleep(1 * time.Second)
		}
	}

	/*
		monit := new(MonitInfo)
		err = json.NewDecoder(r.Body).Decode(monit)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}

		monit.SerialNo = "serial-no"
		monit.UserImsi = "user-imsi"
		monit.Status = []byte("\x01\x02")
		monit.UsedTime = time.Now()
		monit.LastReceivedTime = time.Now()
		monit.WatchDogOn = false

		data, _ := json.Marshal(monit)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(data))
	*/
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.LoginTime = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
	return
}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/login", login)
	mx.HandleFunc("/monitor", monitor)

	http.Handle("/", mx)
	http.ListenAndServe("127.0.0.1:2219", mx)
}
