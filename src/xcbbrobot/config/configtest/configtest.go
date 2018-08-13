package main

import (
	"xcbbrobot/config"
	"fmt"
)

func main() {
	var c config.XcbbRobotConfig
	c.Initconf()
	fmt.Println(c.Logfile)
	fmt.Println(c.Loglevel)
}