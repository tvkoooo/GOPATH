package main

import (
	"time"

	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		go printnew(i)
	}
	time.Sleep(60E9)
}

func printnew(a int) {
	i := 0
	for {
		fmt.Println("线程", a, "第", i, "次打印")
		i++
		time.Sleep(1E9)
		if i > 30 {
			fmt.Println("线程", a, "结束")
			break
		}
	}
}
