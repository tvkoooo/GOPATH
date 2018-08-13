package main

import (
	"time"
	"xcbbrobot/robotlive"
)

func main()  {
	var(
		r1 robotlive.RobotSeed
		r2 robotlive.RobotSeed
		r3 robotlive.RobotSeed
	)


	go r1.RobotWork()
	go r2.RobotWork()
	go r3.RobotWork()

	go func(p *robotlive.RobotSeed) {
		time.Sleep(9E9)
		p.RobotRest()
	}(&r1)

	go func(p *robotlive.RobotSeed) {
		time.Sleep(4E9)
		p.RobotRest()
	}(&r2)

	go func(p *robotlive.RobotSeed) {
		time.Sleep(7E9)
		p.RobotRest()
	}(&r3)

	time.Sleep(10E9)
}

