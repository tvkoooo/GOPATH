//This package deals with the user leaving the studio
//PRealLeaveChannel
package userleave

import (
	"encoding/json"
	"fmt"
	"lj/xcbblinktest/datastream"
	"unsafe"
)

//Peakhead pack head
type Peakhead struct {
	Uri     uint32
	Sid     uint16
	Rescode uint16
	Tag     uint8
}

//PRealLeaveChannel RQ DATA
type PRealLeaveChannelRQ struct {
	Uid uint32
	Sid uint32
}

//PRealLeaveChannelRS RQ DATA
type PRealLeaveChannelRS struct {
	Code uint32
	Desc string
	Sid  uint32
	Uid  uint32
}

//PRealLeaveChannel RQ add peak uri head
func Addpeakhead(uri uint32, inbyte []byte) (outbyte []byte) {
	var ph Peakhead
	ph.Uri = uri
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	outbyte = datastream.AddUint32(ph.Uri, inbyte)
	outbyte = datastream.AddUint16(ph.Sid, outbyte)
	outbyte = datastream.AddUint16(ph.Rescode, outbyte)
	outbyte = datastream.AddUint8(ph.Tag, outbyte)
	return outbyte
}

//PRealLeaveChannel RQ add user body struct to datastream
func ADDsenderbody(uri uint32, rq PRealLeaveChannelRQ) (outbyte []byte) {
	sendstream := make([]byte, 0)
	length := uint32(unsafe.Sizeof(rq)) + 13
	outbyte = datastream.AddUint32(length, sendstream)
	outbyte = Addpeakhead(uri, outbyte)
	outbyte = datastream.AddUint32(rq.Uid, outbyte)
	outbyte = datastream.AddUint32(rq.Sid, outbyte)
	jsonrq, _ := json.Marshal(rq)
	fmt.Println("Send body \n length:", length, "uri:", uri, "rq = ", string(jsonrq))
	return outbyte
}

//PRealLeaveChannelRS get user body struct from datastream
func Getreceivebody(inbyte []byte) (rs PRealLeaveChannelRS) {
	var length uint32
	var ph Peakhead
	rsdata := make([]byte, 0)
	length, rsdata = datastream.GetUint32(inbyte)
	ph.Uri, rsdata = datastream.GetUint32(rsdata)
	ph.Sid, rsdata = datastream.GetUint16(rsdata)
	ph.Rescode, rsdata = datastream.GetUint16(rsdata)
	ph.Tag, rsdata = datastream.GetUint8(rsdata)

	rs.Code, rsdata = datastream.GetUint32(rsdata)
	rs.Desc, rsdata = datastream.GetString(rsdata)
	rs.Uid, rsdata = datastream.GetUint32(rsdata)
	rs.Sid, rsdata = datastream.GetUint32(rsdata)
	jsonph, _ := json.Marshal(ph)
	jsonrs, _ := json.Marshal(rs)
	fmt.Println("Receive body \n length:", length, "peakhead:", string(jsonph), "rs=", string(jsonrs))
	return
}
