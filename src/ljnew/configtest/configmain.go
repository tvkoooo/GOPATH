package main

import (
	"fmt"
	"ljnew/configtest/configtest"

)

func main() {
	myConfig := new(configtest.Config)
	myConfig.InitConfig("D:/GOPATH/src/ljnew/configtest/configfile.cfg")
	//fmt.Println(myConfig.Read("default", "path"))

	default_path :=myConfig.Read("default", "path")
	fmt.Println(default_path)
	default_version  :=myConfig.Read("default", "version")
	fmt.Println(default_version)

	fmt.Printf("%v \n", myConfig.Mymap)

}