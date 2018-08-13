package message

import (
	"xcbbrobot/common/datastream"
	"errors"
	"fmt"
)

//PRegisteredPI RQ DATA  (101 << 8) | 23
type PRegisteredPI struct {
	Id     uint32
	PIType uint32
	PIPass string
}

//RSC_LIVE_START = 0,
//RSC_LIVE_STOP,
//RSC_LIVE_END,
//RSC_ADD_ROBOT, //param1为要添加的机器人数量
//RSC_REMOVE_ROBOT //param1为要减少的机器人数量
//PRobotServerCmd  DATA  (131 << 8) | 2
type PRobotServerCmd struct {
	Cmd     uint8
	Sid     uint32
	param1  uint32
	param2  uint32
}


func SendPRegisteredPI()(mess []byte){
	var ph Packhead
	ph.Uri = (101 << 8) | 23
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	var robotcon PRegisteredPI
	robotcon.Id = 0
	robotcon.PIType = 64
	robotcon.PIPass = ""

	body:= EncodePRegisteredPIBody(robotcon)
	mess = AddPeakHead(ph ,body)
	return mess
}

func ReceivePRobotServerCmd(mess []byte)(length uint32, ph Packhead, rs PRobotServerCmd ){
	var body []byte
	length,ph,body = PopPeakHead(mess)
	rs , _ = DecodePRobotServerCmdBody(body)
	fmt.Printf("the ph : %+v\n", ph)
	fmt.Printf("the rs : %+v\n", rs)
	return length , ph ,rs
}


func EncodePRegisteredPIBody(rq PRegisteredPI ) (outbyte []byte) {
	body := make([]byte,0)
	outbyte = datastream.AddUint32(rq.Id ,body )
	outbyte = datastream.AddUint32(rq.PIType, outbyte)
	outbyte = datastream.AddString16(rq.PIPass, outbyte)
	return outbyte
}
//used test
func DecodePRegisteredPIBody(inbyte []byte) (rq PRegisteredPI ,err error) {
	body := make([]byte,0)
	rq.Id,body =  datastream.GetUint32(inbyte)
	rq.PIType,body =  datastream.GetUint32(body)
	rq.PIPass,body =  datastream.GetString(body)

	if 0 !=len(body) {
		err = errors.New("解析 PRegisteredPI 失败，解析模具错误")
	}else {
		err = nil
	}
	return rq ,err
}

func DecodePRobotServerCmdBody(inbyte []byte) (rs PRobotServerCmd ,err error) {
	body := make([]byte,0)
	rs.Cmd,body =  datastream.GetUint8(inbyte)
	rs.Sid,body =  datastream.GetUint32(body)
	rs.param1,body =  datastream.GetUint32(body)
	rs.param2,body =  datastream.GetUint32(body)
	if 0 !=len(body) {
		err = errors.New("解析 PRobotServerCmd 失败，解析模具错误")
	}else {
		err = nil
	}
	return rs ,err
}