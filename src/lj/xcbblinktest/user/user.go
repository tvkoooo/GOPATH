//This package imitates the normal behavior of the user
//Robot
package user

import (
	"lj/xcbblinktest/usercome/usercome_sampel"
	"lj/xcbblinktest/tcplink"
	"time"
	"lj/xcbblinktest/userping/userping_sampel"
	"lj/xcbblinktest/userleave/userleave_sampel"
)
type Userrobot struct {
	Ssid uint32
	Version uint32
	Sha1pass string
	Sspass string
	Stampc uint32
	Stamps uint32
}
func userinit()(userinf Userrobot){
	userinf.Ssid = 1
	userinf.Version = 1
	userinf.Sha1pass = ""
	userinf.Sspass = ""
	userinf.Stampc = 0
	userinf.Stamps = 0
	return
}

func Userctrl(ch *chan int,uid uint32,sid uint32)(){
	var userif Userrobot
	userif = userinit()
	conn := tcplink.Tcplink()
	usercome_sampel.Sender(conn,uid ,sid ,userif.Ssid , userif.Version,userif.Sha1pass ,userif.Sspass )
	go usercome_sampel.Recev(conn)

	for i:=0;i<2;i++{
		userping_sampel.Sender(conn,uid,sid ,userif.Stampc,userif.Stamps)
		time.Sleep(5*1E9)
	}
	userleave_sampel.Sender(conn ,uid,sid)
	time.Sleep(3*1E9)
	conn.Close()
	*ch<-1
}