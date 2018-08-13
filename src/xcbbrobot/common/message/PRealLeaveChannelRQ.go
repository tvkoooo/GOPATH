package message

import (
	"xcbbrobot/common/datastream"
)

//PRealJoinChannel RQ DATA  (360 << 8) | 2
type PRealLeaveChannelRQ struct {
	Uid uint32
	Sid uint32
}



func SendPRealLeaveChannelRQ(uid uint32 , sid uint32 )( mess []byte){
	var ph Packhead
	ph.Uri = (360 << 8) | 2
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	var robotLeave PRealLeaveChannelRQ
	robotLeave.Uid = uid
	robotLeave.Sid = sid

	body:= EncodePRealLeaveChannelRQ(robotLeave)
	mess = AddPeakHead(ph ,body)
	return mess
}


func EncodePRealLeaveChannelRQ(rq PRealLeaveChannelRQ ) (outbyte []byte) {
	body := make([]byte,0)
	outbyte = datastream.AddUint32(rq.Uid ,body )
	outbyte = datastream.AddUint32(rq.Sid, outbyte)
	return outbyte
}
