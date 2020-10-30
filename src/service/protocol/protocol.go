package protocol

//package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	COMMAND    = iota
	SUBCOMMAND = iota
	PID        = iota
	KEYLEN     = iota
	KEY        = iota
	DATALEN    = iota
	DATA       = iota
)

type Head struct {
	command    string // main command
	subCommand string // sub command
	pid        uint32 // packet unique id
	keyLen     uint8  // packet unique key length
	key        []byte // packet unqiue key
}

type Body struct {
	dLen uint32 // data length
	data []byte // data
}

type Packet struct {
	tBuf []byte
	tLen uint32

	head Head
	body Body
}

func (p *Packet) makeProtocol() ([]byte, uint32) {
	buf := new(bytes.Buffer)

	// command|subCommand|pid|keyLen|key|dLen|data|
	// SMS|SEND|00 00 00 01|05|51010|00 00 00 05|HELLO|

	buf.Write([]byte(p.head.command))
	buf.Write([]byte("|"))
	buf.Write([]byte(p.head.subCommand))
	buf.Write([]byte("|"))
	if err := binary.Write(buf, binary.BigEndian, p.head.pid); err != nil {
		fmt.Println("[protocol/make] make pid error: ", err)
		return nil, 0
	}
	buf.Write([]byte("|"))
	if err := binary.Write(buf, binary.BigEndian, p.head.keyLen); err != nil {
		fmt.Println("[protocol/make] make key length error: ", err)
		return nil, 0
	}
	buf.Write([]byte("|"))
	if p.head.keyLen > 0 {
		buf.Write(p.head.key)
	}
	buf.Write([]byte("|"))

	if err := binary.Write(buf, binary.BigEndian, p.body.dLen); err != nil {
		fmt.Println("[protocol/make] make key length error: ", err)
		return nil, 0
	}
	buf.Write([]byte("|"))
	if p.body.dLen > 0 {
		buf.Write(p.body.data)
	}
	buf.Write([]byte("|"))

	p.tBuf = buf.Bytes()
	p.tLen = uint32(buf.Len())

	return p.tBuf, p.tLen
}

func (p *Packet) Make(command string, subCommand string, pid uint32, keyLen uint8, key []byte, dLen uint32, data []byte) ([]byte, uint32) {
	if 0 >= len(command) || 0 >= len(subCommand) {
		return nil, 0
	}
	if 0 >= keyLen && len(key) > 0 {
		return nil, 0
	}
	if 0 >= dLen && len(data) > 0 {
		return nil, 0
	}

	p.head.command = command
	p.head.subCommand = subCommand
	p.head.pid = pid
	p.head.keyLen = keyLen
	p.head.key = key[:p.head.keyLen]

	p.body.dLen = dLen
	p.body.data = data[:p.body.dLen]

	buf, len := p.makeProtocol()

	return buf, len
}

/*****************************************************************/

func (p *Packet) parsing() bool {
	tmp := p.tBuf
	var x int

	idx := 0
	for x = bytes.Index(tmp, []byte("|")); x != -1; x = bytes.Index(tmp, []byte("|")) {
		val := tmp[:x]
		fmt.Println(val)
		switch idx {
		case COMMAND:
			p.head.command = string(val)
		case SUBCOMMAND:
			p.head.subCommand = string(val)
		case PID:
			p.head.pid = binary.BigEndian.Uint32(val)
		case KEYLEN:
			p.head.keyLen = val[0]
		case KEY:
			p.head.key = val
		case DATALEN:
			p.body.dLen = binary.BigEndian.Uint32(val)
		case DATA:
			p.body.data = val
		default:
			fmt.Println("[protocol/parsing] unknown data received error")
			return false
		}

		tmp = tmp[x+1:]
		idx++
	}

	fmt.Printf("[protocol/parsing] [head] command    : %s\n", p.head.command)
	fmt.Printf("[protocol/parsing] [head] sub command: %s\n", p.head.subCommand)
	fmt.Printf("[protocol/parsing] [head] pid        : %d\n", p.head.pid)
	fmt.Printf("[protocol/parsing] [head] key length : %d\n", p.head.keyLen)
	fmt.Printf("[protocol/parsing] [head] key        : %s\n", p.head.key)
	fmt.Printf("[protocol/parsing] [body] data length: %d\n", p.body.dLen)
	fmt.Printf("[protocol/parsing] [body] data       : %s\n", p.body.data)

	return true
}

func (p *Packet) Parsing(tBuf []byte, tLen uint32) bool {
	if 0 >= len(tBuf) || 0 >= tLen {
		return false
	}

	p.tBuf = tBuf
	p.tLen = tLen

	p.parsing()
	fmt.Println(p)

	return true
}

func (p *Packet) GetCommand() string {
	return p.head.command
}

func (p *Packet) GetSubCommand() string {
	return p.head.subCommand
}

func (p *Packet) GetPid() uint32 {
	return p.head.pid
}

func (p *Packet) GetKey() ([]byte, uint8) {
	return p.head.key, p.head.keyLen
}

func (p *Packet) GetData() ([]byte, uint32) {
	return p.body.data, p.body.dLen
}

/*****************************************************************/

//func main() {
//	p := Packet{}
//
//	buf, len := p.Make("SMS", "SEND", 1, 3, []byte("510"), 0, nil)
//	fmt.Println(len)
//	fmt.Printf("%s\n", hex.Dump(buf))
//
//	//buf := []byte("SMS|SEND|")
//	//len := uint32(len(buf))
//
//	p.Parsing(buf, len)
//
//	fmt.Println(p.GetCommand())
//	fmt.Println(p.GetSubCommand())
//	fmt.Println(p.GetPid())
//	fmt.Println(p.GetKey())
//	fmt.Println(p.GetData())
//}
//
