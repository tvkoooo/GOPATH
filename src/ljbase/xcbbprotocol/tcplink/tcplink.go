package tcplink

import (
	"fmt"
	"net"
	"os"
)

// Tcplink
func Tcplink(server string) (conn net.Conn) {
	//server := "127.0.0.1:9090"
	//server := "59.110.125.134:30302"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	return conn
}

//
func Tcplisten() (conn net.Conn) {

	//建立socket，监听端口
	netListen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer netListen.Close()

	fmt.Println("Waiting for clients")
	for {
		conn, err = netListen.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		return conn
	}
}
