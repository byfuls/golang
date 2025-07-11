package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	fmt.Println("jobs: ", jobs)
	go func() {
		for {
			j, more := <-jobs
			fmt.Println("more: ", more)
			if more {
				fmt.Println("received job: ", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// for j := 1; j <= 3; j++ {
	// 	jobs <- j
	// 	fmt.Println("sent job: ", j)
	// }
	close(jobs)
	fmt.Println("jobs: ", jobs)
	fmt.Println("sent all jobs")

	<-done
}
