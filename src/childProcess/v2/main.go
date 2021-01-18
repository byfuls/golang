package main

import (
	"bytes"
	"encoding/hex"
	_ "encoding/json"
	"fmt"
	"os"
	"time"

	"childProcess/v2/child"
	"childProcess/v2/logging"
	"childProcess/v2/mediaChannel/header"
)

func readFromChildPcs(cp *child.Process) {
	for {
		buf := make([]byte, 128)
		len, err := cp.Read(&buf)
		if err != nil {
			logging.ErrorLn("[readFromChildPcs] read error: ", err)
			cp.Stop()
			return
		} else {
			logging.DebugF("[readFromChildPcs] read len(%d)\n%s\n",
				len, hex.Dump(buf[:len]))

			mediaMessage := mediaChannel.ParsingMediaMessage(buf[:len])
			if mediaMessage.Head == "mediaAddress" {
				logging.TraceF("[readFromChildPcs] get media channel address[%s:%s]\n",
					mediaMessage.Ip, mediaMessage.Port)
			} else if mediaMessage.Head == "timeout" {
				cp.Stop()
				logging.TraceF("[readFromChildPcs] TIMEOUT, shut down child process\n")
				return
			} else if mediaMessage.Head == "error" {
				cp.Stop()
				logging.TraceF("[readFromChildPcs] ERROR(1), shut down child process\n")
				return
			} else if bytes.Contains(buf[:len], []byte("error")) {
				cp.Stop()
				logging.TraceF("[readFromChildPcs] ERROR(2), shut down child process\n")
				return
			} else {
				cp.Stop()
				logging.TraceF("[readFromChildPcs] ERROR(*), unknown message received, shut down child process\n")
				return
			}
		}
	}
}

func sendToChildPcs(cp *child.Process) {
	buf := []byte("TEMP TEST MESSAGE")
	for {
		if len, err := cp.Write(buf); err != nil {
			logging.ErrorLn("[sendToChildPcs] write error: ", err)
			cp.Stop()
			cp = nil
			return
		} else {
			logging.DebugF("[sendToChildPcs] write len(%d)\n%s\n",
				len, hex.Dump(buf))
		}

		time.Sleep(3 * time.Second)
	}
}

func main() {
	if _loggingPath := os.Getenv("B_LOG"); len(_loggingPath) > 0 {
		if !logging.Init(os.Getenv("B_LOG"), "parent_child.log") {
			fmt.Println("logging init error")
			panic("logging init error")
		}
	} else {
		panic("logging env path check again ...")
	}

	cp := child.Init()
	if err := cp.Run("./mediaChannel/mediaChannel"); err != nil {
		logging.ErrorLn("Run child process error: ", err)
		panic(err)
	} else {
		logging.TraceLn("Run child process success")
		go readFromChildPcs(cp)
		//go sendToChildPcs(cp)
	}

	var count int
	for {
		if cp.Status() {
			time.Sleep(1 * time.Second)
			continue
		}

		if !cp.Status() {
			if count > 5 {
				cp = child.Init()
				if err := cp.Run("./mediaChannel/mediaChannel"); err != nil {
					logging.ErrorLn("Run child process error: ", err)
					panic(err)
				} else {
					logging.TraceLn("Run child process success")
					go readFromChildPcs(cp)
					//go sendToChildPcs(cp)
				}
				count = 0
			} else {
				count++
			}
		}
		time.Sleep(1 * time.Second)
	}
}
