//This package imitates the normal behavior of the user
//Robot
package user

import (
	"lj/xcbblinktest/tcplink"
	"lj/xcbblinktest/usercome/usercome_sampel"
	"time"
	//	"lj/xcbblinktest/userping/userping_sampel"
	//	"lj/xcbblinktest/userleave/userleave_sampel"
	"fmt"
	"lj/xcbblinktest/userping/userping_sampel"
	"xcbbrobot/common/message"
)

type Userrobot struct {
	Ssid     uint32
	Version  uint32
	Sha1pass string
	Sspass   string
	Stampc   uint32
	Stamps   uint32
}

func userinit() (userinf Userrobot) {
	userinf.Ssid = 1
	userinf.Version = 1
	userinf.Sha1pass = ""
	userinf.Sspass = ""
	userinf.Stampc = 0
	userinf.Stamps = 0
	return
}

func Userctrl(ch *chan int, uid uint32, sid uint32) {
	var userif Userrobot
	userif = userinit()
	conn := tcplink.Tcplink()
	fmt.Println(uid, "进入", sid, "房间")
	usercome_sampel.Sender(conn, uid, sid, userif.Ssid, userif.Version, userif.Sha1pass, userif.Sspass)

	go usercome_sampel.Recev(conn)
	for i := 0; i < 2; i++ {
		time.Sleep(6 * 1E9)
		userif.Stampc = uint32(time.Now().Unix())
		userif.Stamps = uint32(time.Now().UnixNano() % 1E9)
		fmt.Println(uid, "发送第 ", i+1, " 次 ping")
		userping_sampel.Sender(conn, uid, sid, userif.Stampc, userif.Stamps)
	}
	//userleave_sampel.Sender(conn ,uid,sid)
	time.Sleep(3 * 1E9)

	//发送机器人离场
	sendLeave := message.SendPRealLeaveChannelRQ(uid, sid)
	_, err := conn.Write(sendLeave)
	fmt.Println("send message sendLeave:", sendLeave)
	if nil != err {
		fmt.Println("机器人离场失败")
	}

	time.Sleep(6 * 1E9)
	userif.Stampc = uint32(time.Now().Unix())
	userif.Stamps = uint32(time.Now().UnixNano() % 1E9)
	fmt.Println(uid, "发送第 ", 0000, " 次 ping")
	userping_sampel.Sender(conn, uid, sid, userif.Stampc, userif.Stamps)
	time.Sleep(6 * 1E9)

	fmt.Println(uid, "退出", sid, "房间")
	conn.Close()
	*ch <- 1
}
