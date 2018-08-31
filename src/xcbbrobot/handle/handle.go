package handle

import (
	"fmt"
	"time"
	"xcbbrobot/common/datagroove"
	"xcbbrobot/common/message"
	"xcbbrobot/robotcontrol"
	"xcbbrobot/common/maths"
)

//PRobotServerCmd 解包
func Decode_PRobotServerCmd(d *datagroove.DataBuff, i interface{}, length int) {
	var b message.PRobotServerCmd
	b.ReadPackBody(d, length)
	//【INFO】
	fmt.Printf("消息包内容: %+v \n", b)
	var a  *robotcontrol.RobotControl
	a = i.(*robotcontrol.RobotControl)

	fmt.Println("b.Cmd: ",b.Cmd)
	switch b.Cmd {
	case 0:
		fmt.Println("房间: ", b.Sid, "START")
		//创建房间
		a.RoomCreate(b.Sid)

		// 随机时间(0.5s-1s 增加一个) 共 3个机器人
		go func() {
			for iii := 0; iii < 200; iii++ {
				time.Sleep(time.Duration(1E6 * maths.BetweenRand(50, 100)))
				go a.RoomAddRobot(b.Sid)

			}
		}()
	case 1:
		//把房间里面的机器人全部移除
		numSidRobot := a.MapRoomRobotLen(b.Sid)
		fmt.Println( b.Sid,"房间机器人数: ", numSidRobot, " 关闭房间:",b.Sid)
		a.RoomDel(b.Sid)
	case 2:
		fmt.Println("目前没有用")
	case 3:
		fmt.Println("用户进场 ADD_ROBOT")
		//增加一个机器人
		go a.RoomAddRobot(b.Sid)

	case 4:
		fmt.Println("用户出场 REMOVE_ROBOT")
		//减少一个机器人
		numSidRobot := a.MapRoomRobotLen(b.Sid)
		if numSidRobot > 0 {
			go a.RoomSubRobot(b.Sid)
		}
	default:
		fmt.Println("收到 房间服发来的 cmd:", b.Cmd, " 不在正常cmd(0,1,2,3,4)范围")
	}
}
