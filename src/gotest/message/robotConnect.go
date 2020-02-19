package message

import (
	"gotest/common/datagroove"
)

//PRegisteredPI RQ DATA  (101 << 8) | 23
type PRegisteredPI struct {
	Id     uint32
	PIType uint32
	PIPass string
}

//RSC_LIVE_START = 0,
//RSC_LIVE_STOP,
//RSC_LIVE_END,
//RSC_ADD_ROBOT, //param1为要添加的机器人数量
//RSC_REMOVE_ROBOT //param1为要减少的机器人数量
//PRobotServerCmd  DATA  (131 << 8) | 2
type PRobotServerCmd struct {
	Cmd    uint8
	Sid    uint32
	Param1 uint32
	Param2 uint32
}

/////////////------------------------//////////////////////////////////////
//拼装注册消息
func WritePRegisteredPI(d *datagroove.DataBuff) {
	var robotCon PRegisteredPI
	robotCon.Id = 0
	robotCon.PIType = 64
	robotCon.PIPass = ""

	robotCon.WriteMessage(d)
}

/////////////------------------------//////////////////////////////////////
//注册消息写入数据槽
func (b *PRegisteredPI) WriteMessage(d *datagroove.DataBuff) {
	var ph PackHead
	ph.Uri = (101 << 8) | 23
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	ph.Length = uint32(13 + 8 + 2 + len(b.PIPass))

	ph.WritePackHead(d)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+13, b.Id)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+17, b.PIType)
	d.DataSlotWriteString16(d.LenRemove+d.LenData+21, &b.PIPass)
	d.LenData += uint16(ph.Length)
}

//只用于测试，是否可以还原数据
//读出注册消息体
func (b *PRegisteredPI) ReadPackBody(d *datagroove.DataBuff, length uint16) {
	b.Id = d.DataSlotReadUint32(d.LenRemove + 13)
	b.PIType = d.DataSlotReadUint32(d.LenRemove + 17)
	d.DataSlotReadString16(d.LenRemove+21, &b.PIPass)
	d.LenRemove += length
	d.LenData -= length
}

/////////////------------------------//////////////////////////////////////
//只用于测试，仿造服务器数据
//控制机器人消息写入数据槽
func (b *PRobotServerCmd) WriteMessage(d *datagroove.DataBuff) {
	var ph PackHead
	ph.Uri = (131 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	ph.Length = uint32(13 + 13)

	ph.WritePackHead(d)
	d.DataSlotWriteUint8(d.LenRemove+d.LenData+13, b.Cmd)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+14, b.Sid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+18, b.Param1)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+22, b.Param2)
	d.LenData += uint16(ph.Length)
}

//读出控制机器人消息体
func (b *PRobotServerCmd) ReadPackBody(d *datagroove.DataBuff, length uint16) {
	b.Cmd = d.DataSlotReadUint8(d.LenRemove + 13)
	b.Sid = d.DataSlotReadUint32(d.LenRemove + 14)
	b.Param1 = d.DataSlotReadUint32(d.LenRemove + 18)
	b.Param2 = d.DataSlotReadUint32(d.LenRemove + 22)
	d.LenRemove += length
	d.LenData -= length
}

/////////////------------------------//////////////////////////////////////
