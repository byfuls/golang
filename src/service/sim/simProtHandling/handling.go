package simProtHandling

import (
	"bytes"
	"encoding/binary"

	"service/sim/clientHandling"
	"module/logging"
)

/*************************************************************************************/
func Recv_AS90(pdata *parsedData) bool {
	logging.DebugF("[Recv_AS90] start, imsi[%d:%s]", len(pdata.Head.Key), pdata.Head.Key)

	if ret := clientHandling.Update(pdata.Head.Key); !ret {
		logging.ErrorLn("[Recv_AS90] update error, imsi: ", pdata.Head.Key)
		return false
	}

	return true
}

/*************************************************************************************/
func Send_AS03(pdata *parsedData) ([]byte, int) {
	logging.DebugLn("[Send_AS03] start")

	/* get mapped imsi & client info via gateway */
	if ret, addr, gateway := clientHandling.GetMappedAddressNGatewayViaImsi(pdata.Head.Key); !ret {
		logging.ErrorLn("[Send_AS03] get mapped client info error, imsi: ", pdata.Head.Key)
		return nil, -1
	} else {
		pdata.Addr = addr
		pdata.Head.Key = gateway
	}

	buf := new(bytes.Buffer)
	protocols := [][]byte{
		[]byte("AS03"),
		[]byte(pdata.Head.Key),
		{0x00, 0x00, 0x00, 0x00}, // Seq, not used(?)
		{0x00, 0x00, 0x00, 0x00}, // Rev, not used(?)
		pdata.Body,
	}
	for _, b := range protocols {
		if err := binary.Write(buf, binary.BigEndian, b); err != nil {
			logging.ErrorLn("[Send_AS03] make protocol error")
			return nil, -1
		}
	}

	return buf.Bytes(), buf.Len()
}

func Recv_AS03(pdata *parsedData) bool {
	logging.DebugLn("[Recv_AS03] start")

	if ret := clientHandling.Update(pdata.Head.Key); !ret {
		logging.ErrorLn("[Recv_AS03] update error, imsi: ", pdata.Head.Key)
		return false
	}

	return true
}

func Recv_AA00(pdata *parsedData) bool {
	logging.DebugLn("[Recv_AA00] start")

	if ret := clientHandling.Update(pdata.Head.Key); !ret {
		logging.ErrorLn("[Recv_AA00] update error, imsi: ", pdata.Head.Key)
		return false
	}
	return true
}

func Resp_AA00(pdata *parsedData) ([]byte, int) {
	logging.DebugLn("[Resp_AA00] start")

	buf := new(bytes.Buffer)
	protocols := [][]byte{
		[]byte("AA00"),
		[]byte(pdata.Head.Key),
		{0x00, 0x00, 0x00, 0x00}, // Seq, not used(?)
		{0x00, 0x00, 0x00, 0x00}, // Rev, not used(?)
	}
	for _, b := range protocols {
		if err := binary.Write(buf, binary.BigEndian, b); err != nil {
			logging.ErrorLn("[Resp_AA00] make protocol error")
			return nil, -1
		}
	}

	return buf.Bytes(), buf.Len()
}

func Send_AC03(pdata *parsedData) ([]byte, int) {
	logging.DebugLn("[Send_AC03] start")

	/* get mapped imsi & client info via gateway */
	if ret, addr, imsi := clientHandling.GetMappedAddressNImsiViaGateway(pdata.Head.Key); !ret {
		logging.ErrorLn("[Send_AC03] get mapped client info error, gateway: ", pdata.Head.Key)
		return nil, -1
	} else {
		pdata.Addr = addr
		pdata.Head.Key = imsi
	}

	buf := new(bytes.Buffer)
	protocols := [][]byte{
		[]byte("AC03"),
		[]byte(pdata.Head.Key),
		{0x00, 0x00, 0x00, 0x00}, // Seq, not used(?)
		{0x00, 0x00, 0x00, 0x00}, // Rev, not used(?)
		pdata.Body,
	}
	for _, b := range protocols {
		if err := binary.Write(buf, binary.BigEndian, b); err != nil {
			logging.ErrorLn("[Send_AC03] make protocol error")
			return nil, -1
		}
	}

	return buf.Bytes(), buf.Len()
}

