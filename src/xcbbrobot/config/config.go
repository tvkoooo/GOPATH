package config

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

var Conf *XcbbRobotConfig

func Initconf() {
	var c XcbbRobotConfig
	yamlFile, err := ioutil.ReadFile("./xcbbrobot/config/config.yaml")
	if err != nil {
		fmt.Println("加载配置文件 config.yaml 失败",err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println("解析配置文件 config.yaml 失败",err.Error())
	}
	Conf = &c
}











