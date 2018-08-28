package message

import (
	"xcbbrobot/common/datastream"
	"xcbbrobot/common/typechange"
	"encoding/json"
)

//PEnterChannelRQ RQ DATA  (1 << 8) | 23
type PEnterChannlRQ struct {
	Str32 string
}

type PEnterRQJson struct {
	Cmd string `json:"cmd"`
	Uid string `json:"uid"`
	Sender string `json:"sender"`
	Uid_onmic string `json:"uid_onmic"`
	Sid string `json:"sid"`
	Socket_id string `json:"socket_id"`
}


func SendPEnterChannel(robotId uint32 ,uid uint32 , sid uint32 )( mess []byte){
	var ph PackHead
	ph.Uri = (1 << 8) | 23
	ph.Sid = 0
	ph.ResCode = 200
	ph.Tag = 1

	var robotEnterChanlJson PEnterRQJson
	robotEnterChanlJson.Cmd = "PEnterChannel"
	robotEnterChanlJson.Uid = typechange.Int2StringRe(int(robotId))
	robotEnterChanlJson.Sender = "niMeiDe"
	robotEnterChanlJson.Uid_onmic = typechange.Int2StringRe(int(uid))
	robotEnterChanlJson.Sid = typechange.Int2StringRe(int(sid))
	robotEnterChanlJson.Socket_id = "0"


	jsonSend ,_ := json.Marshal(robotEnterChanlJson)

	var sendPEnterChannlRQ PEnterChannlRQ
	sendPEnterChannlRQ.Str32 = string(jsonSend)

	body:= EncodePEnterChannel(sendPEnterChannlRQ)
	mess = AddPeakHead(ph ,body)
	return mess
}


func EncodePEnterChannel(rq PEnterChannlRQ ) (outbyte []byte) {
	body := make([]byte,0)
	outbyte = datastream.AddString32(rq.Str32, body)
	return outbyte
}
