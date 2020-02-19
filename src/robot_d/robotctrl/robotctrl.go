package robotctrl

import (
	"net"
	"robot_d/common/datagroove"
	"robot_d/common/fileopr"
	"robot_d/common/logfile"
	"robot_d/common/maths"
	"robot_d/common/tcplink"
	"robot_d/config"
	"robot_d/message"
	"robot_d/robotvector"
	"sync"
	"time"
)

const MAXRECEIVEBUFF = 1024 * 5
const SUBNUMBER = 60
const HURRYSUBNUMBER = 70

//线上机器人连接map  sid:robotId:*net.Conn
//备注：len(SidRobot.M)是共需要那么多机器人进入房间，但是实际进入房间的机器人数量是Num
type SidRobot struct {
	M   map[uint32]*net.Conn
	Num int
	L   *sync.RWMutex
}

//线上 即将发送 RobotJoin MSG 机器人
type RobotJoin struct {
	sid  uint32
	conn *net.Conn
	t    time.Time
}

//线上 即将发送 RobotQuit MSG 机器人
type RobotQuit struct {
	sid  uint32
	conn *net.Conn
	t    time.Time
}

//线上 即将发送 RobotClose MSG 机器人
type RobotClose struct {
	sid  uint32
	conn *net.Conn
	t    time.Time
}

//机器人控制器
type RobotCtrl struct {
	MAppRobot  robotvector.MapAppRobot
	MOnlineSid map[uint32]*SidRobot
	LOnlineSid *sync.RWMutex
	MJoin      map[uint32]*RobotJoin
	LJoin      *sync.RWMutex
	MQuit      map[uint32]*RobotQuit
	LQuit      *sync.RWMutex
	MClose     map[uint32]*RobotClose
	LClose     *sync.RWMutex
	SubNum     int
	HurryNum   int
	Famine     bool
}

//机器人控制器初始化
func (p *RobotCtrl) RobotCtrlInit() {
	var listPath string
	listPath = "D:/work_git/C/xcbb_project.git/server.kr.2/server/feature_new_robot/robot_config/" + config.Conf.ServiceNum + "/" + "robot_" + config.Conf.Instance + ".list"
	logfile.SystemLogPrintln("info", "RobotCtrlInit::Loading robot list files Path:", listPath)
	p.MAppRobot.RobotFreeInit()
	p.MAppRobot.LoadRobot(listPath)
	p.MOnlineSid = make(map[uint32]*SidRobot)
	p.LOnlineSid = new(sync.RWMutex)
	p.MJoin = make(map[uint32]*RobotJoin)
	p.LJoin = new(sync.RWMutex)
	p.MQuit = make(map[uint32]*RobotQuit)
	p.LQuit = new(sync.RWMutex)
	p.MClose = make(map[uint32]*RobotClose)
	p.LClose = new(sync.RWMutex)
	p.SubNum = SUBNUMBER
	p.HurryNum = HURRYSUBNUMBER
	p.Famine = false
}

///////////////////////////////////////////////////////////////////////
//尽量不要使用，只用于测试，查看当前在线房间
func (p *RobotCtrl) PrintOnlineSidMap() {
	logfile.GlobalLog.Debugln("PrintOnlineSidMap:", p.MOnlineSid)
}

//尽量不要使用，只用于测试，查看当前在线房间里面的机器人数据
func (p *RobotCtrl) PrintSidRobotMap(sid uint32) {
	logfile.GlobalLog.Debugln("PrintSidRobotMap:", p.MOnlineSid[sid].M)
}

//即将 发送进入 消息的机器人
func (p *RobotCtrl) PrintRobotJoinMap() {
	logfile.GlobalLog.Infoln("PrintRobotJoinMap:", p.MJoin)
}

//即将 发送退出 消息的机器人
func (p *RobotCtrl) PrintRobotQuitMap() {
	logfile.GlobalLog.Infoln("PrintRobotQuitMap:", p.MQuit)
}

//即将 关闭连接 消息的机器人
func (p *RobotCtrl) PrintRobotCloseMap() {
	logfile.GlobalLog.Infoln("PrintRobotCloseMap:", p.MClose)
}

