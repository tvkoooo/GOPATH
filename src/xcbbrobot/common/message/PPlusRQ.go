package message

import (
	"xcbbrobot/common/datastream"
	"errors"
	"time"
	"xcbbrobot/common/datagroove"
)

//PPlus RQ DATA  (12 << 8) | 4
type PPlusRQ struct {
	Uid uint32
	Sid uint32
	Stampc uint32
	Stamps uint32
}

func SlotSendPPlus(d *datagroove.DataBuff ,uid uint32 , sid uint32)() {
	var ph Packhead
	ph.Uri = (12 << 8) | 4
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	var robotping PPlusRQ
	robotping.Uid = uid
	robotping.Sid = sid
	robotping.Stampc = uint32(time.Now().Unix())
	robotping.Stamps = uint32(time.Now().UnixNano() % 1E9)
	ph.Length = uint32(13+16)

	WritePeakHead(d , &ph)
	WritePPlus(d, &robotping)
	d.LenData += int(ph.Length)
}
func WritePPlus(d *datagroove.DataBuff ,rq *PPlusRQ ) () {
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+13,rq.Uid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+17,rq.Sid)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+21,rq.Stampc)
	d.DataSlotWriteUint32(d.LenRemove+d.LenData+25,rq.Stamps)
}




func SendPPlusRQ(uid uint32 , sid uint32 )( mess []byte){
	var ph Packhead
	ph.Uri = (12 << 8) | 4
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	var robotping PPlusRQ
	robotping.Uid = uid
	robotping.Sid = sid
	robotping.Stampc = uint32(time.Now().Unix())
	robotping.Stamps = uint32(time.Now().UnixNano() % 1E9)


	body:= EncodePPlusBody(&robotping)
	mess = AddPeakHead(ph ,body)
	return mess
}



func EncodePPlusBody(rq *PPlusRQ ) (outbyte []byte) {
	body := make([]byte,0)
	outbyte = datastream.AddUint32(rq.Uid ,body )
	outbyte = datastream.AddUint32(rq.Sid, outbyte)
	outbyte = datastream.AddUint32(rq.Stampc, outbyte)
	outbyte = datastream.AddUint32(rq.Stamps, outbyte)
	return outbyte
}
//used test
func DecodePPlusBody(inbyte []byte) (rq *PPlusRQ ,err error) {
	body := make([]byte,0)
	rq.Uid,body =  datastream.GetUint32(inbyte)
	rq.Sid,body =  datastream.GetUint32(body)
	rq.Stampc,body =  datastream.GetUint32(body)
	rq.Stamps,body =  datastream.GetUint32(body)
	if 0 !=len(body) {
		err = errors.New("解析 PRealJoinChannelRQ 失败，解析模具错误")
	}else {
		err = nil
	}
	return rq ,err
}
