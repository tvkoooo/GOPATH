package message

import (
	"xcbbrobot/common/datastream"
	"xcbbrobot/common/typechange"
)

//Peakhead pack head
type Packhead struct {
	Uri     uint32
	Sid     uint16
	Rescode uint16
	Tag     uint8
}

//ph.Uri = uri
//ph.Sid = 0
//ph.Rescode = 200
//ph.Tag = 1
func AddPeakHead(ph Packhead ,inbyte []byte) (outbyte []byte) {
	length := uint32(len(inbyte))
	mess := make([]byte ,0,length + 13)
	outbyte = datastream.AddUint32(length + 13, mess)
	outbyte = datastream.AddUint32(ph.Uri, outbyte)
	outbyte = datastream.AddUint16(ph.Sid, outbyte)
	outbyte = datastream.AddUint16(ph.Rescode, outbyte)
	outbyte = datastream.AddUint8(ph.Tag, outbyte)
	outbyte = datastream.AddBytes(outbyte,inbyte)
	return outbyte
}

func PopPeakHead(inbyte []byte) (length uint32 ,ph Packhead ,body []byte) {
	length,body = datastream.GetUint32(inbyte)
	ph.Uri,body =  datastream.GetUint32(body)
	ph.Sid,body =  datastream.GetUint16(body)
	ph.Rescode,body =  datastream.GetUint16(body)
	ph.Tag,body =  datastream.GetUint8(body)
	return length,ph,body
}

func CheckPeakHead(inbyte []byte, uri uint32) (bool) {
	getUri := typechange.Slice_2_uint32(inbyte[4:8])
	if getUri == uri{
		return true
	}else {
		return false
	}
}
