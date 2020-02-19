package handle

import (
	"robot_d/common/datagroove"
	"robot_d/common/logfile"
	"robot_d/common/maths"
	"robot_d/message"
	"robot_d/robotctrl"
	"time"
)

//PRobotServerCmd 解包
func Decode_PRobotServerCmd(d *datagroove.DataBuff, i interface{}, length int) {
	var b message.PRobotServerCmd
	var t time.Duration
	b.ReadPackBody(d, length)
	//【INFO】
	logfile.GlobalLog.Debugf("Decode_PRobotServerCmd::Message package content: %+v \n", b)
	var a *robotctrl.RobotCtrl
	a = i.(*robotctrl.RobotCtrl)
	//fmt.Println("b.Cmd: ",b.Cmd)
	switch b.Cmd {
	case 0:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:", b.Sid, "RSC_LIVE_START")
		//创建房间
		a.RoomCreate(b.Sid)
		//在6s 内 迅速进入10人
		for iii := 0; iii < 10; iii++ {
			//确定房间还存在情况下进行加人操作
			if _, ok := a.MOnlineSid[b.Sid]; ok {
				t = time.Duration(1E7 * maths.BetweenRand(10, 60) * iii)
				a.AddRobot(b.Sid, t)
			}
		}
		//6s 过后 4分钟内随机缓慢再增加 40人
		for iii := 0; iii < 40; iii++ {
			//确定房间还存在情况下进行加人操作
			if _, ok := a.MOnlineSid[b.Sid]; ok {
				t = time.Duration(6E7*maths.BetweenRand(50, 100)*iii + 6E9)
				a.AddRobot(b.Sid, t)
			}
		}

	case 1:
		//把房间里面的机器人全部移除
		//确定该房间存在情况下
		if sidInfo, ok := a.MOnlineSid[b.Sid]; ok {
			logfile.GlobalLog.Infoln(b.Sid, "Decode_PRobotServerCmd::sid:", b.Sid, "RSC_LIVE_STOP , All robot num:", a.SidRobotLen(b.Sid), "Real robot num:", sidInfo.Num, "will leave the field")
			a.RoomDel(b.Sid)
		}
	case 2:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::RSC_LIVE_END No use at present.")
	case 3:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:", b.Sid, "RSC_ADD_ROBOT")
		//确定该房间存在情况下
		if _, ok := a.MOnlineSid[b.Sid]; ok {
			t = time.Duration(1E7 * maths.BetweenRand(50, 100))
			//增加一个机器人
			a.AddRobot(b.Sid, t)
		} else {
			logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:", b.Sid, "Not yet created,please wait .")
			//因为当前房间不存在，需要先创建一个房间，然后特殊处理，只是缓慢增加40人
			//创建房间
			a.RoomCreate(b.Sid)
			//缓慢增加40人
			for iii := 0; iii < 40; iii++ {
				//确定房间还存在情况下进行加人操作
				if _, ok := a.MOnlineSid[b.Sid]; ok {
					t = time.Duration(6E7 * maths.BetweenRand(50, 100) * iii)
					a.AddRobot(b.Sid, t)
				}
			}
		}
	case 4:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:", b.Sid, "RSC_REMOVE_ROBOT")
		//减少一个机器人
		numSidRobot := a.SidRobotLen(b.Sid)
		if numSidRobot > 0 {
			//低于最低人数，机器人不离开房间
			if numSidRobot < 30 {
				logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:", b.Sid, "robot num:", numSidRobot, "Robots are too few to order.")
			} else {
				a.SubRobot(b.Sid)
			}
		} else {
			logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:", b.Sid, "has no robot")
		}
	default:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:", b.Sid, "the cmd:", b.Cmd, " not in(0,1,2,3,4)")
	}
}
