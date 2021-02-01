package main

type Component interface {
	Operator(string)
}

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	// Send Data
	sendData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {

}
