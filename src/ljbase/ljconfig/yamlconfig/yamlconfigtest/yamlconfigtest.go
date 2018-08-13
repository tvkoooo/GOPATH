package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

func main() {
	var c conf
	conf:=c.getConf()
	fmt.Println(conf.Host)
	fmt.Println(*conf)
}

//profile variables
type conf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}
func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("./ljbase/ljconfig/yamlconfig/yamlconfigtest/yamlconfig.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}