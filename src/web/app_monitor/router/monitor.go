package router

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"

	"web/app_monitor/package/statusFlags"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	ONCE = "once"
	LOOP = "loop"
)

type monitorStatusFlags struct {
	AlterUser       bool `json:"alterUser"`
	LurUpRequest    bool `json:"lurUpRequest"`
	SmsSend         bool `json:"smsSend"`
	Call            bool `json:"call"`
	CallDrop        bool `json:"callDrop"`
	PagingRequest   bool `json:"pagingRequest"`
	CallRecvRequest bool `json:"callRecvRequest"`
}

type monitorStatusInfo struct {
	Index int `json:"index"`
	Count int `json:"count"`

	SerialNo         string             `json:"serialNo"`
	UserImsi         string             `json:"userImsi"`
	Status           monitorStatusFlags `json:"status"`
	UsedTime         time.Time          `json:"usedTime"`
	LastReceivedTime time.Time          `json:"lastReceivedTime"`
	WatchDogOn       bool               `json:"watchDogOn"`
}

type message struct {
	Command    string               `json:"command"`
	SubCommand string               `json:"subCommand"`
	Result     bool                 `json:"result"`
	ResultMsg  string               `json:"resultMsg"`
	Data       []*monitorStatusInfo `json:"data"`
}

//func monitorResponse(w *http.ResponseWriter, command string, subCommand string, result bool, resultMsg string, data statusInfo, httpStatus int) {
//	responseMsg := message{
//		Command:    command,
//		SubCommand: subCommand,
//		Result:     result,
//		ResultMsg:  resultMsg,
//		Data:       data,
//	}
//	rd.JSON((*w), httpStatus, responseMsg)
//}

func all(r *message) message {
	allData := []*monitorStatusInfo{}

	var count = 1 // TEST
	for i := 0; i < count; i++ {
		data := new(monitorStatusInfo)

		data.Index = i
		data.Count = count
		data.SerialNo = "serial no:" + strconv.Itoa(i)
		data.UserImsi = "user imsi:" + strconv.Itoa(i)
		data.UsedTime = time.Now()
		data.LastReceivedTime = time.Now()
		data.WatchDogOn = false

		status := statusFlags.Status{}
		for _, flag := range []statusFlags.Bits{
			statusFlags.AlterUser,
			statusFlags.LurUpRequest,
			statusFlags.SmsSend,
			statusFlags.Call,
			statusFlags.CallDrop,
			statusFlags.PagingRequest,
			statusFlags.CallRecvRequest} {

			if status.Has(flag) {
				switch flag {
				case 1: /* AlterUser */
					data.Status.AlterUser = true
				case 2: /* LurUpRequest */
					data.Status.LurUpRequest = true
				case 4: /* SmsSend */
					data.Status.SmsSend = true
				case 8: /* Call */
					data.Status.Call = true
				case 16: /* CallDrop */
					data.Status.CallDrop = true
				case 32: /* PagingRequest */
					data.Status.PagingRequest = true
				case 64: /* CallRecvRequest */
					data.Status.CallRecvRequest = true
				}
			}
		}
		allData = append(allData, data)
	}

	return message{
		Command:    r.Command,
		SubCommand: r.SubCommand,
		Result:     true,
		ResultMsg:  "success",
		Data:       allData,
	}
}

func monitoring(conn *websocket.Conn, r *http.Request) {
	for {
		m := &message{}
		err := conn.ReadJSON(m)
		if err != nil {
			log.Println("[monitorHandler] read error: ", err)
			return
		} else {
			log.Println("[monitorHandler] read ok: ", m)
			log.Println("[monitorHandler] command   : ", m.Command)
			log.Println("[monitorHandler] subCommand: ", m.SubCommand)
			log.Println("[monitorHandler] result    : ", m.Result)
			log.Println("[monitorHandler] resultMsg : ", m.ResultMsg)
			log.Println("[monitorHandler] data      : ", m.Data)

			var data message
			data = all(m)
			err = conn.WriteJSON(data)
			if err != nil {
				log.Println("[monitorHandler] write error: ", err)
				return
			} else {
				log.Println("[monitorHandler] write ok: ", m)
			}
		}
	}
}

func (rt *RouterHandler) monitorHandler(w http.ResponseWriter, r *http.Request) {
	//sessionId := getSessionID(r)
	//log.Println("[monitorHandler] sessionId: ", sessionId)
	// rt.database.ReadUserSession(sessionId)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	go monitoring(conn, r)
}
