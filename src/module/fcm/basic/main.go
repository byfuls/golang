package main

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
)

const (
	serverKey = "AAAAPnQqdlw:APA91bEzssQ_He1a5eVj5snwaG7MgDBabtMxDb1cqAAjSqDYGxiJhtF3xo1SXfmY6vDOLVbeOD-GiEetdTdmclXvX8VRwhUNBYbld15iStVAJCsg0cviMEJzUAage6WFZ7yr77JsUneo"
)

func main() {
	//data := map[string]string{
	//	"msg": "Hello world1",
	//	"sum": "Happy day",
	//}
	data := map[string][]byte{
		"msg": []byte("\x01x02x03"),
	}
	ids := []string{
		"c8FLM-SQSoifTT4iMWtV1D:APA91bGad1aeuUUSYX15IDmTSPcOdAtbD4oPzDcVdgUYLpaLTYkVDbb5XaLMa5CXuoXjNOhSrDVcfpEp2srw3AhvJn2bzoyLEdSiNNTG3r6_EmOsTfppsYmkSbhVkICZ9-thWxlaI93h",
	}
	fmt.Println("ids length: ", len(ids))

	fmt.Println("dest: ", ids)
	fmt.Println("Message: ", data)

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)

	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
