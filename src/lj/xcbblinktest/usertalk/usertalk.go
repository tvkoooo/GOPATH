//This package deals with the user talk the studio
//PFeatureRequest (1 << 8) | 23
package usertalk

import (
	"fmt"
	"lj/xcbblinktest/datastream"
	"encoding/json"
)

//Peakhead pack head
type Peakhead struct {
	Uri     uint32
	Sid     uint16
	Rescode uint16
	Tag     uint8
}

type PEnterChannel struct {
	PIType uint32
	Cmd string
}

//PFeatureRequest RQ DATA
type PFeatureRequestRQ struct {
	Cmd string `json:"cmd"`
	Uid uint32 `json:"uid"`
	Singerid uint32 `json:"singerid"`
	Sid uint32 `json:"sid"`
	Context string `json:"context"`
}

//PFeatureRequest RS DATA
type PPlusRS struct {
	Code uint32
	Desc string
	Sid  uint32
	Uid  uint32
}

//PPlus RQ add peak uri head
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
func ADDsenderbody(uri uint32 ,rq PFeatureRequestRQ) (outbyte []byte) {
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
//PPlusRS get user body struct from datastream
func Getreceivebody(inbyte []byte) () {
	fmt.Println("Receive body \n :",inbyte )
}
