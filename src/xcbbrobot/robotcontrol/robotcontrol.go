package robotcontrol

import (
	"xcbbrobot/robotlife"
	"fmt"
	"xcbbrobot/robotvector"
)
//房间里的机器人结构    key:房间    k:机器人 robotId   v:机器人生命指针 *robotlife.RobotLife
type RoomRobot struct {
	mapRobotRoom map[uint32]*robotlife.RobotLife
}
//应用程序房间结构    key:应用程序    k:房间sid   v:房间里的机器人指针 *RoomRobot
type AppRoom struct {
	mapAppRoom map[uint32]*RoomRobot
}
//机器人控制结构
type RobotControl struct {
	appRobot robotvector.AppRobot
	appRoom AppRoom
}
//机器人控制初始化
func (m *RobotControl)RobotControlInit()() {
	m.appRobot.RobotFreeInit(128*128)
	m.appRobot.LoadRobot("./xcbbrobot/robotvector/rob.list")
	m.appRoom.mapAppRoom = make(map[uint32]*RoomRobot,128)
}
//创建房间
func (m *RobotControl)RoomCreate(sid uint32 )() {
	var rb RoomRobot
	rb.mapRobotRoom = make(map[uint32]*robotlife.RobotLife,128)
	(m.appRoom.mapAppRoom)[sid] = &rb
}
//移除房间
func (m *RobotControl)RoomDel(sid uint32 )() {
	delete(m.appRoom.mapAppRoom ,sid)
}
//清理所有房间
func (m *RobotControl)RoomClean()() {
	m.appRoom.mapAppRoom = make(map[uint32]*RoomRobot,128)
}
//app 房间总数
func (m *RobotControl)AppRoomLen()(int) {
	return len(m.appRoom.mapAppRoom)
}
//App 机器人总数
func (m *RobotControl)AppRobotLen()(int) {
	return m.appRobot.Len()
}
//房间 机器人总数
func (m *RobotControl)RoomRobotLen(sid uint32 )(int) {
	if _, ok := m.appRoom.mapAppRoom[sid]; ok {
		return len((m.appRoom.mapAppRoom[sid]).mapRobotRoom)
	} else {
		fmt.Println("房间不存在")
		return -1
	}

}

//房间 增加一个机器人
func (m *RobotControl)RoomAddRobot(sid uint32 )() {
	//创建一个机器人
	var mr robotlife.RobotLife
	//空闲机器人列表减少一个机器人 robotId
	robotId := m.appRobot.PopRobot()
	//初始化机器人
	mr.RobotInit(robotId,sid)
	//机器人 robotId 进行工作
	go mr.RobotWork()
	//把该机器人的生命交给房间 sid 下面对应机器人 robotId
	(*(m.appRoom.mapAppRoom[sid])).mapRobotRoom[robotId] = &mr
	fmt.Println(sid,"机器人数量:",m.RoomRobotLen(sid))
	fmt.Println(sid,"机器人:")
	m.PrintRoomRobot(sid)
}
//房间 减少一个机器人
func (m *RobotControl)RoomSubRobot(sid uint32)()  {
	//房间机器人大于 0
	if m.RoomRobotLen(sid)>0{
		for k, v := range (m.appRoom.mapAppRoom[sid]).mapRobotRoom{
			//机器人停止工作
			go v.RobotRest()
			//房间移除机器人
			delete((m.appRoom.mapAppRoom[sid]).mapRobotRoom , k)
			//空闲机器人增加    移除的机器人
			m.appRobot.AddRobot(k)
			//找到一个机器人后就退出搜索
			break
		}
	}
	fmt.Println(sid,"机器人数量:",m.RoomRobotLen(sid))
	fmt.Println(sid,"机器人:")
	m.PrintRoomRobot(sid)
}
//输出房间机器人（数量太大请禁用，只用于小量测试）
func (m *RobotControl)PrintRoomRobot(sid uint32)(){
	for k, v := range (m.appRoom.mapAppRoom[sid]).mapRobotRoom{
		fmt.Println("k:", k, "   v:", v)
	}
}
//输出系统机器人（数量太大请禁用，只用于小量测试）
func (m *RobotControl)PrintAppRobot()(){
	fmt.Println("AppRobot : ")
	m.appRobot.PrintRobotMap()
}




