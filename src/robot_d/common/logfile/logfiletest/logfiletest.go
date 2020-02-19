package main

import (
	"fmt"
	"os"
	"robot_d/common/logfile"
	"robot_d/config"
	"time"
)

func main() {

	//【程序init】初始化配置
	config.AppConfigNew()
	config.Conf.AppConfigInit()

	//【程序init】初始化日志
	logfile.LogFileNew()
	var nowdate string
	nowdate = time.Now().Format("20060102")
	pathsss := config.Conf.LogFilePath + nowdate + "/" + "robot_d.log"

	_dir := config.Conf.LogFilePath + nowdate
	// 创建文件夹
	err := os.Mkdir(_dir, os.ModePerm)
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
	}

	fmt.Println("pathsss", pathsss)
	logfile.GlobalLog.StartLogFile(pathsss)
	defer logfile.GlobalLog.LogFileClosed()
	logfile.GlobalLog.SetLogLevel(config.Conf.LogLevel)

	//
	logfile.GlobalLog.Debugln("2w2w2w22w2w2w2w2w2w2w2w2")
	logfile.GlobalLog.Debugln("aaaaaaaaaaaaaa")
	logfile.GlobalLog.Infoln("bbbbbbbbbbbbbb")
	logfile.GlobalLog.Errorln("ccccccccccccccccccc")
	logfile.GlobalLog.Fatalln("ddddddddddddddddd")

	//打开本地文件 读取出全部数据
	fin, err := os.Open(config.Conf.LogFilePath + "robot_d_" + time.Now().Format("20060102") + ".log")
	defer fin.Close()
	if err != nil {
		return
	}
	fmt.Println(fin.Seek(0, os.SEEK_END))

	//文件指针指向文件末尾 获取文件大小保存于buf_len
	buf_len, _ := fin.Seek(0, os.SEEK_END)
	fmt.Println("buf_len", buf_len)
	//获取buf_len后把文件指针重新定位于文件开始
	fin.Seek(0, os.SEEK_SET)

	buf := make([]byte, buf_len)
	fin.Read(buf)
	fmt.Println(string(buf[:]), len(buf))

}
