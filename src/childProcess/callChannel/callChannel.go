package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	_ "fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"childProcess/callChannel/bypasser"
	"childProcess/callChannel/ioCommunication"
	"childProcess/logging"
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
			buf := []byte("TIMEOUT") // get sending data via channel
			if len, err := ioCommunication.Write(buf); err != nil {
				logging.ErrorLn("[pipeWriter] write error: ", err)
				panic(err)
			} else {
				logging.DebugF("[pipeWriter] write success: %d\n%s\n",
					len, hex.Dump(buf))
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
		logging.DebugF("[callChannel] start\n")
		for i := 0; i < len(os.Args); i++ {
			logging.DebugF("[callChannel] args(%d)=[%s]\n", i, os.Args[i])
		}
	}

	ip := flag.String("ip", "", "call channel ip address")
	port := flag.Int("port", 0, "call cahnnel port number")
	flag.Parse()
	if flag.NFlag() == 0 {
		logging.ErrorLn("[option] check parameter")
		flag.Usage()
		return
	} else {
		logging.DebugF("[option] ip: %s\n", *ip)
		logging.DebugF("[option] port: %d\n", *port)
	}

	ioCommunication.Init()
	bypasser.Init(*ip, *port, 32) // get ip address from parent
	bypasser.Operator(3)

	timeoutSig := make(chan bool)
	go pipeReader() // not used now
	go pipeWriter(timeoutSig)
	go channelTimeout(timeoutSig)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		sig := <-sigs
		logging.TraceF("[callChannel] (1) End: %d\n", sig)
		bypasser.Done()
		logging.TraceF("[callChannel] (2) End: %d\n", sig)
		ioCommunication.Done()
		logging.TraceF("[callChannel] (3) End: %d\n", sig)
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()

	for {
		logging.TraceF("[callChannel] ------\n")
		time.Sleep(3 * time.Second)
	}
}
