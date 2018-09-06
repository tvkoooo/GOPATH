package robotcontrol

import (
	"fmt"
	"xcbbrobot/robotvector"
	"xcbbrobot/config"
	"sync"
	"net"
	"xcbbrobot/robotonline"
	"time"
)
//房间里的机器人结构    key:房间    k:机器人 robotId   v:机器人生命指针 *robotlife.RobotLife
type MapRoomRobot struct {
	M map[uint32]*net.Conn
	L *sync.RWMutex
}
//应用程序房间结构    key:应用程序    k:房间sid   v:房间里的机器人指针 *MapRoomRobot
type MapAppRoom struct {
	M map[uint32]*MapRoomRobot
	L *sync.RWMutex
}
//机器人控制结构
type RobotControl struct {
	appRobot robotvector.MapAppRobot
	AppRobotOnline robotonline.MapRobotInfo
	appRoom MapAppRoom
}
//机器人控制初始化
func (p *RobotControl)RobotControlInit()() {
	p.appRobot.RobotFreeInit()
	p.AppRobotOnline.RobotFreeInit()
	p.appRobot.LoadRobot(config.Conf.Yaml.Robotlist)
	p.appRoom.M = make(map[uint32]*MapRoomRobot,128)
	p.appRoom.L = new(sync.RWMutex)
}
//创建房间
func (p *RobotControl)RoomCreate(sid uint32 )() {
	var rb MapRoomRobot
	rb.M = make(map[uint32]*net.Conn,128)
	rb.L = new(sync.RWMutex)
	p.appRoom.L.Lock()
	(p.appRoom.M)[sid] = &rb
	p.appRoom.L.Unlock()
}
//移除房间
func (p *RobotControl)RoomDel(sid uint32 )() {
	if 0!=p.MapRoomRobotLen(sid){
		for k, v := range (p.appRoom.M[sid]).M{
			//关掉连接
			(*v).Close()
			//释放 MapRoomRobot v值 conn 连接为 nil
			v = nil
			// 释放 AppRobotOnline v值 conn 连接为 nil
			p.AppRobotOnline.M[k] = nil
			// 删除 AppRobotOnline 元素 k （robotId）
			p.AppRobotOnline.L.Lock()
			delete(p.AppRobotOnline.M, k)
			p.AppRobotOnline.L.Unlock()
			// 把这个机器人 robotId 放回 下线机器人列表
			p.appRobot.AddRobot(k)
		}
	}
	//释放 appRoom v值 conn 连接为 nil
	p.appRoom.M[sid] = nil
	// 删除 appRoom 元素 k （sid）
	p.appRoom.L.Lock()
	delete(p.appRoom.M ,sid)
	p.appRoom.L.Unlock()
}
//清理所有房间
func (p *RobotControl)RoomClean()() {
	p.appRoom.L.Lock()
	p.appRoom.M = make(map[uint32]*MapRoomRobot,128)
	p.appRoom.L.Unlock()
}
//app 房间总数
func (p *RobotControl)MapAppRoomLen()(int) {
	return len(p.appRoom.M)
}
//App 下线机器人总数
func (p *RobotControl)MapAppRobotLen()(int) {
	return p.appRobot.Len()
}
//App 上线机器人总数
func (p *RobotControl)MapAppRobotOnlineLen()(int) {
	return p.AppRobotOnline.Len()
}
//房间 机器人总数
func (p *RobotControl)MapRoomRobotLen(sid uint32 )(l int) {
	if _, ok := p.appRoom.M[sid]; ok {
		l= len((p.appRoom.M[sid]).M)
	} else {
		fmt.Println("房间不存在")
		l= -1
	}
	return l
}

//房间 增加一个机器人
func (p *RobotControl)RoomAddRobot(sid uint32 )() {
	getV ,ok := p.appRoom.M[sid]
	if  ok{
		//空闲机器人列表减少一个机器人 robotId
		robotId := p.appRobot.PopRobot()

		//上线机器人增加 robotId
		con :=p.AppRobotOnline.AddRobot(robotId , sid)

		//房间增加一个机器人
		p.appRoom.L.Lock()
		(*getV).M[robotId] = con
		p.appRoom.L.Unlock()
		fmt.Println(sid,"机器人数量:",p.MapRoomRobotLen(sid))
	}
}

//房间 减少一个机器人
func (p *RobotControl)RoomSubRobot(sid uint32)()  {
	//房间机器人大于 0
	if p.MapRoomRobotLen(sid)>0{

		for k, _ := range (p.appRoom.M[sid]).M{
			//机器人停止工作,在线 map 去掉机器人 k
			p.AppRobotOnline.DelRobot(k , sid)
			
			//房间移除机器人，只是删除了房间map对应sid的机器人寻找指针，此时这个指针在 AppRobotOnline 里面还要继续处理
			p.appRoom.L.Lock()
			delete((p.appRoom.M[sid]).M , k)
			p.appRoom.L.Unlock()
			
			//空闲机器人增加    移除的机器人，本操作需要进行延时，防止删除的机器人还未完全断开连接
			go func() {
				//3s 后，机器人下线
				time.Sleep(3E9)
				p.appRobot.AddRobot(k)
			}()			
			
			//找到一个机器人后就退出搜索
			break
		}

	}
	fmt.Println(sid,"机器人数量:",p.MapRoomRobotLen(sid))
}
//输出房间机器人（数量太大请禁用，只用于小量测试）
func (p *RobotControl)PrintMapRoomRobot(sid uint32)(){
	fmt.Println("PrintMapRoomRobot:",p.appRoom.M)
}
//输出系统机器人（数量太大请禁用，只用于小量测试）
func (p *RobotControl)PrintMapAppRobot()(){
	p.appRobot.PrintRobotMap()
}




