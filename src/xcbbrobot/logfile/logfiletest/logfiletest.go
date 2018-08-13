package main

import (
	"xcbbrobot/logfile"
	"xcbbrobot/config"
	"time"
)

func main()  {

	//初始化配置
	config.Initconf()

	//init logfile
	var l logfile.LogFileName
	l.LogFileOpen(config.Conf.Logfile + "robotlog" + time.Now().Format("20060102") + ".log")
	defer l.LogFileClosed()
	l.SetLoglevel(config.Conf.Loglevel)

	//
	l.Debugln("2w2w2w22w2w2w2w2w2w2w2w2")
	l.Debugln("aaaaaaaaaaaaaa")
	l.Infoln("bbbbbbbbbbbbbb")
	l.Errorln("ccccccccccccccccccc")
	l.Fatalln("ddddddddddddddddd")
}















