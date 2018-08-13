package main

import (
	"time"

	"fmt"
	"xcbbrobot/config"
	"xcbbrobot/logfile"
	"xcbbrobot/robotcontrol"
)

func main() {
	//初始化配置
	config.Initconf()

	//初始化日志
	logfile.LogFileInit()
	logfile.GlobalLog.LogFileOpen(config.Conf.Logfile + "robotlog" + time.Now().Format("20060102") + ".log")
	defer logfile.GlobalLog.LogFileClosed()
	logfile.GlobalLog.SetLoglevel(config.Conf.Loglevel)

	var robctrl robotcontrol.RobotControl
	robctrl.RobotControlInit()
	fmt.Println("AppRobotLen", robctrl.AppRobotLen())
	fmt.Println("AppRoomLen", robctrl.AppRoomLen())
	robctrl.PrintAppRobot()
	go func() {
		for i := 0; i < 55; i++ {
			fmt.Println("i:", i)
			time.Sleep(1E9)
		}
	}()
	robctrl.RoomCreate(102692)
	go robctrl.RoomAddRobot(102692)
	time.Sleep(1E9)
	robctrl.PrintAppRobot()
	fmt.Println("AppRobotLen", robctrl.AppRobotLen())
	fmt.Println("AppRoomLen", robctrl.AppRoomLen())
	fmt.Println("RoomRobotLen", robctrl.RoomRobotLen(102692))
	go func() {
		time.Sleep(25E9)
		fmt.Println("RoomSubRobot now")
		robctrl.RoomSubRobot(102692)
	}()
	time.Sleep(40E9)
	fmt.Println("AppRobotLen", robctrl.AppRobotLen())
	fmt.Println("AppRoomLen", robctrl.AppRoomLen())
	fmt.Println("RoomRobotLen", robctrl.RoomRobotLen(102692))
	time.Sleep(1E9)
	robctrl.RoomDel(102692)
	time.Sleep(1E9)
	fmt.Println("AppRobotLen", robctrl.AppRobotLen())
	fmt.Println("AppRoomLen", robctrl.AppRoomLen())
	fmt.Println("RoomRobotLen", robctrl.RoomRobotLen(102692))
	time.Sleep(1E9)
}
