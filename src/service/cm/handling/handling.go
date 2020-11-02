package handling

import (
	"encoding/hex"
	"fmt"
	"net"

	"service/protocol"
	"service/proxy/channelManager"
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

func Deliver(chToDv chan Message, pxToDv chan Message, dvToPx chan Message, dvToCh chan Message, no int) {
	fmt.Printf("[Deliver/%d] start\n", no)

	for {
		select {
		case message := <-chToDv:
			switch message.From {
			case CH:
				fmt.Printf("[Deliver/%d] (CH) receive data\n", no)
				fmt.Printf("[Deliver/%d] receive channel: \n%s\n", no, hex.Dump(message.Buf))

				dvToPx <- Message{
					From: message.From,
					Buf:  message.Buf,
				}
			default:
				fmt.Printf("[Deliver/%d] (from Channel) unknown FromMessage error: %d\n", no, message.From)
			}
		case message := <-pxToDv:
			switch message.From {
			case PROXY:
				fmt.Printf("[Deliver/%d] (PROXY) receive data\n", no)
				fmt.Printf("[Deliver/%d] receive channel: \n%s\n", no, hex.Dump(message.Buf))

				p := protocol.Packet{}
				if ret := p.Parsing(message.Buf, uint32(len(message.Buf))); ret {
					key, len := p.GetKey()
					fmt.Printf("[Deliver/%d] key:%s, len:%d\n", no, key, len)

					if tmp, ret := channelManager.Get(string(key)); ret {
						ch := tmp.(CHManage)
						ch.ToCH <- Message{
							From: PROXY,
							Buf:  message.Buf,
						}
					} else {
						fmt.Printf("[Deliver/%d] (PROXY) not found channel in map\n", no)
						continue
					}
				} else {
					fmt.Printf("[Deliver/%d] protocol parsing error\n", no)
					continue
				}
			default:
				fmt.Printf("[Deliver/%d] (from Proxy) unknown FromMessage error: %d\n", no, message.From)
			}
		}
	}
}
