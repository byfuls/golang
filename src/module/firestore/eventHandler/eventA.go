package eventHandler

import "time"

type EventA struct {
	Id              string
	Field_string    string    `firestore:"field_string"`
	Field_timeStamp time.Time `firestore:"field_timeStamp"`
	Field_number    int       `firestore:"field_number"`
}
