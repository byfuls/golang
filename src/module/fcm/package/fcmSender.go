package fcmSender

import (
	"fmt"

	"github.com/NaySoftware/go-fcm"
)

type FcmHandler struct {
	ServerKey string
	Ids       []string
	Message   map[string][]byte

	Fcm *fcm.FcmClient
}

func (f *FcmHandler) GenerateSend(ids []string, protocol []byte) (error, bool) {
	if 0 >= len(ids) {
		return nil, false
	}

	f.Ids = ids
	f.Message["protocol"] = protocol

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

func (f *FcmHandler) Init(serverKey string) bool {
	if 0 >= len(serverKey) {
		return false
	}

	f.ServerKey = serverKey
	f.Message = make(map[string][]byte)
	f.Fcm = fcm.NewFcmClient(f.ServerKey)

	return true
}
