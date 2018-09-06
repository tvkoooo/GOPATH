package handle

import (
	"xcbbrobot/common/datagroove"
	"xcbbrobot/common/message"
	"xcbbrobot/robotctrl"
	"time"
	"xcbbrobot/common/maths"
	"xcbbrobot/logfile"
)

//PRobotServerCmd 解包
func Decode_PRobotServerCmd(d *datagroove.DataBuff, i interface{}, length int) {
	var b message.PRobotServerCmd
	var t time.Duration
	b.ReadPackBody(d, length)
	//【INFO】
	logfile.GlobalLog.Infof("消息包内容: %+v \n", b)
	var a  *robotctrl.RobotCtrl
	a = i.(*robotctrl.RobotCtrl)
	//fmt.Println("b.Cmd: ",b.Cmd)
	switch b.Cmd {
	case 0:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::房间: ", b.Sid, "START")
		//创建房间
		a.RoomCreate(b.Sid)
		//增加机器人
		for iii := 0; iii < 3; iii++ {
			//确定房间还存在情况下进行加人操作
			if _, ok := a.MOnlineSid[b.Sid]; ok {
				t = time.Duration(1E7 * maths.BetweenRand(50,100) * iii)
				a.AddRobot(b.Sid , t)
			}
		}

	case 1:
		//把房间里面的机器人全部移除
		numSidRobot := a.SidRobotLen(b.Sid)
		logfile.GlobalLog.Infoln( b.Sid,"房间机器人数: ", numSidRobot, " 关闭房间:",b.Sid)
		a.RoomDel(b.Sid)
	case 2:
		logfile.GlobalLog.Infoln("目前没有用")
	case 3:
		logfile.GlobalLog.Infoln("用户进场 ADD_ROBOT")
		//增加一个机器人
		t = time.Duration(1E7 * maths.BetweenRand(50,100))
		a.AddRobot(b.Sid ,t)

	case 4:
		logfile.GlobalLog.Infoln("用户出场 REMOVE_ROBOT")
		//减少一个机器人
		numSidRobot := a.SidRobotLen(b.Sid)
		if numSidRobot > 0 {
			a.SubRobot(b.Sid)
		}else {
			logfile.GlobalLog.Infoln("房间sid:",b.Sid,"已经没有机器人")
		}
	default:
		logfile.GlobalLog.Infoln("收到 房间服发来的 cmd:", b.Cmd, " 不在正常cmd(0,1,2,3,4)范围")
	}
}
