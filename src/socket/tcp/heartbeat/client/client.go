package main
import (
	"net"
	"os"
	"time"
	"syscall"
	"log"
)

func setKeepaliveParameters(conn *net.TCPConn) {
    rawConn, err := conn.SyscallConn()
    if err != nil {
        log.Println("on getting raw connection object for keepalive parameter setting", err.Error())
    }

    rawConn.Control(
        func(fdPtr uintptr) {
            // got socket file descriptor. Setting parameters.
            fd := int(fdPtr)
            //Number of probes.
            err := syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPCNT, 3)
            if err != nil {
                log.Println("on setting keepalive probe count", err.Error())
            }
            //Wait time after an unsuccessful probe.
            err = syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL, 3)
            if err != nil {
                log.Println("on setting keepalive retry interval", err.Error())
            }
	})
}

func main() {
	strEcho := "Halo"
	servAddr := "localhost:1234"
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

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Second * 3)
	setKeepaliveParameters(conn)

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

	for {

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

		time.Sleep(3 * time.Second)
	}

	conn.Close()
}