///////////////////////////////////////////////////////////////////////
//获得本程序所管辖上线房间数量
func (p *RobotCtrl) OnlineSidLen() int {
	return len(p.MOnlineSid)
}

//获得上线机器人  房间sid机器人数量
func (p *RobotCtrl) SidRobotLen(sid uint32) (n int) {
	if _, ok := p.MOnlineSid[sid]; ok {
		n = len(p.MOnlineSid[sid].M)
	} else {
		logfile.GlobalLog.Debugln("SidRobotLen::The room sid:", sid, " does not exist.")
		n = 0
	}
	return n
}

//获得上线机器人  即将入场数量
func (p *RobotCtrl) RobotJoinLen() int {
	return len(p.MJoin)
}

//获得上线机器人  即将出场数量
func (p *RobotCtrl) RobotQuitLen() int {
	return len(p.MQuit)
}

//获得上线机器人  即将断开连接数量
func (p *RobotCtrl) RobotCloseLen() int {
	return len(p.MClose)
}

///////////////////////////////////////////////////////////////////////
//创建房间
func (p *RobotCtrl) RoomCreate(sid uint32) {
	var r SidRobot
	r.M = make(map[uint32]*net.Conn)
	r.L = new(sync.RWMutex)
	r.Num = 0
	logfile.GlobalLog.Infoln("RoomCreate::Create a new room sid:", sid)
	p.LOnlineSid.Lock()
	(p.MOnlineSid)[sid] = &r
	p.LOnlineSid.Unlock()
}

//移除房间
func (p *RobotCtrl) RoomDel(sid uint32) {
	if sidInfo, ok := p.MOnlineSid[sid]; ok {
		if p.SidRobotLen(sid) > 0 {
			for k, v := range sidInfo.M {
				//如果机器人正式入场后，才会有tcp 连接，需要关掉连接
				if v != nil {
					logfile.GlobalLog.Debugln("RoomDel::Room Sid :", sid, "RobotId:", k, "Disconnect:", v)
					(*v).Close()
					//释放 SidRobot v值 conn 连接为 nil
					v = nil
				}
				//无论是否有连接，房间里面的机器人都需要从加人列表当做移除，如果已经移除可以再移除一次
				p.LJoin.Lock()
				delete(p.MJoin, k)
				p.LJoin.Unlock()

				// 把所有加入改房间的机器人 robotId 放回 下线机器人列表
				p.MAppRobot.AddRobot(k)
			}

			logfile.GlobalLog.Infoln("RoomDel::Remove the room sid:", sid, "Disconnect robot num:", sidInfo.Num, "leave the field robot num:", p.SidRobotLen(sid))
			//因为是直接关播房间，机器人是被直接提出房间，无须发送离场命令,清退所有机器人后房间人数清零
			sidInfo.Num = 0
		}
		// 删除 MOnlineSid 元素 k （sid）
		p.LOnlineSid.Lock()
		delete(p.MOnlineSid, sid)
		p.LOnlineSid.Unlock()

	} else {
		logfile.GlobalLog.Infoln("RoomDel::the room sid:", sid, "does not exist.")
	}
}

//清理所有房间
func (p *RobotCtrl) RoomClean() {
	if 0 != p.OnlineSidLen() {
		for k, _ := range p.MOnlineSid {
			p.RoomDel(k)
		}
	}
	logfile.GlobalLog.Infoln("RoomClean:: Clear all rooms")
	p.LOnlineSid.Lock()
	p.MOnlineSid = make(map[uint32]*SidRobot)
	p.LOnlineSid.Unlock()
}

