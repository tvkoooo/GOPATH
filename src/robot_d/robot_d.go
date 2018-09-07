package main

import (
	"robot_d/robotapp"
	"robot_d/config"
	"robot_d/common/logfile"
	"time"
	"os"
	"os/signal"
	"robot_d/common/fileopr"
)

func main() {

	//设定系统日志路径
	logfile.SystemLogSetDefaultPath()
	logfile.SystemLogPrintln("SYSTEM","Program startup==>>:",logfile.FilePath)

	//程序输入参数初始化
	confInit()

	//【程序init】创建 app 结构
	var aP robotapp.AppProgram

	//【程序init】App 初始化配置
	aP.AppInit()

	//App connect
	for  {
		if aP.AppRobotConn() {
			break
		}
	}

	go aP.RobCtrl.PollOrder()
	go aP.RobCtrl.PollReceive()

	//捕获正常退出
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)

		s := <-c
		logfile.SystemLogPrintln("SYSTEM","Program exit<<==:", logfile.FilePath,"Got signal:", s)

		aP.RobCtrl.RoomClean()
		logfile.GlobalLog.LogFileClosed()
		time.Sleep(5E6)
		os.Exit(0)
	}()

	//循环数据槽解包,错误已经在接收消息包进行了处理
	for{

		recNum, err := aP.ReceiveMessage()
		if nil == err {
			if recNum > 0 {
				aP.DecodeMessage()
			}
		}else {
			aP.RobCtrl.RoomClean()
			if aP.Conn == nil {
				for  {
					if aP.AppRobotConn() {
						break
					}
				}
			}
		}
	}
	logfile.GlobalLog.LogFileClosed()
}

func confInit()  {

	//【程序init】初始化配置
	config.AppConfigNew()
	config.Conf.AppConfigInit()

	//创建/打开 实例号文件夹
	dirPath :=config.Conf.LogFilePath
	logfile.SystemLogPrintln("info","confInit::Function log path:",dirPath)
	fileopr.CreateDir(&dirPath)

	//【程序init】初始化日志文件
	logfile.LogFileNew()
	logfile.GlobalLog.SetLogLevel(config.Conf.LogLevel)
	logfile.GlobalLog.StartLogFile(dirPath +"/"+ "go_robot_d.log")
}

