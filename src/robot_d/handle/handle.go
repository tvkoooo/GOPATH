package handle

import (
	"robot_d/common/datagroove"
	"robot_d/message"
	"robot_d/robotctrl"
	"time"
	"robot_d/common/maths"
	"robot_d/common/logfile"
)

//PRobotServerCmd 解包
func Decode_PRobotServerCmd(d *datagroove.DataBuff, i interface{}, length int) {
	var b message.PRobotServerCmd
	var t time.Duration
	b.ReadPackBody(d, length)
	//【INFO】
	logfile.GlobalLog.Infof("Decode_PRobotServerCmd::Message package content: %+v \n", b)
	var a  *robotctrl.RobotCtrl
	a = i.(*robotctrl.RobotCtrl)
	//fmt.Println("b.Cmd: ",b.Cmd)
	switch b.Cmd {
	case 0:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:", b.Sid, "START")
		//创建房间
		a.RoomCreate(b.Sid)
		//在6s 内 迅速进入10人
		for iii := 0; iii < 10; iii++ {
			//确定房间还存在情况下进行加人操作
			if _, ok := a.MOnlineSid[b.Sid]; ok {
				t = time.Duration(1E7 * maths.BetweenRand(10,60) * iii)
				a.AddRobot(b.Sid , t)
			}
		}
		//6s 过后 4分钟内随机缓慢再增加 40人
		for iii := 0; iii < 40; iii++ {
			//确定房间还存在情况下进行加人操作
			if _, ok := a.MOnlineSid[b.Sid]; ok {
				t = time.Duration(6E7 * maths.BetweenRand(50,100) * iii + 6E9)
				a.AddRobot(b.Sid , t)
			}
		}

	case 1:
		//把房间里面的机器人全部移除
		numSidRobot := a.SidRobotLen(b.Sid)
		logfile.GlobalLog.Infoln( b.Sid,"Decode_PRobotServerCmd::sid all robot num:", numSidRobot, "close sid:",b.Sid)
		a.RoomDel(b.Sid)
	case 2:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::No use at present.")
	case 3:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::User coming, ADD_ROBOT")
		//增加一个机器人
		t = time.Duration(1E7 * maths.BetweenRand(50,100))
		a.AddRobot(b.Sid ,t)

	case 4:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::User out, REMOVE_ROBOT")
		//减少一个机器人
		numSidRobot := a.SidRobotLen(b.Sid)
		if numSidRobot > 0 {
			a.SubRobot(b.Sid)
		}else {
			logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::sid:",b.Sid,"has no robot")
		}
	default:
		logfile.GlobalLog.Infoln("Decode_PRobotServerCmd::the cmd:", b.Cmd, " not in(0,1,2,3,4)")
	}
}