func Recv_AC03(pdata *parsedData) bool {
	logging.DebugLn("[Recv_AC03] start")

	return true
}

/*************************************************************************************/
func Recv_MA06(pdata *parsedData) bool {
	logging.DebugLn("[Recv_MA06] start")
	tmp := pdata.Body
	var imsi string
	var seq int32
	var gatewayId string
	i := 0
	for f := bytes.Index(tmp, []byte("|")); f != -1; f = bytes.Index(tmp, []byte("|")) {
		switch i {
		case 0:
			imsi = string(tmp[:f])
		case 1:
			seq = int32(binary.BigEndian.Uint32(tmp[:f]))
		case 2:
			gatewayId = string(tmp[:f])
		}
		i++
		tmp = tmp[f+1:]
	}

	logging.DebugLn("[Recv_MA06] imsi[%s] seq(%d) gatewayId[%s]\n", imsi, seq, gatewayId)

	if ret := clientHandling.Mapping(pdata.Addr, gatewayId, pdata.Head.Key); ret == false {
		logging.ErrorF("[Recv_MA06] mapping error, gateway[%s] imsi[%s]\n", gatewayId, pdata.Head.Key)
		return false
	}

	return true
}

/*************************************************************************************/
func Send_LUR(pdata *parsedData) ([]byte, int) {
	logging.DebugLn("[Send_LUR] start")

	buf := new(bytes.Buffer)
	protocols := [][]byte{
		[]byte("LUR"),
		[]byte("|"),
		[]byte("TRY"),
		[]byte("|"),
		[]byte(pdata.Head.Key),
		[]byte("|"),
	}
	for _, b := range protocols {
		if err := binary.Write(buf, binary.BigEndian, b); err != nil {
			logging.ErrorLn("[Send_LUR] make protocol error")
			return nil, -1
		}
	}

	return buf.Bytes(), buf.Len()
}

func Recv_AS07(pdata *parsedData) bool {
	logging.DebugF("[Recv_AS07] start, imsi[%d:%s]", len(pdata.Head.Key), pdata.Head.Key)
	if ret := clientHandling.Insert(pdata.Addr, pdata.Head.Key, pdata.Head.Seq); ret == false {
		logging.ErrorLn("[AS07] insert error")
		return false
	}

	return true
}

func Resp_AS07(pdata *parsedData) ([]byte, int) {
	logging.DebugLn("[Resp_AS07] start")

	lurInterval := 60

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, []byte(pdata.Head.Command)); err != nil {
		logging.ErrorLn("[Resp_AS07] error: ", err)
		return nil, -1
	}
	if err := binary.Write(buf, binary.BigEndian, []byte(pdata.Head.Key)); err != nil {
		logging.ErrorLn("[Resp_AS07] error: ", err)
		return nil, -2
	}
	if err := binary.Write(buf, binary.BigEndian, pdata.Head.Seq); err != nil {
		logging.ErrorLn("[Resp_AS07] error: ", err)
		return nil, -3
	}
	if err := binary.Write(buf, binary.BigEndian, pdata.Head.Rev); err != nil {
		logging.ErrorLn("[Resp_AS07] error: ", err)
		return nil, -4
	}
	if err := binary.Write(buf, binary.BigEndian, uint32(lurInterval)); err != nil {
		logging.ErrorLn("[Resp_AS07] error: ", err)
		return nil, -5
	}

	logging.DebugLn(buf.Bytes())
	logging.DebugLn(buf.Len())

	return buf.Bytes(), buf.Len()
}
