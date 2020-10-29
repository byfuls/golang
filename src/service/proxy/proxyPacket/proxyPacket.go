package proxyPacket

import (
	"fmt"
)

type head struct {
	command    string
	subCommand string
	bodyLength uint32
}

type code struct {
	mcc  string
	mnc  string
	code string
}

type Packet struct {
	Head head
	Code code

	buf         []byte
	totalLength uint32
}

func (p *Packet) Buf(buf []byte, tlen uint32) bool {
	if 0 >= len(buf) || 0 >= tlen {
		return false
	}

	fmt.Println("[proxyPacket] ", tlen, buf)

	p.buf = buf
	p.totalLength = tlen

	return true
}
