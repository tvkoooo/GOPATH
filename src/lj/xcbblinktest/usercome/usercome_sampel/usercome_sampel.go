//This package deals with the user leaving the studio
//PRealLeaveChannel  use the test
package usercome_sampel

import (
	"net"
	"fmt"
	"os"
	"lj/xcbblinktest/tcplink"
	"time"
	"lj/xcbblinktest/usercome"
)
//use the default data for example to send
func Sendbody(sid uint32,uid uint32,s_sender string,uid_onmic uint32)(outbyte []byte){
	var mysend usercome.PEnterChannelRQ
	var uri uint32 = (253 << 8) | 2
	mysend.Cmd = "PEnterChannel"
	mysend.Sid = sid
	mysend.Uid = uid
	mysend.Sender = s_sender
	mysend.Uid_onmic = uid_onmic
	outbyte = usercome.ADDsenderbody(uri,mysend)
	return outbyte
}
//Receive data and decode
func Recebody(inbyte []byte)(){
	usercome.Getreceivebody(inbyte)
}

//tcp send
func sender(conn net.Conn,uid uint32,sid uint32,s_sender string,uid_onmic uint32) {
	senddata := make([]byte , 0)
	senddata = Sendbody(uid,sid,s_sender,uid_onmic)
	conn.Write([]byte(senddata))
	fmt.Println(uid," send over",conn.LocalAddr(),"timenow:",time.Now().Format("2006-01-02 15:04:05"),"\r\n")
}

//tcp rec
func recev(conn net.Conn)() {
	recevdata := make([]uint8, 4096)
	for {
		count, err := conn.Read(recevdata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			break
		}
		if count != 0 {
			//fmt.Println("rec Binary stream", recevdata[:count])
			Recebody(recevdata[:count])
			//fmt.Println("rec string ", string(recevdata[:count]))
		}
	}
	//conn.Close()
	//fmt.Println("connect close success LocalAddr:",conn.LocalAddr(),"RemoteAddr",conn.RemoteAddr(),"timenow:",time.Now().Format("2006-01-02 15:04:05"),"\r\n")

}
//user test
func PEnterChannel(uid uint32,sid uint32,s_sender string,uid_onmic uint32,ch *chan int) {

	conn := tcplink.Tcplink()
	sender(conn,uid,sid,s_sender,uid_onmic)
	//time.Sleep(1E9)
	recev(conn)
	time.Sleep(1E9)
	*ch<-1
}
