package message

import (
	"robot_d/common/datagroove"
	"time"
)

/////////////------------------------//////////////////////////////////////
//PPlus RQ DATA  (12 << 8) | 4
type PPlusRQ struct {
	Uid    uint32
	Sid    uint32
	Stampc uint32
	Stamps uint32
}

/////////////------------------------//////////////////////////////////////
//拼装 PPlus 消息
func WritePPlusBuff(d *datagroove.DataBuff, robotId uint32, sid uint32) {
	var p PPlusRQ
	p.Uid = robotId
	p.Sid = sid
	p.Stampc = uint32(time.Now().Unix())
	p.Stamps = uint32(time.Now().UnixNano() % 1E9)
	p.WriteMessage(d)
}

/////////////------------------------//////////////////////////////////////
func (b *PPlusRQ) WriteMessage(d *datagroove.DataBuff) {
	var ph PackHead
	ph.Uri = (12 << 8) | 4
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	ph.Length = uint32(13 + 16)

	ph.WritePackHead(d)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+13, b.Uid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+17, b.Sid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+21, b.Stampc)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+25, b.Stamps)
	d.LenData += int(ph.Length)
}

//只用于测试，是否可以还原数据
func (b *PPlusRQ) ReadPackBody(d *datagroove.DataBuff, length int) {
	b.Uid = d.DataSlotReadUint32(d.LenRemove + 13)
	b.Sid = d.DataSlotReadUint32(d.LenRemove + 17)
	b.Stampc = d.DataSlotReadUint32(d.LenRemove + 21)
	b.Stamps = d.DataSlotReadUint32(d.LenRemove + 25)
	d.LenRemove += length
	d.LenData -= length
}

/////////////------------------------//////////////////////////////////////
