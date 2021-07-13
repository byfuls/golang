package fcmHandler

import (
	"fmt"

	"github.com/NaySoftware/go-fcm"
)

type FcmSender struct {
	ServerKey string
	Ids       []string
	Message   map[string]interface{}

	Fcm *fcm.FcmClient
}

func (f *FcmSender) GenerateSend(ids []string, noti_title string, noti_body string, data_protocol []byte) (error, bool) {
	if 0 >= len(ids) {
		return nil, false
	}

	f.Ids = ids
	f.Message["title"] = noti_title
	f.Message["body"] = noti_body
	f.Message["data"] = map[string]interface{}{
		"message":      data_protocol,
		"click_action": "FLUTTER_NOTIFICATION_CLICK",
	}

	fmt.Println("dest: ", f.Ids)
	fmt.Println("Message: ", f.Message)

	f.Fcm.NewFcmRegIdsMsg(f.Ids, f.Message)

	status, err := f.Fcm.Send()
	if err != nil {
		return err, false
	} else {
		fmt.Println(status)
	}

	return nil, true
}

func (f *FcmSender) Init(serverKey string) bool {
	if 0 >= len(serverKey) {
		return false
	}

	f.ServerKey = serverKey
	f.Message = make(map[string]interface{})
	f.Fcm = fcm.NewFcmClient(f.ServerKey)

	return true
}