///////////////////////////////////////////////////////////////////////
//在房间 sid 增加一个机器人（房间开启一个加人命令，但还未正式入场）
func (p *RobotCtrl) AddRobot(sid uint32, t time.Duration) {
	//确定该房间存在情况下
	if sidInfo, ok := p.MOnlineSid[sid]; ok {
		//空闲机器人列表减少一个机器人 robotId
		robotId := p.MAppRobot.PopRobot()
		//上线机器人 房间 sid 增加一个机器人 robotId
		sidInfo.L.Lock()
		sidInfo.M[robotId] = nil
		sidInfo.L.Unlock()
		//打印房间机器人数量
		//fmt.Println("AddRobot::房间sid:",sid,"机器人数量:",p.SidRobotLen(sid))
		//令机器人 延时发送 入场命令
		var r RobotJoin
		r.sid = sid
		//r.t = time.Now().Add(time.Duration(1E7 * maths.BetweenRand(50,100)))
		r.t = time.Now().Add(t)
		//fmt.Println("AddRobot::房间sid:",sid,"机器人 robotId:",robotId,"稍后发送入场命令")
		p.LJoin.Lock()
		(p.MJoin)[robotId] = &r
		p.LJoin.Unlock()
	}
}

//在房间 sid 减少一个机器人（房间开启一个减人命令，但还未正式离场和关闭连接）
func (p *RobotCtrl) SubRobot(sid uint32) {
	//确定该房间存在情况下
	if sidInfo, ok := p.MOnlineSid[sid]; ok {
		//房间确实有机器人
		if p.SidRobotLen(sid) > 0 {
			for k, v := range sidInfo.M {
				//令机器人 延时发送 退出命令 QuitRobot
				var rq RobotQuit
				rq.sid = sid
				rq.conn = v
				t := time.Duration(1E7 * maths.BetweenRand(10, 49))
				rq.t = time.Now().Add(t)
				//fmt.Println("SubRobot::房间sid:",sid,"机器人 robotId:",k,"稍后发送离场命令")
				p.LQuit.Lock()
				(p.MQuit)[k] = &rq
				p.LQuit.Unlock()
				//令机器人 延时关闭连接 closeConn，备注，需要比 发送退出命令 QuitRobot  要晚才行
				var rc RobotClose
				rc.sid = sid
				rc.conn = v
				t = time.Duration(1E7 * maths.BetweenRand(50, 100))
				rc.t = time.Now().Add(1E9)
				//fmt.Println("SubRobot::房间sid:",sid,"机器人 robotId:",k,"稍后关闭连接conn:",v)
				p.LClose.Lock()
				(p.MClose)[k] = &rc
				p.LClose.Unlock()
				//把该机器人移除线上列表
				sidInfo.L.Lock()
				delete(sidInfo.M, k)
				sidInfo.L.Unlock()
				//把机器人放回线下
				p.MAppRobot.AddRobot(k)
				//找到一个机器人后可以退出流程
				break
			}
		}
	}
}

