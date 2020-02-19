//This package deals with the user leaving the studio
//PRealJoinChannel
package usercome

import (
	"encoding/json"
	"fmt"
	"lj/xcbblinktest/datastream"
)

//Peakhead pack head
type Peakhead struct {
	Uri     uint32
	Sid     uint16
	Rescode uint16
	Tag     uint8
}

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

//PRealJoinChannel RQ add peak uri head
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

//PRealJoinChannel RQ add user body struct to datastream
func ADDsenderbody(uri uint32, rq PRealJoinChannelRQ) (outbyte []byte) {
	sendstream := make([]byte, 0)
	length := uint32(13 + 4 + 2 + 4 + 4 + 2 + 4)
	outbyte = datastream.AddUint32(length, sendstream)
	outbyte = Addpeakhead(uri, outbyte)
	outbyte = datastream.AddUint32(rq.Uid, outbyte)
	outbyte = datastream.AddString16(rq.Sha1Pass, outbyte)
	outbyte = datastream.AddUint32(rq.Sid, outbyte)
	outbyte = datastream.AddUint32(rq.Ssid, outbyte)
	outbyte = datastream.AddString16(rq.SsPass, outbyte)
	outbyte = datastream.AddUint32(rq.Version, outbyte)

	jsonrq, _ := json.Marshal(rq)
	fmt.Println("Send body \n length:", length, "uri:", uri, "rq = ", string(jsonrq))
	fmt.Println(outbyte)
	return outbyte
}

//PRealJoinChannel get user body struct from datastream
func Getreceivebody(inbyte []byte) {
	var length uint32
	var rechead Peakhead
	var outbyte []byte
	length, outbyte = datastream.GetUint32(inbyte)
	rechead.Uri, outbyte = datastream.GetUint32(outbyte)
	fmt.Println("Receive body  length :", length, "Receive body  Uri :", rechead.Uri)
	//fmt.Println("Receive body \n :",inbyte )
}
