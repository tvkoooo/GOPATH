package message

import (
	"errors"
	"fmt"
	"xcbbrobot/common/datagroove"
	"xcbbrobot/common/datastream"
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
	d.LenData += int(ph.Length)
}

//只用于测试，是否可以还原数据
//读出注册消息体
func (b *PRegisteredPI) ReadPackBody(d *datagroove.DataBuff, length int) {
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
	d.LenData += int(ph.Length)
}

//读出控制机器人消息体
func (b *PRobotServerCmd) ReadPackBody(d *datagroove.DataBuff, length int) {
	b.Cmd = d.DataSlotReadUint8(d.LenRemove + 13)
	b.Sid = d.DataSlotReadUint32(d.LenRemove + 14)
	b.Param1 = d.DataSlotReadUint32(d.LenRemove + 18)
	b.Param2 = d.DataSlotReadUint32(d.LenRemove + 22)
	d.LenRemove += length
	d.LenData -= length
}

/////////////------------------------//////////////////////////////////////
func (b *PRobotServerCmd) WriteBody(pB *PackBody) {
	pB.StreamWriteUint8(b.Cmd)
	pB.StreamWriteUint32(b.Sid)
	pB.StreamWriteUint32(b.Param1)
	pB.StreamWriteUint32(b.Param2)
}

func (b *PRobotServerCmd) ReadBody(pB *PackBody) {
	b.Cmd = pB.StreamReadUint8()
	b.Sid = pB.StreamReadUint32()
	b.Param1 = pB.StreamReadUint32()
	b.Param2 = pB.StreamReadUint32()
}

func SendPRegisteredPI() (mess []byte) {
	var ph PackHead
	ph.Uri = (101 << 8) | 23
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	var robotcon PRegisteredPI
	robotcon.Id = 0
	robotcon.PIType = 64
	robotcon.PIPass = ""

	body := EncodePRegisteredPIBody(robotcon)
	mess = AddPeakHead(ph, body)
	return mess
}

func ReceivePRobotServerCmd(mess []byte) (length uint32, ph PackHead, rs PRobotServerCmd) {
	var body []byte
	length, ph, body = PopPeakHead(mess)
	rs, _ = DecodePRobotServerCmdBody(body)
	fmt.Printf("the ph : %+v\n", ph)
	fmt.Printf("the rs : %+v\n", rs)
	return length, ph, rs
}

func EncodePRegisteredPIBody(rq PRegisteredPI) (outbyte []byte) {
	body := make([]byte, 0)
	outbyte = datastream.AddUint32(rq.Id, body)
	outbyte = datastream.AddUint32(rq.PIType, outbyte)
	outbyte = datastream.AddString16(rq.PIPass, outbyte)
	return outbyte
}

//used test
func DecodePRegisteredPIBody(inbyte []byte) (rq PRegisteredPI, err error) {
	body := make([]byte, 0)
	rq.Id, body = datastream.GetUint32(inbyte)
	rq.PIType, body = datastream.GetUint32(body)
	rq.PIPass, body = datastream.GetString(body)

	if 0 != len(body) {
		err = errors.New("解析 PRegisteredPI 失败，解析模具错误")
	} else {
		err = nil
	}
	return rq, err
}

func DecodePRobotServerCmdBody(inbyte []byte) (rs PRobotServerCmd, err error) {
	body := make([]byte, 0)
	rs.Cmd, body = datastream.GetUint8(inbyte)
	rs.Sid, body = datastream.GetUint32(body)
	rs.Param1, body = datastream.GetUint32(body)
	rs.Param2, body = datastream.GetUint32(body)
	if 0 != len(body) {
		err = errors.New("解析 PRobotServerCmd 失败，解析模具错误")
	} else {
		err = nil
	}
	return rs, err
}

//数据槽增加 机器人注册 消息
func Write_PRegisteredPI(d *datagroove.DataBuff) {
	var ph PackHead
	ph.Uri = (101 << 8) | 23
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	var robotCon PRegisteredPI
	robotCon.Id = 0
	robotCon.PIType = 64
	robotCon.PIPass = ""

	ph.Length = uint32(13 + 8 + 2 + len(robotCon.PIPass))

	var pB PackBody
	pB.Body = d.SGroove
	pB.OffSet = d.LenRemove + d.LenData
	pB.Len = 13

	d.LenData += int(13)

	pB.Body = d.SGroove
	pB.OffSet = d.LenRemove + d.LenData + 13
	pB.Len = 8 + 2 + len(robotCon.PIPass)

	d.LenData += int(8 + 2 + len(robotCon.PIPass))
}
