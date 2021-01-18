package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"childProcess/v2/logging"
	"childProcess/v2/mediaChannel/bypasser"
	"childProcess/v2/mediaChannel/header"
	"childProcess/v2/mediaChannel/ioCommunication"
)

func pipeReader() {
	for {
		buf, len, err := ioCommunication.Read(32)
		if err != nil {
			logging.ErrorLn("[pipeReader] read error: ", err)
			panic(err)
		} else {
			logging.DebugF("[pipeReader] read success: %d\n%s\n",
				len, hex.Dump(buf))
		}
	}
}

func pipeWriter(timeoutSig chan bool) {
	for {
		tmout := <-timeoutSig
		if tmout {
			/* inform timeout message to parent */
			//mediaAddressInfo := &mediaChannel.MediaMessage{
			//	Head: "timeout",
			//	Ip:   "",
			//	Port: "",
			//}
			mediaAddressInfo := mediaChannel.GenerateMediaMessage("timeout", "", "", "")
			mediaAddressByteArray, mediaAddressError := json.Marshal(mediaAddressInfo)
			if mediaAddressError != nil {
				logging.ErrorF("[pipeWriter] json marshal error\n")
				panic(mediaAddressError)
			}

			if len, err := ioCommunication.Write(mediaAddressByteArray); err != nil {
				logging.ErrorLn("[pipeWriter] write error: ", err)
				panic(err)
			} else {
				logging.DebugF("[pipeWriter] write success: %d\n%s\n",
					len, hex.Dump(mediaAddressByteArray))
			}
		}
	}
}

func channelTimeout(timeoutSig chan bool) {
	for {
		timeout := bypasser.Timeout()
		timeoutSig <- timeout
	}
}

func main() {
	if _loggingPath := os.Getenv("B_LOG"); len(_loggingPath) > 0 {
		if !logging.Init(os.Getenv("B_LOG"), "parent_child.log") {
			fmt.Println("logging init error")
			panic("logging init error")
		}
	} else {
		logging.DebugF("[mediaChannel] start\n")
		for i := 0; i < len(os.Args); i++ {
			logging.DebugF("[mediaChannel] args(%d)=[%s]\n", i, os.Args[i])
		}
	}

	ioCommunication.Init()

	/* open new udp server */
	port, _ := strconv.Atoi("0")
	address := net.UDPAddr{Port: port, IP: net.ParseIP("127.0.0.1")}
	udpSock, err := net.ListenUDP("udp", &address)
	if err != nil {
		logging.ErrorLn("[Init] error: ", err)
		return
	}
	tmp := udpSock.LocalAddr().String()
	addr := strings.Split(tmp, ":")
	addrIp := addr[0]
	addrPort, _ := strconv.Atoi(addr[1])
	udpAddr := net.UDPAddr{Port: addrPort, IP: net.ParseIP(addrIp)}
	bypasser.Init(udpAddr, udpSock, 32)
	bypasser.Operator(3)

	/* inform new udp address to parent */
	mediaAddressInfo := mediaChannel.GenerateMediaMessage("mediaAddress", "", addr[0], addr[1])
	mediaAddressByteArray, mediaAddressError := json.Marshal(mediaAddressInfo)
	if mediaAddressError != nil {
		logging.ErrorF("[mediaChannel] json marshal error[%s:%s]\n", addr[0], addr[1])
		return
	} else {
		if len, err := ioCommunication.Write(mediaAddressByteArray); err != nil {
			logging.ErrorLn("[mediaChannel] write error: ", err)
			panic(err)
		} else {
			logging.DebugF("[mediaChannel] write success: %d\n%s\n",
				len, hex.Dump(mediaAddressByteArray))
		}
	}

	timeoutSig := make(chan bool)
	go pipeReader() // not used now
	go pipeWriter(timeoutSig)
	go channelTimeout(timeoutSig)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		sig := <-sigs
		logging.TraceF("[mediaChannel] received signal, end: %d\n", sig)
		bypasser.Done()
		ioCommunication.Done()
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()

	for {
		time.Sleep(3 * time.Second)
	}
}
