//This package deals with the user comeing the studio
//PEnterChannel   use the PForwardToPlug,uri = (253 << 8 | 2)
package usercomephp

import (
	"fmt"
	"lj/xcbblinktest/datastream"
	//"unsafe"
	"encoding/json"
)

//Peakhead pack head
type Peakhead struct {
	Uri     uint32
	Sid     uint16
	Rescode uint16
	Tag     uint8
}

//PEnterChannel DATA
type PEnterChannel struct {
	PIType uint32
	Cmd string
}

//PEnterChannel RQ DATA
type PEnterChannelRQ struct {
	Cmd string `json:"cmd"`
	Sid uint32 `json:"sid"`
	Uid uint32 `json:"uid"`
	Sender string `json:"sender"`
	Uid_onmic uint32 `json:"uid_onmic"`
}

//PEnterChannel RS DATA
type REnterChannel struct {
	Cmd string
	Uid uint32
	PalyTotalTime uint32
	Money uint32
	TotalSun uint32
	CameraStatus uint32
	MicroStatus uint32
	WeekStar uint32
	Weektool_id uint32
	Week_ranking uint32
	Weektool_img string
	Weektool_name string
}

//PEnterChannel RQ add peak uri head
func Addpeakhead(uri uint32 ,inbyte []byte) (outbyte []byte) {
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

//PEnterChannel RQ add user body struct to datastream
func ADDsenderbody(uri uint32 ,rq PEnterChannelRQ) (outbyte []byte) {
	sendstream := make([]byte, 0)
	jsonrq,_ := json.Marshal(rq)
	//length := uint32(unsafe.Sizeof(rq))+13+4
	length := uint32(len(jsonrq)+13+4+4)  //json长度 + （包头13） + （PIType 4） + （cmd长度 4）
	outbyte = datastream.AddUint32(length, sendstream)
	outbyte = Addpeakhead(uri,outbyte)

	var datasend PEnterChannel
	datasend.PIType = 6
	datasend.Cmd = string(jsonrq)
	outbyte = datastream.AddUint32(datasend.PIType, outbyte)
	outbyte = datastream.AddString32(datasend.Cmd,outbyte)

	fmt.Println("Send body \n length:", length,"uri:",uri,"rq = ", string(jsonrq))
	fmt.Println(outbyte)
	return outbyte
}
//PEnterChannel get user body struct from datastream
func Getreceivebody(inbyte []byte) () {

	fmt.Println("Receive body \n length:",inbyte )
}
