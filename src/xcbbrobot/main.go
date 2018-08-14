package main

import (
	"xcbbrobot/config"
	"xcbbrobot/logfile"
	"time"
	"xcbbrobot/common/tcplink"
	"xcbbrobot/common/message"
	"xcbbrobot/common/datagroove"
	"fmt"
	"xcbbrobot/robotcontrol"
	"xcbbrobot/common/maths"
	"os"
	"path/filepath"
	"log"
)

func main()  {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	//初始化配置
	config.Initconf()

	//初始化日志
	logfile.LogFileInit()
	logfile.GlobalLog.LogFileOpen(config.Conf.Logfile + "robotlog" + time.Now().Format("20060102") + ".log")
	defer logfile.GlobalLog.LogFileClosed()
	logfile.GlobalLog.SetLoglevel(config.Conf.Loglevel)

	//启动机器人控制初始化
	var robctrl robotcontrol.RobotControl
	robctrl.RobotControlInit()

	//连接tcp
	conn := tcplink.Tcplink(config.Conf.Server)

	//注册机器人程序
	sendRobotCon := message.SendPRegisteredPI()
	numSend , err := conn.Write(sendRobotCon)
	fmt.Println("send message:", sendRobotCon)
	if nil != err || numSend != 23{
		logfile.GlobalLog.Fatalln("注册机器人失败")
	}



	//等待房间服给出机器人执行命令
	//开辟一个数据接收槽
	var dat datagroove.DataBuff
	dat.BufferInit()
	recData := make([]byte, 1024)
	for {
		count, err := conn.Read(recData)
		if err != nil {
			fmt.Println("Fatal error:" , err.Error())
			break
		}

		if count != 0 {
			//fmt.Println("rec Binary stream", recData[:count])
			dat.DataAppend(recData[:count])
			//fmt.Println("append dat ", dat)

			popData := dat.MessagePop()
			//fmt.Println("after pop dat", dat,"popData dat", popData)
			for ;nil != popData ;{

				cmdRobot := uint32((131 << 8) | 2 )
				if message.CheckPeakHead(popData, cmdRobot){
					lengrs , ph , robotctl := message.ReceivePRobotServerCmd(popData)
					fmt.Println("房间发送数据num: ",lengrs ,"房间发送数据 包头: ",ph ,"房间发送数据 内容: ",robotctl)

					switch robotctl.Cmd {
					case 0:
						fmt.Println("房间 START",robotctl.Sid)
						//创建房间
						robctrl.RoomCreate(robotctl.Sid)

						// 随机时间(0.5s-1s 增加一个) 共 3个机器人
						go func() {
							for iii:=0; iii<3;iii++  {
								time.Sleep(time.Duration(1E7 * maths.BetweenRand(50,100)))
								go robctrl.RoomAddRobot(robotctl.Sid)
							}
						}()
					case 1:
						fmt.Println("房间 STOP")
						//把房间里面的机器人全部移除
						numSid := robctrl.RoomRobotLen(robotctl.Sid)
						for iii:=0; iii<numSid;iii++ {
							go robctrl.RoomSubRobot(robotctl.Sid)
						}
						//给个1延时操作，等待机器人全部移除后关闭房间
						go func() {
							time.Sleep(1E9)
							robctrl.RoomDel(robotctl.Sid)
						}()
					case 2:
						fmt.Println("目前没有用")
					case 3:
						fmt.Println("用户进场 ADD_ROBOT")
						//增加一个机器人
						go robctrl.RoomAddRobot(robotctl.Sid)

					case 4:
						fmt.Println("用户出场 REMOVE_ROBOT")
						//减少一个机器人
						numSid := robctrl.RoomRobotLen(robotctl.Sid)
						if numSid>0 {
							go robctrl.RoomSubRobot(robotctl.Sid)
						}
					default:
						fmt.Println("Default")
					}

				}
				popData = dat.MessagePop()
				//fmt.Println("after pop dat", dat,"popData dat", popData)
			}
		}
	}
}

