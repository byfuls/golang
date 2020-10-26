package simProtHandling

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"service/sim/clientHandling"
)

func Recv_AS07(pdata *parsedData) bool {
	fmt.Println("[Recv_AS07] start")
	fmt.Println(pdata)
	if ret := clientHandling.Insert(pdata.Addr, pdata.Head.Imsi, pdata.Head.Seq); ret == false {
		fmt.Println("[AS07] insert error")
		return false
	}

	return true
}

func Resp_AS07(pdata *parsedData) ([]byte, int) {
	fmt.Println("[Resp_AS07] start")

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, []byte(pdata.Head.Command)); err != nil {
		fmt.Println("[Resp_AS07] error: ", err)
		return nil, -1
	}
	if err := binary.Write(buf, binary.BigEndian, []byte(pdata.Head.Imsi)); err != nil {
		fmt.Println("[Resp_AS07] error: ", err)
		return nil, -2
	}
	if err := binary.Write(buf, binary.BigEndian, pdata.Head.Seq); err != nil {
		fmt.Println("[Resp_AS07] error: ", err)
		return nil, -3
	}
	if err := binary.Write(buf, binary.BigEndian, pdata.Head.Rev); err != nil {
		fmt.Println("[Resp_AS07] error: ", err)
		return nil, -4
	}

	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())

	return buf.Bytes(), buf.Len()
}
