package message

import (
	"xcbbrobot/common/datastream"
	"errors"
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

func SendPRealJoinChannel(uid uint32 , sid uint32 )( mess []byte){
	var ph Packhead
	ph.Uri = (32 << 8) | 2
	ph.Sid = 0
	ph.Rescode = 200
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