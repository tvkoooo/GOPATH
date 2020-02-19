//如何加载外部配置，需要给出一个配置路径初始化。
//本配置只支持二级加载   [test]something  = wrong
//可以直接通过 ljconfig.ConfigUnix 解析后 myConfig.Mymap 进行调用
package main

import (
	"fmt"
	"ljbase/ljconfig/unixconfig"
)

func main() {
	myConfig := new(ljconfig.ConfigUnix)
	myConfig.InitConfig("./ljbase/ljconfig/unixconfig/unixconfigtest/unixconfig.uni")
	default_path := myConfig.Read("default", "path")
	fmt.Println(default_path)
	default_version := myConfig.Read("default", "version")
	fmt.Println(default_version)
	test_num := myConfig.Read("test", "num")
	fmt.Println(test_num)
	//
	fmt.Println(myConfig.Mymap["test.something"])
	//
	fmt.Printf("%v \n", myConfig.Mymap)
	//
	fmt.Println(myConfig.Mymap)
}
