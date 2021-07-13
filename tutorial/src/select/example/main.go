package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	var i int
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)

	go push(c)

	timerChan := time.After(10 * time.Second)   // time.After 일정 시간 뒤에 채널을 리턴함 (1번만)
	tickTimerChan := time.Tick(2 * time.Second) // time.Tick 일정 주기마다 계속해서 채널을 리턴 (주기적)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		//default:
		//	fmt.Println("idle")
		//	time.Sleep(1 * time.Second)
		case <-timerChan:
			fmt.Println("timeout")
			return
		case <-tickTimerChan:
			fmt.Println("tick")
		}
	}
}
