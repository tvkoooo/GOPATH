package message

import (
	"xcbbrobot/common/datastream"
	"xcbbrobot/common/datagroove"

)

//PRealJoinChannel RQ DATA  (360 << 8) | 2
type PRealLeaveChannelRQ struct {
	Uid uint32
	Sid uint32
}

/////////////------------------------//////////////////////////////////////
func (b *PRealLeaveChannelRQ )WriteMessageWriteMessage( d *datagroove.DataBuff ) () {
	var ph PackHead
	ph.Uri = (360 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	ph.Length = uint32(13+8)

	ph.WritePackHead(d)

	d.DataSlotWriteUint32(d.LenRemove + d.LenData + 13 , b.Uid)
	d.DataSlotWriteUint32(d.LenRemove + d.LenData + 17 , b.Sid)

	d.LenData += int(ph.Length)
}
//只用于测试，是否可以还原数据
func (b *PRealLeaveChannelRQ )ReadPackBody(d *datagroove.DataBuff , length int) () {
	b.Uid = d.DataSlotReadUint32(d.LenRemove+ 13)
	b.Sid = d.DataSlotReadUint32(d.LenRemove+ 17)
	d.LenRemove += length
	d.LenData -= length
}
/////////////------------------------//////////////////////////////////////













func SlotSendPRealLeaveChannelRQ(d *datagroove.DataBuff ,uid uint32 , sid uint32)() {
	var ph PackHead
	ph.Uri = (360 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	var robotLeave PRealLeaveChannelRQ
	robotLeave.Uid = uid
	robotLeave.Sid = sid

	ph.Length = uint32(13+8)

	WritePeakHead(d , &ph)
	WritePRealLeaveChannelRQ(d, &robotLeave)
	d.LenData += int(ph.Length)
}
func WritePRealLeaveChannelRQ(d *datagroove.DataBuff ,rq *PRealLeaveChannelRQ ) () {
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+13,rq.Uid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+17,rq.Sid)

}


func SendPRealLeaveChannelRQ(uid uint32 , sid uint32 )( mess []byte){
	var ph PackHead
	ph.Uri = (360 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
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
