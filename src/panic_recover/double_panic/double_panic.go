package main

// ref >> https://cjwoov.tistory.com/8

import "log"

var (
	panicChannel chan bool = nil
)

func closeRecover() {
	log.Printf("[closeRecover] catch panic")

	if rcv := recover(); rcv != nil {
		log.Printf("[closeRecover] catch >> ", rcv)
	}

	log.Printf("[closeRecover] after panic")
}

func closeFunc() {
	defer closeRecover()
	// closeRecover()

	log.Printf("[closeFunc] catch panic & panic occur")

	if rcv := recover(); rcv != nil {
		log.Printf("[closeFunc] catch >> ", rcv)
	}

	close(panicChannel)

	log.Printf("[closeFunc] catch panic & after panic")
}

func main() {
	// defer closeRecover()
	defer closeFunc()

	log.Printf("[main] panic occur")

	close(panicChannel)

	log.Printf("[main] after panic")
}
