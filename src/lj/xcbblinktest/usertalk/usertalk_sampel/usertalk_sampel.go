//This package deals with the user talk the studio
//PFeatureRequest (1 << 8) | 23
package usertalk_sampel

import (
	"fmt"
	"lj/xcbblinktest/tcplink"
	"lj/xcbblinktest/userping"
	"lj/xcbblinktest/usertalk"
	"net"
	"os"
	"time"
)

//use the default data for example to send
func Sendbody(uid uint32, singerid uint32, sid uint32, str string) (outbyte []byte) {
	var mysend usertalk.PFeatureRequestRQ
	var uri uint32 = (253 << 8) | 2
	mysend.Cmd = "PTextChat"
	mysend.Uid = uid
	mysend.Singerid = singerid
	mysend.Sid = sid
	mysend.Context = str
	outbyte = usertalk.ADDsenderbody(uri, mysend)
	return outbyte
}

//Receive data and decode
func Recebody(inbyte []byte) {
	userping.Getreceivebody(inbyte)
}

//tcp send
func Sender(conn net.Conn, uid uint32, singerid uint32, sid uint32, str string) {
	senddata := make([]byte, 0)
	senddata = Sendbody(uid, singerid, sid, str)
	num, err := conn.Write([]byte(senddata))
	fmt.Println(uid, " send over", "timenow:", time.Now().Format("2006-01-02 15:04:05"), "num:", num, "\r\n")
	fmt.Println(err)
}

//tcp rec
func Recev(conn net.Conn) {
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
	conn.Close()
	//fmt.Println("connect close success LocalAddr:",conn.LocalAddr(),"RemoteAddr",conn.RemoteAddr(),"timenow:",time.Now().Format("2006-01-02 15:04:05"),"\r\n")

}

//user test
func PFeatureRequest(uid uint32, singerid uint32, sid uint32, str string, ch *chan int) {

	conn := tcplink.Tcplink()
	Sender(conn, uid, singerid, sid, str)
	//time.Sleep(1E9)
	Recev(conn)
	time.Sleep(1E9)
	*ch <- 1
}
