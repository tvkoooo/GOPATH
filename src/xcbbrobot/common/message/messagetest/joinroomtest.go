package main

import (
	"fmt"
	"xcbbrobot/common/message"
)

func main() {
	var ph message.Packhead
	ph.Uri = (32 << 8) | 2
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	var robotcome message.PRealJoinChannelRQ
	robotcome.Uid = 10000099
	robotcome.Sha1Pass = ""
	robotcome.Sid = 102692
	robotcome.Ssid = 1
	robotcome.SsPass = ""
	robotcome.Version = 1

	body := message.EncodeJoinChanneBody(robotcome)
	mess := message.AddPeakHead(ph, body)
	fmt.Println("meassage datastream: ", mess)

	lengthde, phde, bodyde := message.PopPeakHead(mess)
	fmt.Println("length: ", lengthde)
	fmt.Printf("the ph : %+v\n", phde)
	fmt.Println("body datastream: ", bodyde)
	de_rq, err := message.DecodeJoinChanneBody(bodyde)
	if nil == err {
		fmt.Printf("the rq :%+v\n", de_rq)
	}

}
