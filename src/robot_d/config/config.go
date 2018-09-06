package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"robot_d/common/typechange"
	"robot_d/common/logfile"
	"strings"
)

type XcbbRobotConfig struct {
	Loglevel int `yaml:"Loglevel"`
	Logfile string `yaml:"Logfile"`
	Robotlist string `yaml:"Robotlist"`
	Server string `yaml:"Server"`
}

type AppConfig struct {
	LogFilePath string
	LogLevel int
	Instance string
	ServiceNum string
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
	if args == nil || len(args) <6{
		logfile.SystemLogPrintln("WARN","参数不够，举个例子: \n ./rc_robot_d /data/vnc_log/vnc/vnc_robot_2 6 2 3 172.16.39.170-30302")
		a.LogFilePath = "/home/longjia/go/log/vnc_robot_1"
		a.LogLevel = 7
		a.Instance = "1"
		a.ServiceNum = "1"
		getObjectNet := "172.16.39.170-30301"
		a.ObjectNet = strings.Replace(getObjectNet,"-",":",-1)
		logfile.SystemLogPrintln("info","\n","Set log path:",a.LogFilePath,"\n","Set log level:",a.LogLevel,"\n","Set Case number:",a.Instance,"\n","Set ServiceNum:",a.ServiceNum,"\n","Set Connect remote address:",a.ObjectNet)
	}else {
		a.LogFilePath = args[1]
		a.LogLevel = typechange.String2IntRe(args[2])
		a.Instance = args[3]
		a.ServiceNum = args[4]
		getObjectNet := args[5]
		a.ObjectNet = strings.Replace(getObjectNet,"-",":",-1)
		logfile.SystemLogPrintln("info","\n","Set log path:",a.LogFilePath,"\n","Set log level:",a.LogLevel,"\n","Set Case number:",a.Instance,"\n","Set ServiceNum:",a.ServiceNum,"\n","Set Connect remote address:",a.ObjectNet)
	}
	//拿到实例编号以后，给系统日志增加实例编号，如果不做添加，默认实例编号为空
	logfile.SystemLogSetInstance(a.Instance)
}
//加载外部配置参数
func (a *AppConfig)LoadConfig() {
	yamlFile, err := ioutil.ReadFile("../src/robot_d/config/config.yaml")
	if err != nil {
		logfile.SystemLogPrintln("ERROR","加载配置文件 config.yaml 失败",err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &a.Yaml)
	if err != nil {
		logfile.SystemLogPrintln("ERROR","解析配置文件 config.yaml 失败",err.Error())
	}
}











