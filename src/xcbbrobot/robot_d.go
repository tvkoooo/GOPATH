package main

import (
	"xcbbrobot/robotapp"
)

const SYSTEM_TIME_5S  = 5E9


func main() {
	//【程序init】创建 app 结构
	var aP robotapp.AppProgram

	//【程序init】App 初始化配置
	aP.AppInit()

	//App connect
	aP.AppRobotConn()

	//系统做 定时 检测，如果数据槽还有残留数据，需要连接发送
	go aP.SystemSendMessage(SYSTEM_TIME_5S)

	go aP.RobCtrl.AppRobotOnline.SendPing()

	//循环数据槽解包,错误已经在接收消息包进行了处理
	for{

		recNum, err := aP.ReceiveMessage()
		if nil == err {
			if recNum > 0 {
				aP.DecodeMessage()
			}
		}
	}
}