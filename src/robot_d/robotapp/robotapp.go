package robotapp

import (
	"net"
	"syscall"
	"time"
	"robot_d/common/datagroove"
	"robot_d/config"
	"robot_d/common/logfile"
	"robot_d/common/tcplink"
	"robot_d/message"
	"robot_d/handle"
	"robot_d/robotctrl"
)

//发送接收缓冲数据槽 最大容量，超出将断开连接，销毁数据
const GROOVE_CAP_BREAK = 1024 * 8 * 8
//接收缓冲器杯子
const RECEIVE_CUP = 256
//如果发送缓冲槽没有数据发送，返回 -2，跳过本次发送
const SEND_NOTHING = -2


type AppProgram struct {
	RobCtrl  robotctrl.RobotCtrl
	MapUriFunc  message.MapUriFuncDecode
	Conn     net.Conn
	SendBuff datagroove.DataBuff
	RecBuff  datagroove.DataBuff
	NumConnect int
}


func (a *AppProgram) AppInit()(){

	//【程序init】初始化消息 uri 解包函数(注册消息解包函数)
	a.MapUriFunc.UriDecodeHandlerInit()

	//添加uri = (131 << 8) | 2 的解包函数
	//p.M[(131 << 8) | 2] = Decode_PRobotServerCmd
	a.MapUriFunc.ZhuCe((131 << 8) | 2 , handle.Decode_PRobotServerCmd)

	//机器人初始化
	a.RobCtrl.RobotCtrlInit()


	//初始化连接次数 为0
	a.NumConnect = 0
}

func (a *AppProgram) AppRobotConn()( bool)  {
	for {
		//【程序init】tcp （客户端模式） 初始化
		a.Conn = tcplink.Tcplink(config.Conf.ObjectNet)
		if a.Conn != nil {
			break
		}
		time.Sleep(1E8)
	}

	//【程序init】初始化Net 发送和接收数据槽
	a.SendBuff.BufferInit()
	a.RecBuff.BufferInit()
	//【信息拼装与发送】 注册机器人程序 message 进入发送缓冲器
	message.WritePRegisteredPI(&a.SendBuff)

	senNum , err := a.SendMessage()
	if err != nil  {
		logfile.GlobalLog.Fatalln("机器人注册房间服务器失败 err :" ,err)
		return false
	}else {
		if senNum >0 {
			logfile.GlobalLog.Infoln("注册次数:",a.NumConnect,"连接conn:",&a.Conn," AppRobotConn success. LocalAddr: " ,a.Conn.LocalAddr(),"LocalAddr: " ,a.Conn.RemoteAddr(),"发送数据:",senNum)
			a.NumConnect ++
			return true
		}else {
			return false
		}
	}
}

//备注：本函数返回是发送数据和正常socket不一样，无数据发送返回 SEND_NOTHING=-2   关闭连接返回0 正常发送返回发送数据
func (a *AppProgram) SendMessage() (numSendAll int, err error) {
	//如果发送槽里面没有数据要发送，直接返回无数据可发，用于系统轮询发送
	if a.SendBuff.LenData == 0 {
		err = nil
		numSendAll = SEND_NOTHING
		return numSendAll,err
	}
	var numSend int
	numSendAll = 0
	for {
		//【INFO】
		//fmt.Println("SendMessage conn:",a.Conn,"send message:",a.SendBuff.SGroove[a.SendBuff.LenRemove : a.SendBuff.LenRemove+a.SendBuff.LenData])
		if a.Conn !=nil {
			numSend, err = (a.Conn).Write(a.SendBuff.SGroove[a.SendBuff.LenRemove : a.SendBuff.LenRemove+a.SendBuff.LenData])
		}else {
			numSendAll = SEND_NOTHING
			//连接丢失
			err = syscall.ENETRESET
			return numSendAll,err
		}

		if numSend > 0 {
			a.SendBuff.LenRemove += numSend
			a.SendBuff.LenData -= numSend
			numSendAll += numSend
			if 0 == a.SendBuff.LenData {
				break
			} else {
				continue
			}
		}
		if 0 == numSend {
			numSendAll = 0
			//【WARN】
			logfile.GlobalLog.Warnln("SendMessage conn:",a.Conn , "Close() ! Because 0 == numSend")
			if a.Conn != nil {
				a.Conn.Close()
				a.Conn = nil
			}
			break
		}
		if numSend < 0 {
			//等待下次发送
			if err == syscall.EINTR || err == syscall.EWOULDBLOCK || err == syscall.EAGAIN {
				err = nil
				break
			} else {
				numSendAll = 0
				//【WARN】
				logfile.GlobalLog.Warnln("SendMessage conn:",a.Conn , "Close() ! Because numSend < 0")
				if a.Conn != nil {
					a.Conn.Close()
					a.Conn = nil
				}
			}
		}
	}
	return numSendAll, err
}

