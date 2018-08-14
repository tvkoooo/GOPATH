package main

import (
	"xcbbrobot/common/datagroove"
	"fmt"
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
	fmt.Println("in88",in88,"in99",in99)

}