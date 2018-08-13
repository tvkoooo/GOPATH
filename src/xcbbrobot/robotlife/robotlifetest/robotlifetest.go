package main

import (
	"time"
	"xcbbrobot/robotlife"
	"xcbbrobot/config"
	"xcbbrobot/logfile"
	"fmt"
)

func main()  {

	//初始化配置
	config.Initconf()

	//初始化日志
	logfile.LogFileInit()
	logfile.GlobalLog.LogFileOpen(config.Conf.Logfile + "robotlog" + time.Now().Format("20060102") + ".log")
	defer logfile.GlobalLog.LogFileClosed()
	logfile.GlobalLog.SetLoglevel(config.Conf.Loglevel)
	go func() {
		for i := 0; i < 40; i++ {
			fmt.Println("i:", i)
			time.Sleep(1E9)
		}
	}()
	var(
		r1 robotlife.RobotLife
		//r2 robotlife.RobotSeed
		//r3 robotlife.RobotSeed
	)

	r1.RobotInit(10005259,102692)
	go r1.RobotWork()
	//go r2.RobotWork(10005127,102692)
	//go r3.RobotWork(10005131,102692)

	go func(p *robotlife.RobotLife) {
		time.Sleep(30E9)
		p.RobotRest()
	}(&r1)

	//go func(p *robotlife.RobotSeed) {
	//	time.Sleep(4E9)
	//	p.RobotRest()
	//}(&r2)

	//go func(p *robotlife.RobotSeed) {
	//	time.Sleep(7E9)
	//	p.RobotRest()
	//}(&r3)

	time.Sleep(40E9)
}

