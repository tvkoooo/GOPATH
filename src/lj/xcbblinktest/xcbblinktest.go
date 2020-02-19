package main

import (
	"fmt"
	"lj/xcbblinktest/user"
	"os"
	"strconv"
	"time"
)

// main
func main() {
	var num, ms int
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) < 2 {
		Usage() //如果用户没有输入,或参数个数不够,则调用该函数提示用户
		num, ms = 1, 50
	} else {
		num, _ = strconv.Atoi(args[1]) //获取输入的第一个参数,并转换为int
		ms, _ = strconv.Atoi(args[2])  //获取输入的第二个参数,并转换为int
		fmt.Println("循环次数:", num, "\n 线程间隔毫秒:", ms)
	}
	ch := make(chan int, num)

	uid := uint32(10000094) //10000094 10005260 10000108
	sid := uint32(102692)
	for i := 0; i < num; i++ {

		go user.Userctrl(&ch, uid, sid)
		time.Sleep(time.Duration(ms * 1E6))
		uid++

	}
	<-ch
	time.Sleep(3 * 1E9)

	fmt.Println("conn 关闭")
	close(ch)
}

var Usage = func() {
	fmt.Println("使用了默认参数 1 次 50ms")
}
