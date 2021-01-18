package ioCommunication

import (
	"bufio"
	"fmt"
	"os"
)

var io *ioComm

type ioComm struct {
	writer *bufio.Writer
	reader *bufio.Reader
}

func Read(bufLength uint32) ([]byte, int, error) {
	buf := make([]byte, bufLength)
	len, err := io.reader.Read(buf)
	if err != nil {
		fmt.Println("[ioRead] error: ", err)
		return nil, 0, err
	} else {
		return buf, len, err
	}
}

func Write(buf []byte) (int, error) {
	if len, err := io.writer.Write(buf); err != nil {
		fmt.Println("[ioWrite] write error: ", err)
		return 0, err
	} else {
		io.writer.Flush()
		return len, nil
	}
}

func Done() {
	io.writer = nil
	io.reader = nil
	io = nil
}

func Init() {
	io = new(ioComm)

	io.writer = bufio.NewWriter(os.Stdout)
	io.reader = bufio.NewReader(os.Stdin)
}
