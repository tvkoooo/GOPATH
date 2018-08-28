package robotvector

import (
	"os"
	"bufio"
	"strings"
	"io"
	"xcbbrobot/common/typechange"
	"fmt"
	//"sync"

)

type MapAppRobot struct {
	mapRobotBool map[uint32]bool
	//lockMapAppRobot *sync.RWMutex
}


func (p *MapAppRobot)RobotFreeInit(num int){
	p.mapRobotBool = make(map[uint32]bool , num)
	//p.lockMapAppRobot = new(sync.RWMutex)
}

func (p *MapAppRobot)PrintRobotMap()(){
	//p.lockMapAppRobot.RLock()
	for k, v := range p.mapRobotBool{
		fmt.Println("k:", k, "   v:", v)
	}
	//p.lockMapAppRobot.RUnlock()
}

func (p *MapAppRobot)Len()(int ){
	return  len(p.mapRobotBool)
}

func (p *MapAppRobot)handleLoadRobot(line string) {
	var robotId int
	typechange.String2Int(&line, &robotId)
	if 0!=robotId {
		//p.lockMapAppRobot.Lock()
		p.mapRobotBool [uint32(robotId)] = true
		//p.lockMapAppRobot.Unlock()
	}
}
func (p *MapAppRobot)readLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func (p *MapAppRobot)LoadRobot(robotList string )  {
	p.readLine(robotList ,p.handleLoadRobot)
}

func (p *MapAppRobot)AddRobot(robotId uint32)  {
	//p.lockMapAppRobot.Lock()
	(p.mapRobotBool)[uint32(robotId)] = true
	//p.lockMapAppRobot.Unlock()
}

func (p *MapAppRobot)DelRobot(robotId uint32)  {
	//p.lockMapAppRobot.Lock()
	delete(p.mapRobotBool, robotId)
	//p.lockMapAppRobot.Unlock()
}
func (p *MapAppRobot)CleanRobot()  {
	//p.lockMapAppRobot.Lock()
	p.mapRobotBool = make(map[uint32]bool)
	//p.lockMapAppRobot.Unlock()
}
func (p *MapAppRobot)PopRobot()(robotId uint32)  {
	robotId = 0
	if 0!=len(p.mapRobotBool) {
		for k, _ := range p.mapRobotBool {
			robotId = k
			p.DelRobot(k)
			break
		}
	}
	return robotId
}
