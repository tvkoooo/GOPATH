package main

import (
	//"encoding/json"
	"fmt"
	"lj/messagehead"
	"net"
	"os"
)

//type sendword struct {
//	Pid       uint64
//	Mid       uint64
//	Datawords string
//}

func sender(conn net.Conn) {
	//datasends := &sendword{
	//	10086,
	//	800122001,
	//	"hello world!",
	//}
	var messh messagehead.Messagehead
	var send_unm uint32
	var sendmessage []byte
	messh.Uid = 800122001
	messh.Pid = 10086
	messh.Mid = 999666
	speakto := "hello,fuck you!"
	sendbyte := []byte(speakto)
	len_send := uint32(len(sendbyte))
	sendmessage, send_unm = messagehead.Encodemessh(messh, sendbyte, len_send)
	fmt.Print(messh)
	fmt.Print(speakto)
	fmt.Print(sendbyte)
	fmt.Print(len_send)
	fmt.Print(send_unm)
	fmt.Print(sendmessage)
	fmt.Println("send over\r\n")
	var messh1 messagehead.Messagehead
	var send_unm1 uint32
	var sendmessage1 []byte
	messh1, sendmessage1, send_unm1 = messagehead.Decodemessh(sendmessage, send_unm)
	fmt.Print(messh1)
	fmt.Print(sendmessage1[0:send_unm1])
	fmt.Print(send_unm1)
	str2 := string(sendmessage1[0:send_unm1])
	fmt.Print(str2)
	//datasendj, err := json.Marshal(datasends)
	//if err != nil {
	//	fmt.Print(err.Error())
	//}

	//datarec := &sendword{}
	//err = json.Unmarshal([]byte(datasendj), &datarec)

	//fmt.Print(datasends)
	//fmt.Print(string(datasendj),len(datasendj))
	//fmt.Print(datarec)
	conn.Write([]byte(sendmessage))
	fmt.Println("send over\r\n")

}

func main() {
	server := "127.0.0.1:9090"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)

}
