package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

func MakeTire(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Tire, "

		outChan <- car
	}
}

func MakeEngine(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Engine, "

		outChan <- car
	}
}

func StartWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car " + strconv.Itoa(i)}
		i++
	}
}

func main() {
	chan1 := make(chan Car)
	chan2 := make(chan Car)
	chan3 := make(chan Car)

	go StartWork(chan1)
	go MakeTire(chan1, chan2)
	go MakeEngine(chan2, chan3)

	for {
		result := <-chan3
		fmt.Println(result.val)
	}
}
