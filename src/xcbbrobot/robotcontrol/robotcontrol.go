package robotcontrol

import (
	"xcbbrobot/robotlife"
	"fmt"
	"xcbbrobot/robotvector"
	//"sync"
	"xcbbrobot/config"
)
//房间里的机器人结构    key:房间    k:机器人 robotId   v:机器人生命指针 *robotlife.RobotLife
type MapRoomRobot struct {
	m map[uint32]*robotlife.RobotLife
	//l *sync.RWMutex
}
//应用程序房间结构    key:应用程序    k:房间sid   v:房间里的机器人指针 *MapRoomRobot
type MapAppRoom struct {
	m map[uint32]*MapRoomRobot
	//l *sync.RWMutex
}
//机器人控制结构
type RobotControl struct {
	appRobot robotvector.MapAppRobot
	appRoom MapAppRoom
}
//机器人控制初始化
func (m *RobotControl)RobotControlInit()() {
	m.appRobot.RobotFreeInit(128*128)
	m.appRobot.LoadRobot(config.Conf.Yaml.Robotlist)
	m.appRoom.m = make(map[uint32]*MapRoomRobot,128)
	//m.appRoom.l = new(sync.RWMutex)
}
//创建房间
func (m *RobotControl)RoomCreate(sid uint32 )() {
	var rb MapRoomRobot
	rb.m = make(map[uint32]*robotlife.RobotLife,128)
	//rb.l = new(sync.RWMutex)
	//m.appRoom.l.Lock()
	(m.appRoom.m)[sid] = &rb
	//m.appRoom.l.Unlock()
}
//移除房间
func (m *RobotControl)RoomDel(sid uint32 )() {
	//m.appRoom.l.Lock()
	delete(m.appRoom.m ,sid)
	//m.appRoom.l.Unlock()
}
//清理所有房间
func (m *RobotControl)RoomClean()() {
	//m.appRoom.l.Lock()
	m.appRoom.m = make(map[uint32]*MapRoomRobot,128)
	//m.appRoom.l.Unlock()
}
//app 房间总数
func (m *RobotControl)MapAppRoomLen()(int) {
	return len(m.appRoom.m)
}
//App 机器人总数
func (m *RobotControl)MapAppRobotLen()(int) {
	return m.appRobot.Len()
}
//房间 机器人总数
func (m *RobotControl)MapRoomRobotLen(sid uint32 )(l int) {
	if _, ok := m.appRoom.m[sid]; ok {
		l= len((m.appRoom.m[sid]).m)
	} else {
		fmt.Println("房间不存在")
		l= -1
	}
	return l
}

//房间 增加一个机器人
func (m *RobotControl)RoomAddRobot(sid uint32 )() {
	fmt.Println(sid,"机器人数量:",m.MapRoomRobotLen(sid))
	//创建一个机器人
	var mr robotlife.RobotLife
	//空闲机器人列表减少一个机器人 robotId

	robotId := m.appRobot.PopRobot()
	//初始化机器人

	mr.RobotInit(robotId,sid)
	//机器人 robotId 进行工作

	go mr.RobotWork()
	//把该机器人的生命交给房间 sid 下面对应机器人 robotId

	//m.appRoom.l.Lock()
	(*(m.appRoom.m[sid])).m[robotId] = &mr
	//m.appRoom.l.Unlock()
	fmt.Println(sid,"机器人数量5:",m.MapRoomRobotLen(sid))
}
//房间 减少一个机器人
func (m *RobotControl)RoomSubRobot(sid uint32)()  {
	//房间机器人大于 0
	if m.MapRoomRobotLen(sid)>0{

		for k, v := range (m.appRoom.m[sid]).m{
			//机器人停止工作
			v.RobotRest()
			//房间移除机器人，只是删除了房间map对应sid的机器人寻找指针，此时这个指针在 RobotRest 里面还要继续处理
			//m.appRoom.l.Lock()
			delete((m.appRoom.m[sid]).m , k)
			//m.appRoom.l.Unlock()
			//空闲机器人增加    移除的机器人
			m.appRobot.AddRobot(k)
			//找到一个机器人后就退出搜索
			break
		}

	}
	fmt.Println(sid,"机器人数量:",m.MapRoomRobotLen(sid))
}
//输出房间机器人（数量太大请禁用，只用于小量测试）
func (m *RobotControl)PrintMapRoomRobot(sid uint32)(){
	for k, v := range (m.appRoom.m[sid]).m{
		fmt.Println("k:", k, "   v:", v)
	}
}
//输出系统机器人（数量太大请禁用，只用于小量测试）
func (m *RobotControl)PrintMapAppRobot()(){
	fmt.Println("MapAppRobot : ")
	m.appRobot.PrintRobotMap()
}




