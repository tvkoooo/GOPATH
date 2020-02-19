//This package deals with the user ping the studio
//PPlus (12 << 8) | 4
package userping

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

//PPlus RQ DATA
type PPlusRQ struct {
	Uid    uint32
	Sid    uint32
	Stampc uint32
	Stamps uint32
}

//PPlus RQ DATA
type PPlusRS struct {
	Code uint32
	Desc string
	Sid  uint32
	Uid  uint32
}

//PPlus RQ add peak uri head
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

//PPlus RQ add user body struct to datastream
func ADDsenderbody(uri uint32, rq PPlusRQ) (outbyte []byte) {
	sendstream := make([]byte, 0)
	length := uint32(unsafe.Sizeof(rq)) + 13
	outbyte = datastream.AddUint32(length, sendstream)
	outbyte = Addpeakhead(uri, outbyte)
	outbyte = datastream.AddUint32(rq.Uid, outbyte)
	outbyte = datastream.AddUint32(rq.Sid, outbyte)
	outbyte = datastream.AddUint32(rq.Stampc, outbyte)
	outbyte = datastream.AddUint32(rq.Stamps, outbyte)
	jsonrq, _ := json.Marshal(rq)
	fmt.Println("Send body \n length:", length, "uri:", uri, "rq = ", string(jsonrq))
	fmt.Println(outbyte)
	return outbyte
}

//PPlusRS get user body struct from datastream
func Getreceivebody(inbyte []byte) {
	fmt.Println("Receive body \n :", inbyte)
}
