package main

import (
	"os"
	"os/signal"
	"time"
	"fmt"
)

const (
	STATECLOSE    = 0
	STATEMOTION   = 1
	STATEFINISH   = 2
)

type Chantests struct {
	ThreadContext chan int
	StateThread int
}

func (p *Chantests)dinit()  {
	p.ThreadContext = make(chan int)
	p.StateThread = STATECLOSE
}

func (p *Chantests)ddestroy()  {
	p.ThreadContext = nil
	p.StateThread = STATECLOSE
}


func (p *Chantests)start()  {
	if STATEFINISH == p.StateThread {
		p.StateThread = STATECLOSE
	}else {
		p.StateThread = STATEMOTION
	}
	go p.loop()
}

func (p *Chantests)shutDown()  {
	p.StateThread = STATEFINISH
}


func (p *Chantests)join()  {
	p.StateThread = <-p.ThreadContext
}


func (p *Chantests)loop()  {
	fmt.Println("Loop now")
	var iii int = 0
	for ; p.StateThread == STATEMOTION; {
		iii ++
		fmt.Println("print now:",iii)
		time.Sleep(1E9)
	}
	p.ThreadContext<-p.StateThread
}

func tttt(channn chan os.Signal,ppppp *Chantests)  {
	//<-channn
	time.Sleep(9E9)
	ppppp.shutDown()
}

func main()  {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	var chtest Chantests

	go tttt(c,&chtest)

	chtest.dinit()
	chtest.start()
	chtest.join()
	chtest.ddestroy()

}