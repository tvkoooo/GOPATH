package tcplink

import (
	"net"
	"gotest/common/logfile"
)

// Tcplink
func Tcplink(server string)(conn net.Conn){
	//server := "127.0.0.1:9090"
	//server := "59.110.125.134:30302"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		//【Warn】
		logfile.GlobalLog.Warnln("Tcplink::TCP link error:", err.Error(),"Please wait a little later")
		conn = nil
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		//【Warn】
		logfile.GlobalLog.Warnln("Tcplink::TCP link error:", err.Error(),"Please wait a little later")
		conn = nil
	}
	//fmt.Println("connect success:",server)
	return conn
}
//
func Tcplisten(server string)(conn net.Conn){
	//建立socket，监听端口
	//server := "localhost:9090"
	netListen, err := net.Listen("tcp", server)
	if err != nil {
		logfile.GlobalLog.Warnln("Tcplisten::TCP listen error:", err.Error(),"Please wait a little later")
		conn = nil
	}
	defer netListen.Close()

	//logfile.SystemLogPrintln("info","Waiting for clients")
	for {
		conn, err = netListen.Accept()
		if err != nil {
			continue
		}
		//logfile.SystemLogPrintln("info",conn.RemoteAddr().String(), " tcp connect success")
		return conn
	}
}



