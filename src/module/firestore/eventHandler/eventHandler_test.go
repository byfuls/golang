package eventHandler

import (
	"fmt"
	"testing"
	"time"
)

func go_SnapShots(eventHandler *EventHandler) {
	eventHandler.Snapshots("test")
}

func go_getSnapShots(eventHandler *EventHandler) {
	for {
		sdocs := <-eventHandler.SnapShotDocs
		fmt.Println("=>get")
		fmt.Println(sdocs)
	}
}

func TestMain(t *testing.T) {
	eventHandler := EventHandler{}
	eventHandler.Init("../ttgo-b29bf-firebase-adminsdk-qg1jz-73e3e1bf64.json")

	//go go_SnapShots(&eventHandler)
	go eventHandler.Snapshots("test")
	go go_getSnapShots(&eventHandler)

	//fmt.Println("docs count: ", cnt)
	//fmt.Println(docs)

	for {
		time.Sleep(3 * time.Second)
	}
}