//轮询所有机器人延时命令，触发延时命令
func (p *RobotCtrl) PollOrder() {
	times := 0
	var sendBuff datagroove.DataBuff
	sendBuff.BufferInit()
	for {
		times++
		//系统10s(0.2 * 50)做一次状态维护 send ping
		if p.OnlineSidLen() > 0 && times%50 == 0 {
			for sid, sidInfo := range p.MOnlineSid {
				Num := p.SidRobotLen(sid)
				if Num > 0 {
					for robotId, robotInfo := range sidInfo.M {
						if robotInfo != nil {
							//sendPing := message.SendPPlusRQ(k2 ,k)
							message.WritePPlusBuff(&sendBuff, robotId, sid)
							_, err := (*robotInfo).Write(sendBuff.SGroove[sendBuff.LenRemove : sendBuff.LenRemove+sendBuff.LenData])
							if nil != err {
								//fmt.Println("PollOrder::机器人 robotId:",k2,"连接conn:",v2," 发送心跳失败")
								//如果发送失败，说明机器人不在房间，必须去掉该机器人（备注由于发送缓冲区会比该函数慢，因此会出现执行函数时候，连接正常，但是发送数据出现失败）
								p.delRoomRobot(sid, robotId)
								//同时在send ping表中删除该机器人
								sidInfo.L.Lock()
								delete(sidInfo.M, robotId)
								sidInfo.L.Unlock()

							} else {
								//fmt.Println("PollOrder::机器人 robotId:",k2,"连接conn:",v2,"send sendPing message")
							}
							//无论是否发送成功，清空数据槽，下次使用
							sendBuff.LenRemove = 0
							sendBuff.LenData = 0
						}
					}
				}
				//输出房间当前人数
				logfile.GlobalLog.Infoln("PollOrder::sid:", sid, "Robot number:", sidInfo.Num)
				//如果该房间人数太多，适当安排机器人离开房间,70开始离场，80加速
				if sidInfo.Num >= p.SubNum {
					p.SubRobot(sid)
					if sidInfo.Num >= p.HurryNum {
						p.SubRobot(sid)
					}
				}
			}
			//输出本程序还剩余的机器人数量
			OfflineNum := p.MAppRobot.Len()
			allNum := p.MAppRobot.Num
			percent := float32(float32(OfflineNum) / float32(allNum))

			//考虑日志太多，只有在机器人总数低于一定情况才反复输出。
			if percent < 0.3 {
				logfile.GlobalLog.Warnln("PollOrder::Offline robot num:", OfflineNum, "all robot num:", allNum, "Offline robot percent:", percent)
			} else {
				if times%300 == 0 {
					logfile.GlobalLog.Infoln("PollOrder::Offline robot num:", OfflineNum, "all robot num:", allNum, "Offline robot percent:", percent)
				}
			}

			//如果线下机器人不足20%，机器人处于饥荒状态，需要加速房间机器人离场。当机器人回复50%，则可以正常进行
			if p.Famine == false && percent < 0.2 {
				p.SubNum = 40
				p.HurryNum = 50
				p.Famine = true
			}
			if p.Famine == true && percent > 0.5 {
				p.SubNum = SUBNUMBER
				p.HurryNum = HURRYSUBNUMBER
				p.Famine = false
			}

			//判断日志文件大小，决定是否新建新文件
			fileSize := fileopr.CheckFileSize(&logfile.FunLogPath, logfile.MAXLOGSIZE)
			if -1 == fileSize || 0 == fileSize {
				logfile.SystemLogPrintln("SYSTEM", "file:", logfile.FunLogPath, "Size:", fileSize, "will create new.")
				logfile.GlobalLog.LogFileClosed()
				logfile.GlobalLog.LogFileOpen(&logfile.FunLogPath, fileSize)

			} else {
				//logfile.SystemLogPrintln("SYSTEM","file:",logfile.FunLogPath,"Size:",fileSize)
			}
		}

		//系统检测入场包 send delay RealJoinChannel
		if p.RobotJoinLen() > 0 {
			for k, v := range p.MJoin {
				if time.Now().Sub(v.t) > 0 {
					//sendJoin :=message.SendPRealJoinChannel(k , v.sid)
					//确定房间还存在情况下进行加人操作
					if sidInfo, ok := p.MOnlineSid[v.sid]; ok {
						//连接tcp
						conn := tcplink.Tcplink(config.Conf.ObjectNet)
						//因为确定进入房间，给线上机器人正式 tcp 连接，同时放入 ping /leave/close/ map 当中
						sidInfo.L.Lock()
						sidInfo.M[k] = &conn
						sidInfo.L.Unlock()
						message.WritePRealJoinChannelBuff(&sendBuff, k, v.sid)
						_, err := conn.Write(sendBuff.SGroove[sendBuff.LenRemove : sendBuff.LenRemove+sendBuff.LenData])
						if nil != err {
							logfile.GlobalLog.Warnln("PollOrder::robotId:", k, "conn:", v, "Sending PRealJoinChannel failed")
							//如果发送失败，说明机器人并未进入房间，必须去掉该机器人
							p.delRoomRobot(v.sid, k)
						} else {
							sidInfo.Num += 1
							logfile.GlobalLog.Debugln("PollOrder::robotId:", k, "conn:", &conn, "Sending PRealJoinChannel success .Now sid:", v.sid, "Robot number:", sidInfo.Num)
						}
						//无论是否发送成功，清空数据槽，下次使用
						sendBuff.LenRemove = 0
						sendBuff.LenData = 0
					}
					//去掉延时命令
					p.LJoin.Lock()
					delete(p.MJoin, k)
					p.LJoin.Unlock()
				}
			}
		}

		//系统检测离场包 send delay PRealLeaveChannel
		if p.RobotQuitLen() > 0 {
			for k, v := range p.MQuit {
				if time.Now().Sub(v.t) > 0 {
					if v != nil {
						//发送机器人离场
						//sendLeave :=message.SendPRealLeaveChannelRQ(k , v.sid)
						//如果连接已经断开，无需再发送离场命令
						if v.conn == nil {
							//如果机器人已经断开连接，需要把机器人放回线下
							p.delRoomRobot(v.sid, k)
						} else {
							//确定房间还存在
							if sidInfo, ok := p.MOnlineSid[v.sid]; ok {
								message.WritePRealLeaveChannelBuff(&sendBuff, k, v.sid)
								_, err := (*v.conn).Write(sendBuff.SGroove[sendBuff.LenRemove : sendBuff.LenRemove+sendBuff.LenData])
								if nil != err {
									logfile.GlobalLog.Warnln("PollOrder::robotId:", k, "conn:", v, "Sending PRealLeaveChannel failed")
									//如果发送失败，说明机器人并未进入房间，必须去掉该机器人
									p.delRoomRobot(v.sid, k)
								} else {
									sidInfo.Num -= 1
									logfile.GlobalLog.Debugln("PollOrder::robotId:", k, "conn:", &v.conn, "send PRealLeaveChannel success .Now sid:", v.sid, "Robot number:", sidInfo.Num)
								}
								//无论是否发送成功，清空数据槽，下次使用
								sendBuff.LenRemove = 0
								sendBuff.LenData = 0
							}
						}
						//去掉延时命令
						p.LQuit.Lock()
						delete(p.MQuit, k)
						p.LQuit.Unlock()
					}
				}
			}
		}

		//系统检测关闭连接 关闭延时后的指定连接
		if p.RobotCloseLen() > 0 {
			for k, v := range p.MClose {
				if time.Now().Sub(v.t) > 0 {
					if v != nil {
						//关闭机器人连接
						if v.conn != nil {
							logfile.GlobalLog.Debugln("PollOrder::sid:", v.sid, "RobotId:", k, "close conn:", &v.conn)
							(*v.conn).Close()
							v.conn = nil
							v = nil
						}
						//去掉延时命令
						p.LClose.Lock()
						delete(p.MClose, k)
						p.LClose.Unlock()
					}
				}
			}
		}

		//小循环间隔周期 时间间隔 0.2s 进行轮询
		time.Sleep(2E8)
	}
}

