package main

import (
	"fmt"
	"xcbbrobot/common/message"
)

func main() {
	var ph message.Packhead
	ph.Uri = (12 << 8) | 4
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	var robotcon message.PRegisteredPI
	robotcon.Id = 0
	robotcon.PIType = 64
	robotcon.PIPass = ""

	body := message.EncodePRegisteredPIBody(robotcon)
	mess := message.AddPeakHead(ph, body)
	fmt.Println("meassage datastream: ", mess)

	lengthde, phde, bodyde := message.PopPeakHead(mess)
	fmt.Println("length: ", lengthde)
	fmt.Printf("the ph : %+v\n", phde)
	fmt.Println("body datastream: ", bodyde)
	de_rq, err := message.DecodePRegisteredPIBody(bodyde)
	if nil == err {
		fmt.Printf("the rq :%+v\n", de_rq)
	}

}
