package main

import (
	"time"
	"fmt"
	"strings"
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
	fmt.Printf("%+v \n", aaa)

	var nowBefore time.Time
	nowBefore = time.Now().Add(7E8)

	iii := 0
	for  {

		time.Sleep(2E8)
		intev :=time.Now().Sub(nowBefore)
		if intev>0 {
			fmt.Println("intev >0",intev,iii)
		}
		fmt.Println(intev,iii)
		iii ++
		if iii==1 {
			break
		}
	}

	var ssssss string
	ssssss = "59.110.125.134-30302"
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println(strings.Replace(ssssss, "-", ":", -1))
}