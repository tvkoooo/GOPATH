//This package deals with the user ping the studio
//PPlus (12 << 8) | 4
package userping_sampel

import (
	"fmt"
	"lj/xcbblinktest/tcplink"
	"lj/xcbblinktest/userping"
	"net"
	"os"
	"time"
)

//use the default data for example to send
func Sendbody(uid uint32, sid uint32, stampc uint32, stamps uint32) (outbyte []byte) {
	var mysend userping.PPlusRQ
	var uri uint32 = (12 << 8) | 4
	mysend.Uid = uid
	mysend.Sid = sid
	mysend.Stampc = stampc
	mysend.Stamps = stamps
	outbyte = userping.ADDsenderbody(uri, mysend)
	return outbyte
}

//Receive data and decode
func Recebody(inbyte []byte) {
	userping.Getreceivebody(inbyte)
}

//tcp send
func Sender(conn net.Conn, uid uint32, sid uint32, stampc uint32, stamps uint32) {
	senddata := make([]byte, 0)
	senddata = Sendbody(uid, sid, stampc, stamps)
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
func PPlus(uid uint32, sid uint32, stampc uint32, stamps uint32, ch *chan int) {

	conn := tcplink.Tcplink()
	Sender(conn, uid, sid, stampc, stamps)
	//time.Sleep(1E9)
	Recev(conn)
	time.Sleep(1E9)
	*ch <- 1
}
