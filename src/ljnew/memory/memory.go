package main

import (
	"fmt"
	"ljnew/loglj"
	"ljnew/readfile"
	"ljnew/typechange"
	"time"
)

// System constant
const (
	MAXROBOTNUM     = 128 * 128                                       //MAX robot number int this robot system
	MAXROOMROBOTNUM = 128                                             //MAX robot number int singer room
	MAXSINGERNUM    = 128                                             //MAX singer number can use this robot system
	LOGLEVEL        = loglj.L_DEBUG                                   //Use the DEBUG log level
	LOGPATH         = "D:/GOPATH/resources/ljnew/loglj/"              //Log path
	ROBOTLISTPATH   = "D:/GOPATH/resources/ljnew/memory/robottt.list" //Robot list path
)

// System variable
var (
	logFile            *loglj.LogFileName             //    logfile
	mapRobotFree       map[int32]bool                 //    k:robotid    v:ture
	mapSingerRoomRobot map[int32]int64                //    k:robotid    v:unix time(int64)
	mapOnlineSinger    map[int32](*(map[int32]int64)) //    k:singerid   v:mapSingerRoomRobot pt(Pointer)
	allRobotnum        int                            //    Actual number of robots

)

// System initialization
func init() {
	mapRobotFree = make(map[int32]bool, MAXROBOTNUM)
	mapOnlineSinger = make(map[int32](*(map[int32]int64)), MAXSINGERNUM)
	logFile, _ = loglj.LogFileOpen(LOGPATH + time.Now().Format("20060102") + ".log")
	loglj.SetLoglevel(logFile, LOGLEVEL)
	readfile.ReadLine2Map(ROBOTLISTPATH, handleloadrobot)
	allRobotnum = len(mapRobotFree)
}

// Select a robot for a singer from the system.
func addRobot2Room(singerId int32) (robotid int32) {
	if 0 != len(mapRobotFree) {
		for k, _ := range mapRobotFree {
			robotid = k
			(*(mapOnlineSinger[singerId]))[robotid] = time.Now().Unix()
			delete(mapRobotFree, robotid)
		}
	} else {
		robotid = 0
		loglj.Errorln(logFile, "The robot program has no robot !")
	}
	loglj.Debugf(logFile, "mapRobotFree num is %d ; mapOnlineSinger num is %d ", len(mapRobotFree), len(mapOnlineSinger))
	loglj.Infof(logFile, "Robot %d get into the singer %d room", robotid, singerId)
	return robotid
}

// After the singer began to live broadcast, we need to create a robot map for this singer.
func createNewSingerMap(singerId int32) (singerMapp *(map[int32]int64)) {
	mapSingerRoomRobot = make(map[int32]int64, MAXROOMROBOTNUM)
	mapOnlineSinger[singerId] = &mapSingerRoomRobot
	singerMapp = &mapSingerRoomRobot
	loglj.Infof(logFile, "Singer %d createNewSingerMap . This system load singer host number %d ", singerId, len(mapOnlineSinger))
	return singerMapp
}

// The handle of load robot
func handleloadrobot(line string) {
	var robotid int
	typechange.String2Int(&line, &robotid)
	mapRobotFree[int32(robotid)] = true
}

// The host began to live broadcast
func SingerOpenBroadcastRQ(singerId int32) {
	mapSingerRoomRobot = make(map[int32]int64, MAXROOMROBOTNUM) //    k:robotid    v:unix time(int64)
	mapOnlineSinger[singerId] = &mapSingerRoomRobot
}

func SingerCloseBroadcast(singerId int32) {
	mapSingerRoomRobot = make(map[int32]int64) //clean the singerid map
	delete(mapOnlineSinger, singerId)
}

func SingerAddrobot2Rom(singerId int32, robotNum int32) {

}

func main() {
	mapt := make(map[int]int, 200)
	fmt.Println("lenmap", len(mapt))
	//readfile.ReadLine("D:/GOPATH/resources/ljnew/memory/robottt.list", Printline)
	readfile.ReadLine2Map("D:/GOPATH/resources/ljnew/memory/robottt.list", handleloadrobot)

	for k, v := range mapt {
		fmt.Println("k", k, "v", v)
	}
	fmt.Println("lenmap", len(mapt))
	singer1 := 10005130
	robot1_0 := OutPutData8Map(singer1, &mapt)
	fmt.Println("singer1:", singer1, "robot1_0:", robot1_0)

	loglj.SetLoglevel(loglj.L_DEBUG)
	loglj.SetLogFilePatch("D:/GOPATH/resources/ljnew/loglj/" + time.Now().Format("20060102") + ".log")
	loglj.Linfo("singer1:", singer1, "robot1_0:", robot1_0)
	loglj.Ldebug("singer1:", singer1, "robot1_0:", robot1_0)
	loglj.Lerror("singer1:", singer1, "robot1_0:", robot1_0)
	loglj.Closefile()

	for k, v := range mapt {
		fmt.Println("k", k, "v", v)
	}
	fmt.Println("lenmap", len(mapt))

}

// Providing external H5Q I/O interface for query
func GetRobotAllNum() (num int) {
	return allRobotnum
}
func GetOnlineSingerNum() (num int) {
	return len(mapOnlineSinger)
}
func GetSingerRoomRobotNum(singerId int32) (num int) {
	return len(mapSingerRoomRobot)
}
func GetFreeRobotNum() (num int) {
	return len(mapRobotFree)
}
func ReStartSystem() {
	mapRobotFree = make(map[int32]bool, MAXROBOTNUM)
	mapOnlineSinger = make(map[int32](*(map[int32]int64)), MAXSINGERNUM)
	logFile, _ = loglj.LogFileOpen(LOGPATH + time.Now().Format("20060102") + ".log")
	loglj.SetLoglevel(logFile, LOGLEVEL)
	readfile.ReadLine2Map(ROBOTLISTPATH, handleloadrobot)
	allRobotnum = len(mapRobotFree)
}
