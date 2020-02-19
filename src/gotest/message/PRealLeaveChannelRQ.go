package message

import (
	"gotest/common/datagroove"
)

//PRealJoinChannel RQ DATA  (360 << 8) | 2
type PRealLeaveChannelRQ struct {
	Uid uint32
	Sid uint32
}

/////////////------------------------//////////////////////////////////////
//拼装 PRealJoinChannel 消息
func WritePRealLeaveChannelBuff(d *datagroove.DataBuff, robotId uint32, sid uint32) {
	var p PRealLeaveChannelRQ
	p.Uid = robotId
	p.Sid = sid
	p.WriteMessage(d)
}

/////////////------------------------//////////////////////////////////////
func (b *PRealLeaveChannelRQ) WriteMessage(d *datagroove.DataBuff) {
	var ph PackHead
	ph.Uri = (360 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	ph.Length = uint32(13 + 8)

	ph.WritePackHead(d)

	d.DataSlotWriteUint32(d.LenRemove+d.LenData+13, b.Uid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+17, b.Sid)

	d.LenData += uint16(ph.Length)
}

//只用于测试，是否可以还原数据
func (b *PRealLeaveChannelRQ) ReadPackBody(d *datagroove.DataBuff, length uint16) {
	b.Uid = d.DataSlotReadUint32(d.LenRemove + 13)
	b.Sid = d.DataSlotReadUint32(d.LenRemove + 17)
	d.LenRemove += length
	d.LenData -= length
}

/////////////------------------------//////////////////////////////////////
