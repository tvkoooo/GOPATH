package robotvector

import (
	"os"
	"bufio"
	"strings"
	"io"
	"robot_d/common/typechange"
	"sync"
	"robot_d/common/logfile"
)

type MapAppRobot struct {
	m map[uint32]bool
	l *sync.RWMutex
}


func (p *MapAppRobot)RobotFreeInit(){
	p.m = make(map[uint32]bool)
	p.l = new(sync.RWMutex)
}

func (p *MapAppRobot)PrintRobotMap()(){
	logfile.GlobalLog.Infoln("PrintRobotMap::MapAppRobot:",p.m)
}

func (p *MapAppRobot)Len()(int ){
	return  len(p.m)
}

func (p *MapAppRobot)handleLoadRobot(line string) {
	var robotId int
	typechange.String2Int(&line, &robotId)
	if 0!=robotId {
		p.l.Lock()
		p.m [uint32(robotId)] = true
		p.l.Unlock()
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
	p.l.Lock()
	(p.m)[uint32(robotId)] = true
	p.l.Unlock()
}

func (p *MapAppRobot)DelRobot(robotId uint32)  {
	p.l.Lock()
	delete(p.m, robotId)
	p.l.Unlock()
}
func (p *MapAppRobot)CleanRobot()  {
	p.l.Lock()
	p.m = make(map[uint32]bool)
	p.l.Unlock()
}
func (p *MapAppRobot)PopRobot()(robotId uint32)  {
	robotId = 0
	if 0!=len(p.m) {
		for k, _ := range p.m {
			robotId = k
			p.DelRobot(k)
			break
		}
	}
	return robotId
}