//失败处理，内部函数，去掉房间一个机器人
func (p *RobotCtrl) delRoomRobot(sid uint32, robotId uint32) {
	//确定该房间存在情况下
	if sidInfo, ok := p.MOnlineSid[sid]; ok {
		//房间确实有机器人
		if p.SidRobotLen(sid) > 0 {
			//房间里面确实存在该机器人
			if robot, ok := sidInfo.M[robotId]; ok {
				//如果还要连接，关闭连接
				if robot != nil {
					(*robot).Close()
					robot = nil

					logfile.GlobalLog.Infoln("delRoomRobot::Special execution operation .robotId:", robotId, "Exit the room sid:", sid)
				}
				//无论是否有连接，这个机器人都需要从加人列表当做移除，如果已经移除可以再移除一次
				p.LJoin.Lock()
				delete(p.MJoin, robotId)
				p.LJoin.Unlock()
			}
		}
		//把该机器人放回下线机器人（如果之前已经放回，可以再次覆盖放入）
		p.MAppRobot.AddRobot(robotId)
	}
}

func (p *RobotCtrl) PollReceive() {
	buffer := make([]byte, MAXRECEIVEBUFF)
	for {
		if p.OnlineSidLen() > 0 {
			for k, v := range p.MOnlineSid {
				Num := p.SidRobotLen(k)
				if Num > 0 {
					for _, conn := range v.M {
						if conn != nil {
							for {
								(*conn).SetReadDeadline(time.Now().Add(1E3))
								reNum, _ := (*conn).Read(buffer)
								//fmt.Println("robotId",robotId,"PollReceive::numm:",reNum,"err",err,"timenow2:",time.Now())
								if reNum != MAXRECEIVEBUFF {
									break
								}
							}
						}
					}
				}
			}
		}
		time.Sleep(6E9)
	}
}
