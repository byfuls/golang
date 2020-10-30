package handling

import (
	"encoding/hex"
	"fmt"
	"net"

	"service/proxy/channelManager"

	"service/protocol"
)

const (
	CH    = iota
	PROXY = iota
)

type Message struct {
	From int
	Buf  []byte
}

type CHManage struct {
	Socket *net.TCPConn
	ToCH   chan Message
}

/***********************************/
func registCM() ([]byte, uint32) {
	p := protocol.Packet{}
	TEST_KEY := "TEST"
	buf, len := p.Make("CONNECTION", "REGIST", 0, len(TEST_KEY), TEST_KEY, 0, nil)
	if 0 == len {
		fmt.Println("[registCM] make regist protocol error")
		return nil, 0
	}
	return buf, len
}

/***********************************/

func Deliver(chToDv chan Message, pxToDv chan Message, dvToPx chan Message, dvToCh chan Message) {
	fmt.Println("[Deliver] start")

	for {
		select {
		case message := <-chToDv:
			switch message.From {
			case CH:
				fmt.Println("[Deliver] (CH) receive data")
				fmt.Printf("[Deliver] receive channel: \n%s\n", hex.Dump(message.Buf))

				dvToPx <- Message{
					From: message.From,
					Buf:  message.Buf,
				}
			default:
				fmt.Println("[Deliver] (from Channel) unknown FromMessage error: ", message.From)
			}
		case message := <-pxToDv:
			switch message.From {
			case PROXY:
				fmt.Println("[Deliver] (PROXY) receive data")
				fmt.Printf("[Deliver] receive channel: \n%s\n", hex.Dump(message.Buf))

				if tmp, ret := channelManager.Get("TEST"); ret {
					ch := tmp.(CHManage)
					ch.ToCH <- Message{
						From: PROXY,
						Buf:  message.Buf,
					}
				} else {
					fmt.Println("[Deliver] (PROXY) not found channel in map")
					continue
				}
			default:
				fmt.Println("[Deliver] (from Proxy) unknown FromMessage error: ", message.From)
			}
		}
	}
}
