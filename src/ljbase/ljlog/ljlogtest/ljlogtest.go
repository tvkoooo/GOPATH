package main

import (
	"time"
	"fmt"
	"ljbase/ljlog"
)

func main()  {

	logwork ,err :=ljlog.LogFileOpen("./ljbase/ljlog/ljlogtest/"+time.Now().Format("20060102")+".log")
	if err != nil {
		fmt.Println("日志文件操作失败", err)
	}
	defer ljlog.LogFileClosed(logwork)
	ljlog.SetLoglevel(logwork ,ljlog.L_INFO)


	ljlog.Debugln(logwork,"2w2w2w22w2w2w2w2w2w2w2w2")


	logerror ,err :=ljlog.LogFileOpen("./ljbase/ljlog/ljlogtest/"+time.Now().Format("20060102")+".err")
	if err != nil {
		fmt.Println("日志文件操作失败", err)
	}
	defer ljlog.LogFileClosed(logerror)
	ljlog.SetLoglevel(logerror ,ljlog.L_ERROR)
	ljlog.Debugln(logerror,"90909090909090909090900")


	fmt.Println("logworklevel", ljlog.GetLoglevel(logwork))
	fmt.Println("logerrorlevel", ljlog.GetLoglevel(logerror))

	ljlog.Debugln(logerror,"111111111111111111")
	ljlog.Debugln(logwork,"aaaaaaaaaaaaaa")
	ljlog.Infoln(logerror,"2222222222222222222")
	ljlog.Infoln(logwork,"bbbbbbbbbbbbbb")
	ljlog.Errorln(logerror,"3333333333333333")
	ljlog.Errorln(logwork,"ccccccccccccccccccc")
	ljlog.Fatalln(logerror,"44444444444444444444")
	ljlog.Fatalln(logwork,"ddddddddddddddddd")
}















