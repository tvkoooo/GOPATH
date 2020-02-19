package main

import (
	"fmt"
	"ljnew/loglj"
	"time"
)

func main() {

	logwork, err := loglj.LogFileOpen("D:/GOPATH/resources/ljnew/loglj/" + time.Now().Format("20060102") + ".log")
	if err != nil {
		fmt.Println("日志文件操作失败", err)
	}
	defer loglj.LogFileClosed(logwork)
	loglj.SetLoglevel(logwork, loglj.L_INFO)

	loglj.Debugln(logwork, "2w2w2w22w2w2w2w2w2w2w2w2")

	logdebug, err := loglj.LogFileOpen("D:/GOPATH/resources/ljnew/loglj/" + time.Now().Format("20060102") + ".debug")
	if err != nil {
		fmt.Println("日志文件操作失败", err)
	}
	defer loglj.LogFileClosed(logdebug)
	loglj.SetLoglevel(logdebug, loglj.L_DEBUG)
	loglj.Debugln(logdebug, "90909090909090909090900")

	fmt.Println("logworklevel", loglj.GetLoglevel(logwork))
	fmt.Println("logdebuglevel", loglj.GetLoglevel(logdebug))

	loglj.Debugln(logdebug, "111111111111111111")
	loglj.Debugln(logwork, "aaaaaaaaaaaaaa")
	loglj.Infoln(logdebug, "2222222222222222222")
	loglj.Infoln(logwork, "bbbbbbbbbbbbbb")
	loglj.Errorln(logdebug, "3333333333333333")
	loglj.Errorln(logwork, "ccccccccccccccccccc")
	loglj.Fatalln(logdebug, "44444444444444444444")
	loglj.Fatalln(logwork, "ddddddddddddddddd")
}
