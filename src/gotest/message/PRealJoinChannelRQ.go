package message

import (
	"gotest/common/datagroove"
)

//PRealJoinChannel RQ DATA  (32 << 8) | 2
type PRealJoinChannelRQ struct {
	Uid      uint32
	Sha1Pass string
	Sid      uint32
	Ssid     uint32
	SsPass   string
	Version  uint32
}

//PRealJoinChannelRS RQ DATA
type PRealJoinChannelRS struct {
	Code uint32
	Desc string
	Sid  uint32
	Uid  uint32
}

/////////////------------------------//////////////////////////////////////
//拼装 PRealJoinChannel 消息
func WritePRealJoinChannelBuff(d *datagroove.DataBuff, robotId uint32, sid uint32) {
	var p PRealJoinChannelRQ
	p.Uid = robotId
	p.Sha1Pass = ""
	p.Sid = sid
	p.Ssid = 1
	p.SsPass = ""
	p.Version = 1
	p.WriteMessage(d)

}

/////////////------------------------//////////////////////////////////////
func (b *PRealJoinChannelRQ) WriteMessage(d *datagroove.DataBuff) {
	var ph PackHead
	ph.Uri = (32 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	ph.Length = uint32(13 + 16 + 2 + 2 + len(b.Sha1Pass) + len(b.SsPass))

	ph.WritePackHead(d)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+13, b.Uid)
	d.DataSlotWriteString16(d.LenRemove+d.LenData+17, &b.Sha1Pass)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+19+uint16(len(b.Sha1Pass)), b.Sid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+23+uint16(len(b.Sha1Pass)), b.Ssid)
	d.DataSlotWriteString16(d.LenRemove+d.LenData+27+uint16(len(b.Sha1Pass)), &b.SsPass)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+29+uint16(len(b.Sha1Pass))+uint16(len(b.SsPass)), b.Version)
	d.LenData += uint16(ph.Length)
}

//只用于测试，是否可以还原数据
func (b *PRealJoinChannelRQ) ReadPackBody(d *datagroove.DataBuff, length uint16) {
	b.Uid = d.DataSlotReadUint32(d.LenRemove + 13)
	d.DataSlotReadString16(d.LenRemove+17, &b.Sha1Pass)
	b.Sid = d.DataSlotReadUint32(d.LenRemove + 19 + uint16(len(b.Sha1Pass)))
	b.Ssid = d.DataSlotReadUint32(d.LenRemove + 23 + uint16(len(b.Sha1Pass)))
	d.DataSlotReadString16(d.LenRemove+27+uint16(len(b.Sha1Pass)), &b.SsPass)
	b.Version = d.DataSlotReadUint32(d.LenRemove + 29 + uint16(len(b.Sha1Pass)) + uint16(len(b.SsPass)))
	d.LenRemove += length
	d.LenData -= length
}

/////////////------------------------//////////////////////////////////////
