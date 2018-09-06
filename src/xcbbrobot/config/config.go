package config

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"xcbbrobot/common/typechange"
)

type XcbbRobotConfig struct {
	Loglevel int `yaml:"Loglevel"`
	Logfile string `yaml:"Logfile"`
	Robotlist string `yaml:"Robotlist"`
	Server string `yaml:"Server"`
}

type AppConfig struct {
	Instance string
	LogFilePath string
	LogLevel int
	ObjectNet string
	Yaml XcbbRobotConfig
}

var Conf *AppConfig

//config new
func AppConfigNew()(*AppConfig) {
	var a AppConfig
	Conf = &a
	return &a
}


//对已经创建的config进行初始化
func (a *AppConfig)AppConfigInit()()  {
	a.GetParameter()
	a.LoadConfig()
}
//获得程序输入参数
func (a *AppConfig)GetParameter()  {
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) <5{
		fmt.Println("参数不够，举个例子：")
		fmt.Println("xcbbrobot 0 \"../src/xcbbrobot/log/\" 4 \"59.110.125.134:30302\"")
		a.Instance = "0"
		a.LogFilePath = "../src/xcbbrobot/log/"
		a.LogLevel = 4
		a.ObjectNet = "59.110.125.134:30302"
	}else {
		a.Instance = args[1]
		a.LogFilePath = args[2]
		a.LogLevel = typechange.String2IntRe(args[3])
		a.ObjectNet = args[4]
	}
}
//加载外部配置参数
func (a *AppConfig)LoadConfig() {
	yamlFile, err := ioutil.ReadFile("../src/xcbbrobot/config/config.yaml")
	if err != nil {
		fmt.Println("加载配置文件 config.yaml 失败",err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &a.Yaml)
	if err != nil {
		fmt.Println("解析配置文件 config.yaml 失败",err.Error())
	}
}











