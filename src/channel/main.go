package main

import (
	"fmt"
	"time"
)

func channelReceiver(msg chan string) {
	for {
		// [like blocking]
		//received := <-msg
		//fmt.Println("received message: ", received)

		// [like non-blocking]
		select {
		case received := <-msg:
			fmt.Println("received message: ", received)
		default:
			fmt.Println("no message received")
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	messages := make(chan string)

	go channelReceiver(messages)

	msg := "hi"
	for {
		time.Sleep(3 * time.Second)
		messages <- msg
	}
}
