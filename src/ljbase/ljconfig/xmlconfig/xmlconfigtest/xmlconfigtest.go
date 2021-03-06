//如何加载外部配置，需要给出一个配置路径初始化。
//本配置要求外部json格式，go内部也有相应结构体
//可以直接通过 ljconfig.ConfigJson 解析后 myConfig 进行调用
package main

import (
	"fmt"
	"ljbase/ljconfig/xmlconfig"
)

func main() {
	myConfig := new(ljconfig.ConfigXml)
	myConfig.InitConfig("./ljbase/ljconfig/xmlconfig/xmlconfigtest/xmlconfig.xml")
	//

	fmt.Println(myConfig.Enabled)
	fmt.Println(myConfig.Path)
	fmt.Println(*myConfig)
	fmt.Printf("%+v\n", *myConfig)
	fmt.Printf("%#v\n", *myConfig)
}
