package main

import (
	"fmt"
	"gotest/common/datagroove"
	"gotest/common/logfile"
	"gotest/robotvector"
	"net"
	"os"
	"time"
)

func main() {
	logfile.SystemLogSetDefaultPath()

	args := os.Args //获取用户输入的所有参数

	var m robotvector.MapAppRobot
	m.RobotFreeInit()
	m.LoadRobot("../src/gotest/robotvector/robot_1.list")

	var zBuff datagroove.DataBuff

	zBuff.BufferInit()
	//message.WritePRegisteredPI(&zBuff)
	//
	var netPath string
	if len(args) == 1 {
		netPath = "59.110.125.134:30302"
	} else {
		netPath = args[1]
	}
	//
	//连接tcp
	//connZ := tcplink.Tcplink(netPath)
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", netPath)
	connZ, _ := net.DialTCP("tcp", nil, tcpAddr)

	//connZ.Write(zBuff.SGroove[zBuff.LenRemove : zBuff.LenRemove+zBuff.LenData])
	//
	//
	//go func() {
	//	data := make([]byte,1024*3)
	//	connZ.SetReadDeadline(time.Now().Add(1E3))
	//	connZ.Read(data)
	//}()
	//
	fmt.Println(&connZ)
	time.Sleep(40E9)

	//for  {
	//	num := m.Len()
	//	if num >0 {
	//		id :=m.PopRobot()
	//		logfile.SystemLogPrintln("id",id)
	//		//连接tcp
	//		conn := tcplink.Tcplink(netPath)
	//
	//		var sendBuff datagroove.DataBuff
	//		sendBuff.BufferInit()
	//		message.WritePRealJoinChannelBuff(&sendBuff ,id ,102692 )
	//		conn.Write(sendBuff.SGroove[sendBuff.LenRemove : sendBuff.LenRemove+sendBuff.LenData])
	//	}
	//
	//	time.Sleep(3E9)
	//}

}
