package fcmSender

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fcm := FcmHandler{}

	serverKey := "AAAAPnQqdlw:APA91bEzssQ_He1a5eVj5snwaG7MgDBabtMxDb1cqAAjSqDYGxiJhtF3xo1SXfmY6vDOLVbeOD-GiEetdTdmclXvX8VRwhUNBYbld15iStVAJCsg0cviMEJzUAage6WFZ7yr77JsUneo"
	if ret := fcm.Init(serverKey); !ret {
		fmt.Println("fcm init error")
		return
	}

	Ids := []string{
		"c8FLM-SQSoifTT4iMWtV1D:APA91bGad1aeuUUSYX15IDmTSPcOdAtbD4oPzDcVdgUYLpaLTYkVDbb5XaLMa5CXuoXjNOhSrDVcfpEp2srw3AhvJn2bzoyLEdSiNNTG3r6_EmOsTfppsYmkSbhVkICZ9-thWxlaI93h",
	}
	msg := []byte("\x01x02x03")
	if err, ret := fcm.GenerateSend(Ids, msg); !ret {
		fmt.Println("fcm generate send error: ", err)
		return
	}
}
