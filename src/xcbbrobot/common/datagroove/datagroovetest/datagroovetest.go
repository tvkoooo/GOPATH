package main

import (
	"xcbbrobot/common/datagroove"
	"fmt"
	"xcbbrobot/common/message"
)
//测试需要把槽的容量调成10
func main()  {
	var dat datagroove.DataBuff
	dat.BufferInit()
	fmt.Println("dat:",dat )
	fmt.Println("dat_address :",dat.GetBufferAddress() )
	fmt.Println("dat_address :",dat.GetDataAddress(2) )

	adddat := []byte{1,2,3,4,5,6}
	dat.DataAppend(adddat)
	fmt.Println("dat:",dat )
	fmt.Println("dat_address :",dat.GetBufferAddress() )
	fmt.Println("dat_address :",dat.GetDataAddress(2) )

	popd := dat.DataPop(2)
	fmt.Println("dat:",dat )
	fmt.Println("popd:",popd )
	fmt.Println("dat_address :",dat.GetBufferAddress() )
	fmt.Println("dat_address :",dat.GetDataAddress(0) )

	//popd1 := dat.DataPop(2)
	//fmt.Println("dat:",dat )
	//fmt.Println("popd1:",popd1 )
	//fmt.Println("dat_address :",dat.GetBufferAddress() )
	//fmt.Println("dat_address :",dat.GetDataAddress(0) )

	adddat1 := []byte{7,8,9,10,11,12,13}
	dat.DataAppend(adddat1)
	fmt.Println("dat:",dat )
	fmt.Println("dat_address :",dat.GetBufferAddress() )
	fmt.Println("dat_address :",dat.GetDataAddress(0) )

	dat.DataPop(2)
	fmt.Println("dat:",dat )
	dat.DataPop(2)
	fmt.Println("dat:",dat )
	dat.DataPop(4)
	fmt.Println("dat:",dat )

	adddat2 := []byte{'a'}
	dat.DataAppend(adddat2)
	fmt.Println("dat:",dat )
	dat.DataAppend(adddat2)
	fmt.Println("dat:",dat )
	dat.DataAppend(adddat2)
	fmt.Println("dat:",dat )
	dat.DataAppend(adddat2)
	fmt.Println("dat:",dat )
	dat.DataAppend(adddat2)
	fmt.Println("dat:",dat )
	dat.DataAppend(adddat2)
	fmt.Println("dat:",dat )


	var in88,in99 int8
	//var uin88,uin99 uint8
	in88 = -3

	var slot1 datagroove.DataBuff
	slot1.BufferInit()
	slot1.DataSlotWriteInt8(0,in88)
	in99 = slot1.DataSlotReadInt8(0)
	fmt.Println("in88",in88,"in99",in99,slot1)

	in88 = -66
	slot1.DataSlotWriteInt8(4,in88)
	in99 = slot1.DataSlotReadInt8(4)
	fmt.Println("in88",in88,"in99",in99,slot1)

	var ph message.Packhead
	ph.Uri = (101 << 8) | 23
	ph.Sid = 0
	ph.Rescode = 200
	ph.Tag = 1

	var robotcon message.PRegisteredPI
	robotcon.Id = 0
	robotcon.PIType = 64
	robotcon.PIPass = "nimei"
	var lengda uint32
	lengda = uint32(13 + 4+4+2+len(robotcon.PIPass))
	slot1.AddDataACup(int(lengda))
	slot1.DataSlotWriteUint32(4,ph.Uri)
	slot1.DataSlotWriteUint16(6,ph.Sid)
	slot1.DataSlotWriteUint16(8,ph.Rescode)
	slot1.DataSlotWriteUint8(10,ph.Tag)
	fmt.Println("ph",slot1)


	slot1.DataSlotWriteUint32(11,robotcon.Id)
	slot1.DataSlotWriteUint32(15,robotcon.PIType)
	slot1.DataSlotWriteString16(19,robotcon.PIPass)

	slot1.DataSlotWriteUint32(0,lengda)
	slot1.LenData = int(lengda)
	fmt.Println("all",slot1)
}