func (a *AppProgram) ReceiveMessage() (numRecAll int, err error){
	numRecAll = 0
	recNum := 0
	for  {
		a.RecBuff.AddDataACup(RECEIVE_CUP)
		if a.Conn ==nil {
			//连接丢失
			numRecAll = 0
			err = syscall.ENETRESET
			break
		}else {
			recNum, err = a.Conn.Read(a.RecBuff.SGroove[a.RecBuff.LenRemove+a.RecBuff.LenData : a.RecBuff.LenRemove+a.RecBuff.LenData+RECEIVE_CUP])
		}
		if recNum>0 {
			a.RecBuff.LenData += recNum
			numRecAll += recNum
			break
		}
		if recNum == 0 {
			//需要重新启用新连接
			numRecAll = 0
			//【WARN】
			logfile.GlobalLog.Warnln("ReceiveMessage conn:",&a.Conn , "Close() ! Because recNum = 0 err:",err.Error())
			if a.Conn != nil {
				a.Conn.Close()
				a.Conn = nil
			}
			//a.AppRobotConn()
			break
		}
		if recNum < 0 {
			if err == syscall.EINTR || err == syscall.EWOULDBLOCK || err == syscall.EAGAIN{
				err = nil
				break
			}else {
				//【WARN】
				logfile.GlobalLog.Warnln("ReceiveMessage conn:",a.Conn , "Close() ! Because recNum < 0 err:",err.Error())
				//如果是其他异常错误，需要关闭连接，重新建立新的连接（同时清除发送和接收数据槽数据）
				if a.Conn != nil {
					a.Conn.Close()
					a.Conn = nil
				}
				//a.AppRobotConn()
				break
			}
		}
		//【INFO】
		if err != nil {
			if a.Conn != nil {
				logfile.GlobalLog.Infoln("ReceiveMessage Conn:",a.Conn ,a.Conn.LocalAddr(),"numRecAll:",numRecAll,"err:",err.Error())
			}
		}
	}
	return numRecAll ,err
}

//如果数据槽数据不够一个完整包，返回一个 -1 ，需要等待下次查看数据槽数据是否完整
//如果数据槽数据够一个完整消息包，返回消息包长度 length
func (a *AppProgram) CheckMessageLen( ) (length int ){
	length = -1
	if a.RecBuff.LenData >=4 {
		length = int(a.RecBuff.DataSlotReadUint32(a.RecBuff.LenRemove + 0))
		if int(length) > a.RecBuff.LenData {
			length =  -1
		}
	}
	return length
}
//取出消息uri
func (a *AppProgram) GetUri( ) ( uint32  ){
	return a.RecBuff.DataSlotReadUint32(a.RecBuff.LenRemove + 4)
}

//根据 uri 解包
func (a *AppProgram) DecodeMessage() (){
	for{
		length := a.CheckMessageLen()
		//检测数据槽，如果数据槽有足够消息包数据就解包，否则退出
		if length > 0 {
			uri := a.GetUri()
			a.MapUriFunc.L.RLock()
			getMapV , ok := a.MapUriFunc.M[uri]
			a.MapUriFunc.L.RUnlock()
			if ok {
				//打印消息包头
				var ph message.PackHead
				ph.ReadPackHead(&a.RecBuff)
				//【INFO】
				logfile.GlobalLog.Infoln("收到uri: ",uri,"消息:",ph)
				if ph.ResCode == 200 {
					//对这个 uri 安卓map 里面函数进行处理
					getMapV(&a.RecBuff,&a.RobCtrl ,length )
				}else {
					//【WARN】
					logfile.GlobalLog.Warnln("返回包错误 uri:",uri,"ph.ResCode:",ph.ResCode)
				}

			}else {
				//【DEBUG】如果不在 uri 表里面，则只输出 uri ，跳过消息
				//fmt.Println("收到uri:",uri ," ,但没有解包")
				a.RecBuff.DataJump(length)
			}

		}else {
			break
		}
	}
}
