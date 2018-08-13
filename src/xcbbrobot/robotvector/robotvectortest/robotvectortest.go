package main

import (
	"fmt"
	"xcbbrobot/robotvector"
	"unsafe"
)

func main() {
	var ar robotvector.AppRobot
	ar.RobotFreeInit(100)

	ar.LoadRobot("./xcbbrobot/robotvector/robottt.list")
	ar.PrintRobotMap()

	ar.AddRobot(10005259)
	fmt.Println("add map")
	ar.PrintRobotMap()

	ar.DelRobot(10002846)
	fmt.Println("sub map now")
	ar.PrintRobotMap()



	ar.CleanRobot()
	fmt.Println("new map clean")
	ar.PrintRobotMap()


	ar.AddRobot(10002844)
	fmt.Println("new add map")
	ar.PrintRobotMap()

	fmt.Println("load new map")
	ar.LoadRobot("./xcbbrobot/robotvector/robottt.list")
	ar.PrintRobotMap()

	a:= ar.PopRobot()
	fmt.Println("a", a, "len", ar.Len())
	b:= ar.PopRobot()
	fmt.Println("a", b, "len", ar.Len())
	qq:= ar.PopRobot()
	fmt.Println("a", qq, "len", ar.Len())
	qb:= ar.PopRobot()
	fmt.Println("a", qb, "len", ar.Len())
	wb:= ar.PopRobot()
	fmt.Println("a", wb, "len", ar.Len())
	be:= ar.PopRobot()
	fmt.Println("a", be, "len", ar.Len())
	gb:= ar.PopRobot()
	fmt.Println("a", gb, "len", ar.Len())



	ccc := make(map[int]int )
	fmt.Println(ccc)
	fmt.Println(len(ccc))
	ccc[1] = 1
	ccc[2] = 2
	fmt.Println(ccc)
	fmt.Println(len(ccc))
	fmt.Println(unsafe.Pointer(&(ccc)))
	ccc[3] = 3
	ccc[4] = 4
	fmt.Println(ccc)
	fmt.Println(len(ccc))
	fmt.Println(unsafe.Pointer(&ccc))




	ddd := make(map[int]int , 2)
	fmt.Println(ddd)
	fmt.Println(len(ddd))
	ddd[1] = 1
	ddd[2] = 2
	fmt.Println(ddd)
	fmt.Println(len(ddd))
	fmt.Println(unsafe.Pointer(&ddd))
	ddd[3] = 3
	ddd[4] = 4
	fmt.Println(ddd)
	fmt.Println(len(ddd))
	fmt.Println(unsafe.Pointer(&ddd))


}
