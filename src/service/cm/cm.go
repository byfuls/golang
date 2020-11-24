package main

import (
	"fmt"
	"time"

	"service/cm/bridge"
	"service/cm/channel"
	"service/cm/handling"

	"service/proxy/channelManager"
)

const (
	DEFAULT_TCP_IP   = "127.0.0.1"
	DEFAULT_TCP_PORT = 10100

	DEFAULT_PROXY_TCP_IP   = "127.0.0.1"
	DEFAULT_PROXY_TCP_PORT = 10001

	HANDLER_COUNT = 5
)

func main() {
	fmt.Println("[cm] start ...")

	fmt.Println("[cm] server address: ", DEFAULT_TCP_IP, DEFAULT_TCP_PORT)
	fmt.Println("[cm] remote proxy address: ", DEFAULT_PROXY_TCP_IP, DEFAULT_PROXY_TCP_PORT)

	channelManager.Init()

	chToDv := make(chan handling.Message)
	pxToDv := make(chan handling.Message)

	dvToPx := make(chan handling.Message)
	dvToCh := make(chan handling.Message)

	for no := 0; no < HANDLER_COUNT; no++ {
		go handling.Deliver(chToDv, pxToDv, dvToPx, dvToCh, no)
	}
	go channel.Accepter(chToDv, DEFAULT_TCP_IP, DEFAULT_TCP_PORT)                  // CM - GW
	go bridge.Bridge(pxToDv, dvToPx, DEFAULT_PROXY_TCP_IP, DEFAULT_PROXY_TCP_PORT) // CM - PROXY

	for {
		time.Sleep(1 * time.Second)
	}
}
