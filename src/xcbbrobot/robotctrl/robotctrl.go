package robotctrl

import (
	"fmt"
	"net"
	"sync"
	"time"
	"xcbbrobot/common/datagroove"
	"xcbbrobot/common/message"
	"xcbbrobot/common/tcplink"
	"xcbbrobot/config"
	"xcbbrobot/logfile"
	"xcbbrobot/robotvector"
)

const MAXRECEIVEBUFF = 1024 * 5

//线上机器人连接map  sid:robotId:*net.Conn
type SidRobot struct {
	M map[uint32]*net.Conn
	L *sync.RWMutex
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
}

//机器人控制器初始化
func (p *RobotCtrl) RobotCtrlInit() {
	var listPath string
	listPath = "../src/xcbbrobot/robotvector/robot_" + config.Conf.Instance + ".list"
	logfile.GlobalLog.Infoln("listPath:", listPath)
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
		logfile.GlobalLog.Infoln("SidRobotLen:: 房间 sid:", sid, " 不存在")
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
	logfile.GlobalLog.Infoln("RoomCreate::创建新房间 sid:", sid)
	p.LOnlineSid.Lock()
	(p.MOnlineSid)[sid] = &r
	p.LOnlineSid.Unlock()
}

//移除房间
func (p *RobotCtrl) RoomDel(sid uint32) {
	if sidInfo, ok := p.MOnlineSid[sid]; ok {
		if 0 != p.SidRobotLen(sid) {
			for k, v := range sidInfo.M {
				//如果机器人正式入场后，才会有tcp 连接，需要关掉连接
				if v != nil {
					logfile.GlobalLog.Infoln("RoomDel::房间 sid:", sid, "机器人 RobotId:", k, "关闭连接 conn:", v)
					(*v).Close()
					//释放 SidRobot v值 conn 连接为 nil
					v = nil
				}
				// 把所有加入改房间的机器人 robotId 放回 下线机器人列表
				p.MAppRobot.AddRobot(k)
			}
		}
		// 删除 MOnlineSid 元素 k （sid）
		logfile.GlobalLog.Infoln("RoomDel::移除房间 sid:", sid)
		p.LOnlineSid.Lock()
		delete(p.MOnlineSid, sid)
		p.LOnlineSid.Unlock()

	} else {
		logfile.GlobalLog.Infoln("RoomDel::房间 sid:", sid, " 不存在")
	}
}

//清理所有房间
func (p *RobotCtrl) RoomClean() {
	if 0 != p.OnlineSidLen() {
		for k, _ := range p.MOnlineSid {
			p.RoomDel(k)
		}
	}
	logfile.GlobalLog.Infoln("RoomClean:: 清理所有房间")
	p.LOnlineSid.Lock()
	p.MOnlineSid = make(map[uint32]*SidRobot)
	p.LOnlineSid.Unlock()
}

///////////////////////////////////////////////////////////////////////
//在房间 sid 增加一个机器人
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

//在房间 sid 减少一个机器人
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
				rq.t = time.Now().Add(5E8)
				//fmt.Println("SubRobot::房间sid:",sid,"机器人 robotId:",k,"稍后发送离场命令")
				p.LQuit.Lock()
				(p.MQuit)[k] = &rq
				p.LQuit.Unlock()
				//令机器人 延时关闭连接 closeConn，备注，需要比 发送退出命令 QuitRobot  要晚才行
				var rc RobotClose
				rc.sid = sid
				rc.conn = v
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
		//send ping
		if p.OnlineSidLen() > 0 && times%30 == 0 {
			for k, v := range p.MOnlineSid {
				Num := p.SidRobotLen(k)
				if Num > 0 {
					for k2, v2 := range v.M {
						if v2 != nil {
							//sendPing := message.SendPPlusRQ(k2 ,k)
							message.WritePPlusBuff(&sendBuff, k2, k)
							_, err := (*v2).Write(sendBuff.SGroove[sendBuff.LenRemove : sendBuff.LenRemove+sendBuff.LenData])
							if nil != err {
								//fmt.Println("PollOrder::机器人 robotId:",k2,"连接conn:",v2," 发送心跳失败")
								//如果发送失败，说明机器人不在房间，必须去掉该机器人（备注由于发送缓冲区会比该函数慢，因此会出现执行函数时候，连接正常，但是发送数据出现失败）
								p.delRoomRobot(k, k2)
							} else {
								//fmt.Println("PollOrder::机器人 robotId:",k2,"连接conn:",v2,"send sendPing message")
							}
							//无论是否发送成功，清空数据槽，下次使用
							sendBuff.LenRemove = 0
							sendBuff.LenData = 0
						}
					}
				}
				logfile.GlobalLog.Infoln("PollOrder::房间sid:", k, " 当前机器人人数 Num:", Num-p.RobotJoinLen())
			}
		}

		//send delay RealJoinChannel
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
							logfile.GlobalLog.Warnln("PollOrder::机器人 robotId:", k, "连接conn:", v, " 发送入场失败")
							//如果发送失败，说明机器人并未进入房间，必须去掉该机器人
							p.delRoomRobot(v.sid, k)
						} else {
							logfile.GlobalLog.Infoln("PollOrder::机器人 robotId:", k, "连接conn:", &conn, "send RobotJoinLen message 当前房间sid:", v.sid, "机器人总数", p.SidRobotLen(v.sid)-p.RobotJoinLen()+1)
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

		//send delay PRealLeaveChannel
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
							continue
						}
						message.WritePRealLeaveChannelBuff(&sendBuff, k, v.sid)
						_, err := (*v.conn).Write(sendBuff.SGroove[sendBuff.LenRemove : sendBuff.LenRemove+sendBuff.LenData])
						if nil != err {
							logfile.GlobalLog.Warnln("PollOrder::机器人 robotId:", k, "连接conn:", v, " 发送离场失败")
							//如果发送失败，说明机器人并未进入房间，必须去掉该机器人
							p.delRoomRobot(v.sid, k)
						} else {
							logfile.GlobalLog.Infoln("PollOrder::机器人 robotId:", k, "连接conn:", v, "send sendLeave message 当前房间sid:", v.sid, "机器人总数", p.SidRobotLen(v.sid)-p.RobotJoinLen())
						}
						//去掉延时命令
						p.LQuit.Lock()
						delete(p.MQuit, k)
						p.LQuit.Unlock()
						//无论是否发送成功，清空数据槽，下次使用
						sendBuff.LenRemove = 0
						sendBuff.LenData = 0
					}

				}
			}
		}

		//关闭延时后的指定连接
		if p.RobotCloseLen() > 0 {
			for k, v := range p.MClose {
				if time.Now().Sub(v.t) > 0 {
					if v != nil {
						//关闭机器人连接
						if v.conn != nil {
							logfile.GlobalLog.Infoln("RoomDel::房间 sid:", v.sid, "机器人 RobotId:", k, "关闭连接 conn:", v.conn)
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
		//时间间隔 0.2s 进行轮询
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
				}
			}
		}
		logfile.GlobalLog.Infoln("delRoomRobot::异常处理，机器人 robotId:", robotId, "退出房间 sid:", sid)
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
