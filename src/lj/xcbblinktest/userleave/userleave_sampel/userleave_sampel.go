//This package deals with the user leaving the studio
//PRealLeaveChannel  use the test
package userleave_sampel

import (
	"fmt"
	"lj/xcbblinktest/tcplink"
	"lj/xcbblinktest/userleave"
	"net"
	"os"
	"time"
)

//use the default data for example to send
func Sendbody(uid uint32, sid uint32) (outbyte []byte) {
	var mysend userleave.PRealLeaveChannelRQ
	var uri uint32 = (360 << 8) | 2
	mysend.Uid = uid
	mysend.Sid = sid
	outbyte = userleave.ADDsenderbody(uri, mysend)
	return outbyte
}

//Receive data and decode
func Recebody(inbyte []byte) (rs userleave.PRealLeaveChannelRS) {
	rs = userleave.Getreceivebody(inbyte)
	return rs
}

//tcp send
func Sender(conn net.Conn, uid uint32, sid uint32) {
	senddata := make([]byte, 0)
	senddata = Sendbody(uid, sid)
	conn.Write([]byte(senddata))
	//fmt.Println(uid," send over","timenow:",time.Now().Format("2006-01-02 15:04:05"),"\r\n")
}

//tcp rec
func Recev(conn net.Conn) (rs userleave.PRealLeaveChannelRS) {
	recevdata := make([]uint8, 4096)
	for {
		count, err := conn.Read(recevdata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			break
		}
		if count != 0 {
			//fmt.Println("rec Binary stream", recevdata[:count])
			rs = Recebody(recevdata[:count])
			//fmt.Println("rec string ", string(recevdata[:count]))
			if rs.Code == 0 {
				break
			}
		}
	}
	conn.Close()
	//fmt.Println("connect close success LocalAddr:",conn.LocalAddr(),"RemoteAddr",conn.RemoteAddr(),"timenow:",time.Now().Format("2006-01-02 15:04:05"),"\r\n")
	return rs
}

//user test
func PRealLeaveChannel(uid uint32, sid uint32, ch *chan int) {

	conn := tcplink.Tcplink()
	Sender(conn, uid, sid)
	//time.Sleep(1E9)
	Recev(conn)
	time.Sleep(1E9)
	*ch <- 1
}
