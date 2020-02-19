package main

import (
	"fmt"
	"xcbbrobot/config"
)

func main() {
	var c config.XcbbRobotConfig
	c.Initconf()
	fmt.Println(c.Logfile)
	fmt.Println(c.Loglevel)
}
