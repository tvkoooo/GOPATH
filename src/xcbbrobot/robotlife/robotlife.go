package robotlife

import (
	"time"
	"xcbbrobot/common/message"
	"xcbbrobot/config"
	"xcbbrobot/common/tcplink"
	"fmt"
	"net"
	"xcbbrobot/common/maths"
)

const (
	STATECLOSE    = 0
	STATEMOTION   = 1
	STATEFINISH   = 2
)
//机器人生命
type RobotLife struct {
	threadContext chan int
	stateThread int
	uid uint32
	sid uint32
	conn net.Conn
}

func (p *RobotLife)RobotWork() {
	p.lstart()
	p.ljoin()
	p.ldestroy()
}

func (p *RobotLife)RobotRest() {
	p.lshutDown()
}

func (p *RobotLife)RobotInit(uid uint32 , sid uint32)()  {
	p.threadContext = make(chan int)
	p.stateThread = STATECLOSE
	p.uid = uid
	p.sid = sid
}

func (p *RobotLife)ldestroy()  {
	p.threadContext = nil
	p.stateThread = STATECLOSE
}

func (p *RobotLife)lstart()  {
	if STATEFINISH == p.stateThread {
		p.stateThread = STATECLOSE
	}else {
		p.stateThread = STATEMOTION
	}
	go p.loop()
}

func (p *RobotLife)lshutDown()  {
	p.stateThread = STATEFINISH
}

func (p *RobotLife)ljoin()  {
	p.stateThread = <-p.threadContext
}

func (p *RobotLife)recData()  {
	recData := make([]byte, 1024)
	for ;p.stateThread == STATEMOTION ;  {
		lengg , err :=p.conn.Read(recData)
		if err != nil {
			fmt.Println(p.uid ," recData lengg: ",lengg)
		}
	}
}

func (p *RobotLife)loop()  {
	//连接tcp
	conn := tcplink.Tcplink(config.Conf.Server)
	p.conn = conn

	//不处理接收数据

	go p.recData()
	//延时0.5s-1s 以内
	time.Sleep(time.Duration(1E7 * maths.BetweenRand(50,100)))

	//发送机器人入场socket
	sendJoin :=message.SendPRealJoinChannel(p.uid , p.sid)
	_ , err := conn.Write(sendJoin)
	fmt.Println(p.uid,"send sendJoin message:", sendJoin)
	if nil != err{
		fmt.Println(p.uid," 机器人入场失败")
	}

	//发送机器人心跳
	sendTime := 0
	for ; p.stateThread == STATEMOTION; {
		sendTime ++
		if 0 == sendTime % 10  {
			sendPing := message.SendPPlusRQ(p.uid , p.sid)
			_ , err := conn.Write(sendPing)
			//fmt.Println(p.uid,"send sendPing ",p.sid," message:", sendPing)
			if nil != err{
				fmt.Println(p.uid," 发送心跳失败")
			}
		}
		time.Sleep(1E9)
	}
	//延时0.5s-1s 以内 发送机器人离场
	time.Sleep(time.Duration(1E7 * maths.BetweenRand(50,100)))
	//发送机器人离场
	sendLeave :=message.SendPRealLeaveChannelRQ(p.uid , p.sid)
	_ , err = conn.Write(sendLeave)
	fmt.Println(p.uid," send sendLeave message:", sendLeave)
	if nil != err{
		fmt.Println(p.uid," 机器人离场失败")
	}
	time.Sleep(2E9)
	fmt.Println(p.uid , " conn 关闭")
	conn.Close()
	//流程结束
	p.threadContext<-p.stateThread
}

