package simProtHandling

import (
	"encoding/binary"
	"net"

	"module/logging"
)

type header struct {
	Command string
	Key     string // IMSI or GATEWAY
	Seq     uint32
	Rev     uint32
}

type parsedData struct {
	Addr net.UDPAddr
	Head header
	Body []byte
}

const (
	HEADER  = 32
	COMMAND = 4
	KEY     = 20
	SEQ     = 4
	REV     = 4
)

func Parsing(buf []byte) (bool, parsedData) {
	var pdata parsedData

	if 0 >= len(buf) {
		return false, pdata
	}

	pos := 0
	pdata.Head.Command = string(buf[0:COMMAND])
	pos += COMMAND
	pdata.Head.Key = string(buf[pos : pos+KEY])
	pos += KEY
	pdata.Head.Seq = binary.BigEndian.Uint32(buf[pos : pos+SEQ])
	pos += SEQ
	pdata.Head.Rev = binary.BigEndian.Uint32(buf[pos : pos+REV])
	pos += REV

	switch pdata.Head.Command {
	case "AS07", "AA00", "AS03", "AS90":
		break
	case "MA06", "AC03":
		pdata.Body = buf[pos:]
	default:
		logging.ErrorLn("[parsing] unknown command received, command: ", pdata.Head.Command)
		return false, pdata
	}

	logging.DebugF("[parsing] command: %s\n", pdata.Head.Command)
	logging.DebugF("[parsing] key: %s\n", pdata.Head.Key)
	logging.DebugF("[parsing] seq: %d\n", pdata.Head.Seq)
	logging.DebugF("[parsing] rev: %d\n", pdata.Head.Rev)

	return true, pdata
}
