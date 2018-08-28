package message

import (
	"xcbbrobot/common/datastream"
	"errors"
	"xcbbrobot/common/datagroove"
)

//PRealJoinChannel RQ DATA  (32 << 8) | 2
type PRealJoinChannelRQ struct {
	Uid uint32
	Sha1Pass string
	Sid uint32
	Ssid uint32
	SsPass string
	Version uint32
}

//PRealJoinChannelRS RQ DATA
type PRealJoinChannelRS struct {
	Code uint32
	Desc string
	Sid  uint32
	Uid  uint32
}

/////////////------------------------//////////////////////////////////////
func (b *PRealJoinChannelRQ )WriteMessageWriteMessage( d *datagroove.DataBuff ) () {
	var ph PackHead
	ph.Uri = (32 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	ph.Length = uint32(13+16+2+2+len(b.Sha1Pass)+len(b.SsPass))

	ph.WritePackHead(d)
	d.DataSlotWriteUint32(d.LenRemove + d.LenData + 13 , b.Uid)
	d.DataSlotWriteString16(d.LenRemove + d.LenData + 17 , &b.Sha1Pass)
	d.DataSlotWriteUint32(d.LenRemove + d.LenData + 19+len(b.Sha1Pass) , b.Sid)
	d.DataSlotWriteUint32(d.LenRemove + d.LenData + 23+len(b.Sha1Pass), b.Ssid)
	d.DataSlotWriteString16(d.LenRemove + d.LenData + 27+len(b.Sha1Pass) , &b.SsPass)
	d.DataSlotWriteUint32(d.LenRemove + d.LenData + 29+len(b.Sha1Pass)+len(b.SsPass) , b.Version)
	d.LenData += int(ph.Length)
}
//只用于测试，是否可以还原数据
func (b *PRealJoinChannelRQ )ReadPackBody(d *datagroove.DataBuff , length int) () {
	b.Uid = d.DataSlotReadUint32(d.LenRemove+ 13)
	d.DataSlotReadString16(d.LenRemove+ 17 , &b.Sha1Pass)
	b.Sid = d.DataSlotReadUint32(d.LenRemove+ 19+len(b.Sha1Pass))
	b.Ssid = d.DataSlotReadUint32(d.LenRemove+ 23+len(b.Sha1Pass))
	d.DataSlotReadString16(d.LenRemove+27+len(b.Sha1Pass) , &b.SsPass)
	b.Version = d.DataSlotReadUint32(d.LenRemove+ 29+len(b.Sha1Pass)+len(b.SsPass))
	d.LenRemove += length
	d.LenData -= length
}
/////////////------------------------//////////////////////////////////////














func SlotSendPRealJoinChannel(d *datagroove.DataBuff ,uid uint32 , sid uint32)() {
	var ph PackHead
	ph.Uri = (32 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	var robotcome PRealJoinChannelRQ
	robotcome.Uid = uid
	robotcome.Sha1Pass = ""
	robotcome.Sid = sid
	robotcome.Ssid = 1
	robotcome.SsPass = ""
	robotcome.Version = 1

	ph.Length = uint32(13+16+2+2+len(robotcome.Sha1Pass)+len(robotcome.SsPass))

	WritePeakHead(d , &ph)
	WritePRealJoinChannel(d, &robotcome)
	d.LenData += int(ph.Length)
}
func WritePRealJoinChannel(d *datagroove.DataBuff ,rq *PRealJoinChannelRQ ) () {
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+13,rq.Uid)
	d.DataSlotWriteString16(d.LenRemove+d.LenData+17,&rq.Sha1Pass)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+19,rq.Sid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+23,rq.Ssid)
	d.DataSlotWriteString16(d.LenRemove+d.LenData+27,&rq.SsPass)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+29,rq.Version)
}


func SendPRealJoinChannel(uid uint32 , sid uint32 )( mess []byte){
	var ph PackHead
	ph.Uri = (32 << 8) | 2
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	var robotcome PRealJoinChannelRQ
	robotcome.Uid = uid
	robotcome.Sha1Pass = ""
	robotcome.Sid = sid
	robotcome.Ssid = 1
	robotcome.SsPass = ""
	robotcome.Version = 1


	body:= EncodeJoinChanneBody(robotcome)
	mess = AddPeakHead(ph ,body)
	return mess
}


func EncodeJoinChanneBody(rq PRealJoinChannelRQ ) (outbyte []byte) {
	body := make([]byte,0)
	outbyte = datastream.AddUint32(rq.Uid ,body )
	outbyte = datastream.AddString16(rq.Sha1Pass, outbyte)
	outbyte = datastream.AddUint32(rq.Sid, outbyte)
	outbyte = datastream.AddUint32(rq.Ssid, outbyte)
	outbyte = datastream.AddString16(rq.SsPass, outbyte)
	outbyte = datastream.AddUint32(rq.Version, outbyte)
	return outbyte
}
//used test
func DecodeJoinChanneBody(inbyte []byte) (rq PRealJoinChannelRQ ,err error) {
	body := make([]byte,0)
	rq.Uid,body =  datastream.GetUint32(inbyte)
	rq.Sha1Pass,body =  datastream.GetString(body)
	rq.Sid,body =  datastream.GetUint32(body)
	rq.Ssid,body =  datastream.GetUint32(body)
	rq.SsPass,body =  datastream.GetString(body)
	rq.Version,body =  datastream.GetUint32(body)
	if 0 !=len(body) {
		err = errors.New("解析 PRealJoinChannelRQ 失败，解析模具错误")
	}else {
		err = nil
	}
	return rq ,err
}

func DecodeJoinChanneRs(inbyte []byte) (rs PRealJoinChannelRS ,err error) {
	body := make([]byte,0)
	rs.Code,body =  datastream.GetUint32(inbyte)
	rs.Desc,body =  datastream.GetString(body)
	rs.Sid,body =  datastream.GetUint32(body)
	rs.Uid,body =  datastream.GetUint32(body)
	if 0 !=len(body) {
		err = errors.New("解析 PRealJoinChannelRS 失败，解析模具错误")
	}else {
		err = nil
	}
	return rs ,err
}