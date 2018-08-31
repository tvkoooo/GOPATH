package robotonline

import (
	"fmt"
	"net"
	"xcbbrobot/config"
	"xcbbrobot/common/tcplink"
	"sync"
	"time"
	"xcbbrobot/common/maths"
	"xcbbrobot/common/message"
)

type RobotInfo struct {
	sid uint32
	conn *net.Conn
}
type MapRobotInfo struct {
	M map[uint32]*RobotInfo
	L *sync.RWMutex
}

func (p *MapRobotInfo)RobotFreeInit(num int){
	p.M = make(map[uint32]*RobotInfo , num)
	p.L = new(sync.RWMutex)
}
//尽量不要使用，只用于测试
func (p *MapRobotInfo)PrintRobotMap()(){
	fmt.Println("map MapRobotInfo:",p.M)
}
//获得上线机器人数量
func (p *MapRobotInfo)Len()(int ){
	return  len(p.M)
}

func (p *MapRobotInfo)AddRobot(robotId uint32 , sid uint32) (*net.Conn) {

	//连接tcp
	conn := tcplink.Tcplink(config.Conf.ObjectNet)

	var r RobotInfo
	r.sid = sid
	r.conn = &conn

	p.L.Lock()
	(p.M)[robotId] = &r
	p.L.Unlock()

	//延时执行后释放 go runtime
	go func(c *net.Conn) {
		//延时0.5s-1s 以内
		time.Sleep(time.Duration(1E7 * maths.BetweenRand(50,100)))

		if nil !=c {
			//发送机器人入场socket
			sendJoin :=message.SendPRealJoinChannel(robotId , sid)
			_ , err := (*c).Write(sendJoin)
			fmt.Println(robotId,"send sendJoin message:", sendJoin)
			if nil != err{
				fmt.Println(robotId," 机器人入场失败:",err.Error())
				r.conn = nil
				//由于机器人入场会比控制命令滞后一端时间，在此期间如果房间消失，机器人会入场失败。因此需要确定房间还存在才进入
			}
		}

	}(&(*(p.M)[robotId].conn))
	return &conn
}

func (p *MapRobotInfo)DelRobot(robotId uint32 , sid uint32)  {

	//延时执行后释放 go runtime
	go func() {
		time.Sleep(1E9)
		//发送机器人离场
		sendLeave :=message.SendPRealLeaveChannelRQ(robotId , sid)
		sendNum , err :=(*p.M[robotId].conn).Write(sendLeave)
		fmt.Println(robotId," send sendLeave message:", sendLeave)
		if nil != err{
			fmt.Println(robotId," 机器人离场失败","sendNum:",sendNum)
		}
		time.Sleep(1E9)
		fmt.Println(robotId , " conn 关闭")
		(*p.M[robotId].conn).Close()

		p.L.Lock()
		p.M[robotId] = nil
		delete(p.M, robotId)
		p.L.Unlock()
	}()

}
func (p *MapRobotInfo)CleanRobot()  {
	p.L.Lock()
	p.M = make(map[uint32]*RobotInfo)
	p.L.Unlock()
}

func (p *MapRobotInfo)SendPing()  {

	for{
		if 0!=len(p.M){
			p.L.RLock()
			for k, v := range p.M {
				sendPing := message.SendPPlusRQ(k , v.sid)
				_ , err := (*v.conn).Write(sendPing)
				fmt.Println(k,"send sendPing ",v.sid," message:", sendPing)
				if nil != err{
					fmt.Println(k," 发送心跳失败")
				}
			}
			p.L.RUnlock()
		}
		time.Sleep(6E9)
	}
}