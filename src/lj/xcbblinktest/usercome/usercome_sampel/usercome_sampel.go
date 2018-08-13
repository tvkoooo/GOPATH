//This package deals with the user leaving the studio
//PRealJoinChannel  use the test
package usercome_sampel

import (
	"lj/xcbblinktest/usercome"
	"net"
	"fmt"
	"os"
	"lj/xcbblinktest/tcplink"
	"time"
)
//use the default data for example to send
func Sendbody(uid uint32,sid uint32,ssid uint32, version uint32,sha1pass string,sspass string)(outbyte []byte){
	var mysend usercome.PRealJoinChannelRQ
	var uri uint32 = (32 << 8) | 2
	mysend.Uid = uid
	mysend.Sid = sid
	mysend.Ssid = ssid
	mysend.Version = version
	mysend.Sha1Pass = sha1pass
	mysend.SsPass = sspass
	outbyte = usercome.ADDsenderbody(uri,mysend)
	return outbyte
}
//Receive data and decode
func Recebody(inbyte []byte)(){
	usercome.Getreceivebody(inbyte)
}

//tcp send
func Sender(conn net.Conn,uid uint32,sid uint32,ssid uint32, version uint32,sha1pass string,sspass string) {
	senddata := make([]byte , 0)
	senddata = Sendbody(uid,sid,ssid,version,sha1pass,sspass)
	conn.Write([]byte(senddata))
	fmt.Println(uid," send over","timenow:",time.Now().Format("2006-01-02 15:04:05"),"\r\n")
}

//tcp rec
func Recev(conn net.Conn)() {
	recevdata := make([]uint8, 4096)
	for {
		count, err := conn.Read(recevdata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			break
		}
		if count != 0 {
			fmt.Println("rec Binary stream", recevdata[:count])
			Recebody(recevdata[:count])
		}
	}
	fmt.Println("connect close success LocalAddr:",conn.LocalAddr(),"RemoteAddr",conn.RemoteAddr(),"timenow:",time.Now().Format("2006-01-02 15:04:05"),"\r\n")

}
//user test
func PRealJoinChannel(uid uint32,sid uint32,ssid uint32, version uint32,sha1pass string,sspass string,ch *chan int) {

	conn := tcplink.Tcplink()
	Sender(conn,uid,sid,ssid,version,sha1pass,sspass)
	time.Sleep(1E9)
	Recev(conn)
	time.Sleep(1E9)
	*ch<-1
}
