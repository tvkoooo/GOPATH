package robotvector

import (
	"os"
	"bufio"
	"strings"
	"io"
	"xcbbrobot/common/typechange"
	"fmt"
)

type AppRobot struct {
	mapRobotBool map[uint32]bool
}


func (p *AppRobot)RobotFreeInit(num int){
	p.mapRobotBool = make(map[uint32]bool , num)
}

func (p *AppRobot)PrintRobotMap()(){
	for k, v := range p.mapRobotBool{
		fmt.Println("k:", k, "   v:", v)
	}
}

func (p *AppRobot)Len()(int ){
	return  len(p.mapRobotBool)
}

func handleLoadRobot(line string ,m *map[uint32]bool) {
	var robotId int
	typechange.String2Int(&line, &robotId)
	if 0!=robotId {
		(*m) [uint32(robotId)] = true
	}
}
func (p *AppRobot)readLine(fileName string, handler func(string,*map[uint32]bool)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line ,&(p.mapRobotBool))
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func (p *AppRobot)LoadRobot(robotList string )  {
	p.readLine(robotList ,handleLoadRobot)
}

func (p *AppRobot)AddRobot(robotId uint32)  {
	(p.mapRobotBool)[uint32(robotId)] = true
}

func (p *AppRobot)DelRobot(robotId uint32)  {
	delete(p.mapRobotBool, robotId)
}
func (p *AppRobot)CleanRobot()  {
	p.mapRobotBool = make(map[uint32]bool)
}
func (p *AppRobot)PopRobot()(robotId uint32)  {
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
