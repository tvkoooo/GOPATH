package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"time"
)

func main()  {
	f, err := os.Open("D:/GOPATH/src/lj/readline/text.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line,time.Now())
		if line=="over"{
			break
		}
	}
}