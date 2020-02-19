package message

import (
	"robot_d/common/datagroove"
)

//pack head
type PackHead struct {
	Length  uint32
	Uri     uint32
	Sid     uint16
	ResCode uint16
	Tag     uint8
}

//ph.Uri = uri
//ph.Sid = 0
//ph.Rescode = 200
//ph.Tag = 1

/////////////------------------------//////////////////////////////////////
//在写完消息后，添加数据包头PackHead
func (p *PackHead) WritePackHead(d *datagroove.DataBuff) {
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+0, p.Length)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+4, p.Uri)
	d.DataSlotWriteUint16(d.LenRemove+d.LenData+8, p.Sid)
	d.DataSlotWriteUint16(d.LenRemove+d.LenData+10, p.ResCode)
	d.DataSlotWriteUint8(d.LenRemove+d.LenData+12, p.Tag)
}

//对数据槽进行 完整解数据包包头PackHead
func (p *PackHead) ReadPackHead(d *datagroove.DataBuff) {
	p.Length = d.DataSlotReadUint32(d.LenRemove + 0)
	p.Uri = d.DataSlotReadUint32(d.LenRemove + 4)
	p.Sid = d.DataSlotReadUint16(d.LenRemove + 8)
	p.ResCode = d.DataSlotReadUint16(d.LenRemove + 10)
	p.Tag = d.DataSlotReadUint8(d.LenRemove + 12)
}

//备注：读出包头数据槽读取位置偏移量Offset 是不会偏移的，只有读出完整包消息或者丢弃完整包消息才会进行偏移
//如果数据槽数据不够一个完整包，返回一个 -1 ，需要等待下次查看数据槽数据是否完整
//如果数据槽数据够一个完整包，并且是所需要的uri ，返回0值，可以正常解包
//如果数据槽数据够一个完整包，但不是所需要的uri ，返回数据包长度length ，方便解包操作时数据槽是否丢弃该消息（令数据槽 LenRemove += length 和 lenData -= length）
func CheckUri(d *datagroove.DataBuff, uri uint32) int {
	Length := d.DataSlotReadUint32(d.LenRemove + 0)
	if int(Length) > d.LenData {
		return -1
	}

	checkUri := d.DataSlotReadUint32(d.LenRemove + 4)
	if checkUri == uri {
		return 0
	} else {
		return int(Length)
	}
}

//备注：读出包头数据槽读取位置偏移量Offset 是不会偏移的，只有读出完整包消息或者丢弃完整包消息才会进行偏移
//如果检出 uri 后有必要输出包头，需要把剩余包头数据读出
func (p *PackHead) ReadPackHeadRemain(d *datagroove.DataBuff) {
	p.Sid = d.DataSlotReadUint16(d.LenRemove + 8)
	p.ResCode = d.DataSlotReadUint16(d.LenRemove + 10)
	p.Tag = d.DataSlotReadUint8(d.LenRemove + 12)
}

/////////////------------------------//////////////////////////////////////
