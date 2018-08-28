package main

import (
	"time"
	"fmt"
)

type ntest struct {
	a int
	b int
	c string
}

func main()  {
	var oneS time.Duration
	oneS = 1E9*24*60*60
	t1 := time.Now()
	time.Sleep(2E9)
	sub := time.Now().Sub(t1)
	nps := oneS / sub
	fmt.Println("sub:" ,sub,"nps:" ,nps )

	var aaa ntest
	aaa.a = 3
	aaa.b = 8
	aaa.c = "sjdfoj"
	fmt.Printf("%+v", aaa)
}