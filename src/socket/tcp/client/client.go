package main
import (
	"net"
	"os"
	_"time"
)

func main() {
	strEcho := "Halo"
	servAddr := "localhost:2219"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	//go func() {
	//	for {
			//var loop = 100
			//for i:=0; i<loop; i++ {
				_, err = conn.Write([]byte(strEcho))
				if err != nil {
					println("Write to server failed:", err.Error())
					os.Exit(1)
				}
			//}
			//time.Sleep(1000 * time.Microsecond)
//		}
//	}()

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	conn.Close()
}
