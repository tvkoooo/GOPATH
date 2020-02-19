package main

import (
	"fmt"
	"time"
	"xcbbrobot/common/message"
)

func main() {
	var ph message.Packhead
	ph.Uri = (12 << 8) | 4
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	var robotping message.PPlusRQ
	robotping.Uid = 10000099
	robotping.Sid = 102692
	robotping.Stampc = uint32(time.Now().Unix())
	robotping.Stamps = uint32(time.Now().UnixNano() % 1E9)

	body := message.EncodePPlusBody(robotping)
	mess := message.AddPeakHead(ph, body)
	fmt.Println("meassage datastream: ", mess)

	lengthde, phde, bodyde := message.PopPeakHead(mess)
	fmt.Println("length: ", lengthde)
	fmt.Printf("the ph : %+v\n", phde)
	fmt.Println("body datastream: ", bodyde)
	de_rq, err := message.DecodePPlusBody(bodyde)
	if nil == err {
		fmt.Printf("the rq :%+v\n", de_rq)
	}

}
