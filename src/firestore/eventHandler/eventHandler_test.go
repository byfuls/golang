package eventHandler

import (
	"fmt"
	"testing"
	"time"
)

//type EventA struct {
//	Id              string
//	Field_string    string    `firestore:"field_string"`
//	Field_timeStamp time.Time `firestore:"field_timeStamp"`
//	Field_number    int       `firestore:"field_number"`
//}

func go_getSnapShots(no int, eventHandler *EventHandler) {
	fmt.Printf("[get:%d] start...\n", no)
	for {
		docs := <-eventHandler.SnapShotDocs
		fmt.Println("=>get: ", len(docs))
		for i := 0; i < len(docs); i++ {
			fmt.Printf("[get:%d](%d) id              : [%s]\n", no, i, docs[i].Id)
			fmt.Printf("[get:%d](%d) field_string    : [%s]\n", no, i, docs[i].Field_string)
			fmt.Printf("[get:%d](%d) field_timeStamp : [%s]\n", no, i, docs[i].Field_timeStamp)
			fmt.Printf("[get:%d](%d) field_number    : (%d)\n", no, i, docs[i].Field_number)

			//eventHandler.UpdateDoc("test1/log/tt1", "history")
            doc := EventA {
                Field_string: "add_string",
                Field_number: 0,
            }
			eventHandler.AddDoc("test1/log/tt1", doc)
		}
	}
}

func TestMain(t *testing.T) {
	eventHandler := EventHandler{}
	eventHandler.Init("../ttgo-b29bf-firebase-adminsdk-qg1jz-73e3e1bf64.json")
	defer eventHandler.Term()

	for i := 0; i < 1; i++ {
		go go_getSnapShots(i, &eventHandler)
	}
	time.Sleep(1 * time.Second)
	go eventHandler.Snapshots("test")

	for {
		time.Sleep(3 * time.Second)
	}
}
