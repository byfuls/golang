package main

import "fmt"

func pop(c chan int) {
	fmt.Println("Pop func")
	v := <-c
	fmt.Println(v)
}

func main() {
	/*
		var c chan int
		c = make(chan int, 1)		// channel queue size 1

		c <- 10
		v := <-c

		fmt.Println(v)
	*/

	/*
		var c chan int
		c = make(chan int) // channel queue size 0

		c <- 10				// DeadLock
		v := <-c
		fmt.Println(v)
	*/

	var c chan int
	c = make(chan int)
	// channel queue size 0, 즉 길이가 0인 channel은 다른 Thread에서 데이터를 뺄 때 까지 대기하고 있다
	go pop(c)
	c <- 10
	fmt.Println("Program Done")
}
