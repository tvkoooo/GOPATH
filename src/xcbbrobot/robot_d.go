package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"xcbbrobot/common/fileopr"
	"xcbbrobot/config"
	"xcbbrobot/logfile"
	"xcbbrobot/robotapp"
)

//const SYSTEM_TIME_5S  = 5E9

func main() {
	//程序输入参数初始化
	confInit()

	//捕获正常退出
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)

		s := <-c
		fmt.Println("Got signal:", s)
		os.Exit(0)
	}()

	//【程序init】创建 app 结构
	var aP robotapp.AppProgram

	//【程序init】App 初始化配置
	aP.AppInit()

	//App connect
	for {
		if aP.AppRobotConn() {
			break
		}
	}

	//系统做 定时 检测，如果数据槽还有残留数据，需要连接发送
	//go aP.SystemSendMessage(SYSTEM_TIME_5S)

	go aP.RobCtrl.PollOrder()
	go aP.RobCtrl.PollReceive()

	//循环数据槽解包,错误已经在接收消息包进行了处理
	for {

		recNum, err := aP.ReceiveMessage()
		if nil == err {
			if recNum > 0 {
				aP.DecodeMessage()
			}
		} else {
			aP.RobCtrl.RoomClean()
			if aP.Conn == nil {
				for {
					if aP.AppRobotConn() {
						break
					}
				}
			}
		}
	}
	logfile.GlobalLog.LogFileClosed()
}

func confInit() {
	//【程序init】初始化配置
	config.AppConfigNew()
	config.Conf.AppConfigInit()

	//创建/打开 实例号文件夹
	dirPath := config.Conf.LogFilePath + "vnc_robot_" + config.Conf.Instance
	fmt.Println("dirPath", dirPath)
	fileopr.CreateDir(&dirPath)
	//创建/打开 日期文件夹
	var today string
	today = time.Now().Format("20060102")
	dir := dirPath + "/" + today
	fmt.Println("today", today, "dir", dir)
	fileopr.CreateDir(&dir)

	//【程序init】初始化日志文件
	logfile.LogFileNew()
	logfile.GlobalLog.SetLoglevel(config.Conf.LogLevel)
	logfile.GlobalLog.LogFileOpen(dir + "/" + "robot_d.log")

}
