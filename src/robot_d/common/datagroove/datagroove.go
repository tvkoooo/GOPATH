package datagroove

import (
	"unsafe"
	"encoding/binary"
	"time"
	"errors"
	"robot_d/common/logfile"
)

const GROOVE_CAP = 1024

type DataBuff struct {
	SGroove       []byte
	LenGroove     int
	LenRemove     int
	LenData       int
	basePage	  int
	baseGroove    int
	numCross      int
	maxGroove	  int
	initTime      time.Time
	shortFlag     int
}
// I/O init buffer
func (d *DataBuff)BufferInit() () {
	d.LenRemove = 0
	d.LenData = 0
	d.basePage = 1
	d.baseGroove = GROOVE_CAP
	d.numCross = 0
	d.initTime= time.Now()
	d.shortFlag = 0
	d.LenGroove = d.baseGroove * d.basePage
	d.maxGroove = d.LenGroove
	d.SGroove = make([]byte,  d.LenGroove )
}
//I/O get buffer address 只用于调试
func (d *DataBuff)GetBufferAddress() unsafe.Pointer {
	return unsafe.Pointer(&(d.SGroove[0]))
}
//I/O get data address 只用于调试
func (d *DataBuff)GetDataAddress(pos int) unsafe.Pointer {
	return unsafe.Pointer(&(d.SGroove[d.LenRemove + pos]))
}
//获取数据槽内数据
func (d *DataBuff)PrintDataGroove(){
	logfile.GlobalLog.Infoln("PrintDataGroove::SGroove:" , d.SGroove)
}
//跳过数据槽内数据
func (d *DataBuff)DataJump(length int) (error){
	if d.LenData >= length {
		d.LenRemove += length
		d.LenData -= length
		return nil
	}else {
		return	errors.New("ERROR : Data slot does not have enough data")
	}

}

//获取数据槽工作情况,1天会出现多少次数据槽回归
func (d *DataBuff)PrintStatus() {

	interval := int(time.Now().Sub(d.initTime)/1E9 +1)
	oneDay := int(24*60*60)
	NpD := oneDay/interval * d.numCross
	logfile.GlobalLog.Infoln("PrintStatus::basePage:" , d.basePage," baseGroove:" , d.baseGroove," numCross:" , d.numCross," maxGroove:" , d.maxGroove," shortFlag:" , d.shortFlag,"interval:" , interval," NpD:" , NpD)
}
//本函数只用于变更基本数据槽规格和后续扩容方式
//用于调整数据槽基本槽大小（默认情况最好是1024等常用计算机存储单位）
//用于调整数据槽基本页（数据槽增加是以基本槽1页为单位，但是默认数据槽的页可以  不是 1页）
//因此实际数据槽大小是基本页 和 基本槽大小 的 乘积，增页是按照基本页增加，不是槽大小倍数增加
func (d *DataBuff)SetDataGroove(page int , baseGroove int)  {
	d.baseGroove = baseGroove
	d.basePage = page
	d.crossBufferMove(page)
}

//I/O Add a cup of data. 向后开辟一杯数据
func (d *DataBuff)AddDataACup(maxCup int) {
	if 0 == d.LenData {
		d.LenRemove = 0
	}
	if d.baseGroove * d.basePage != d.LenGroove  {
		if d.LenData + maxCup <= d.baseGroove * d.basePage{
			d.shortFlag += 1
		}else {
			d.shortFlag = 0
		}
		if d.shortFlag >= 5{
			d.crossBufferMove(d.basePage)
		}
	}
	if d.LenGroove-d.LenData-d.LenRemove < maxCup {
		if d.LenGroove-d.LenData > maxCup{
			d.bufferMove()
		}else {
			page := (d.LenData + maxCup)/(d.baseGroove) + 1
			d.crossBufferMove(page)
		}
	}
}

//I/O get data address
func (d *DataBuff)DataAppend(addData []byte)() {
	length := len(addData)
	if d.baseGroove * d.basePage != d.LenGroove  {
		if d.LenData + length <= d.baseGroove * d.basePage{
			d.shortFlag += 1
		}else {
			d.shortFlag = 0
		}
		if d.shortFlag >= 5{
			d.crossBufferMove(d.basePage)
			d.bufferAppend(addData)
			return
		}
	}
	if (d.LenGroove - d.LenData - length ) >=  0{
		if (d.LenGroove - d.LenRemove - d.LenData ) >=  length {
			d.bufferAppend(addData)
		}else{
			d.bufferMove()
			d.bufferAppend(addData)
		}
	}else{
		page := (d.LenData + length)/(d.baseGroove * d.basePage) + 1
		d.crossBufferMove(page)
		d.bufferAppend(addData)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////
// package function move buffer
func (d *DataBuff)bufferMove()  {
	copy(d.SGroove[0:d.LenData], d.SGroove[d.LenRemove:(d.LenRemove + d.LenData)])
	d.LenRemove = 0
}
// package function move buffer
func (d *DataBuff)crossBufferMove(page int)  {
	groove := make([]byte, d.baseGroove * page)
	copy(groove[0:d.LenData], d.SGroove[d.LenRemove:(d.LenRemove + d.LenData)])
	d.LenGroove = d.baseGroove * page
	d.LenRemove = 0
	d.shortFlag = 0
	d.SGroove = groove
	d.numCross ++
	if d.LenGroove > d.maxGroove {
		d.maxGroove = d.LenGroove
	}
}
//package function buffer append slices
func (d *DataBuff)bufferAppend(addData []byte)  {
	copy(d.SGroove[(d.LenRemove + d.LenData):(d.LenRemove + d.LenData)+ len(addData)], addData)
	d.LenData += len(addData)
}
//package function buffer pop slices
func (d *DataBuff)bufferPop(length int)(popData []byte)  {
	if length<= d.LenData {
		popData = d.SGroove[d.LenRemove:(d.LenRemove + length)]
		d.LenRemove += length
		d.LenData -= length
	}else {
		popData = nil
	}
	if 0 == d.LenData {
		d.LenRemove = 0
	}
	return
}
////////////////////////////////////////////////////////////////////////////////////////
//Data slot processing
////////////////////////////////////////////////////////////////////////////////////////
//DataSlotWriteUint8
func (d *DataBuff)DataSlotWriteByte(position int , wData byte)  {
	d.SGroove[position] = wData
}
//DataSlotReadUint8
func (d *DataBuff)DataSlotReadByte(position int ) ( byte) {
	return d.SGroove[position]
}//DataSlotWriteUint8
func (d *DataBuff)DataSlotWriteUint8(position int , wData uint8)  {
	d.SGroove[position] = wData
}
//DataSlotReadUint8
func (d *DataBuff)DataSlotReadUint8(position int ) ( uint8) {
	return d.SGroove[position]
}
//DataSlotWriteInt8
func (d *DataBuff)DataSlotWriteInt8(position int , wData int8) {
	if wData < 0 {
		d.SGroove[position] = byte(2<<(8-1) + int(wData))
	} else {
		d.SGroove[position] = byte(wData)
	}
}
//DataSlotReadInt8
func (d *DataBuff)DataSlotReadInt8(position int ) ( int8 ) {
	return int8(d.SGroove[position])
}
//DataSlotWriteUint16
func (d *DataBuff)DataSlotWriteUint16(position int , wData uint16)  {
	binary.LittleEndian.PutUint16(d.SGroove[position:], wData)
}
//DataSlotReadUint16
func (d *DataBuff)DataSlotReadUint16(position int ) ( uint16 ) {
	return binary.LittleEndian.Uint16(d.SGroove[position:])
}
//DataSlotWriteUint32
func (d *DataBuff)DataSlotWriteUint32(position int , wData uint32)  {
	binary.LittleEndian.PutUint32(d.SGroove[position:], wData)
}
//DataSlotReadUint32
func (d *DataBuff)DataSlotReadUint32(position int ) ( uint32 ) {
	return binary.LittleEndian.Uint32(d.SGroove[position:])
}
//DataSlotWriteUint64
func (d *DataBuff)DataSlotWriteUint64(position int , wData uint64)  {
	binary.LittleEndian.PutUint64(d.SGroove[position:], wData)
}
//DataSlotReadUint64
func (d *DataBuff)DataSlotReadUint64(position int ) ( uint64 ) {
	return binary.LittleEndian.Uint64(d.SGroove[position:])
}
//DataSlotWriteString16
func (d *DataBuff)DataSlotWriteString16(position int , s *string)()  {
	lenDat :=uint16(len(*s))
	d.DataSlotWriteUint16(position , lenDat)
	copy(d.SGroove[position+2:position+2+int(lenDat)],*s)
}
//DataSlotReadString16
func (d *DataBuff)DataSlotReadString16(position int ,s *string)()  {
	length := d.DataSlotReadUint16(position)
	if 0!=length {
		*s = string(d.SGroove[position+2:position+2+int(length)])

	}else {
		*s =  ""
	}
}
//DataSlotWriteString32
func (d *DataBuff)DataSlotWriteString32(position int , s *string)()  {
	lenDat :=uint32(len(*s))
	d.DataSlotWriteUint32(position , lenDat)
	copy(d.SGroove[position+4:position+4+int(lenDat)],*s)
}
//DataSlotReadString32
func (d *DataBuff)DataSlotReadString32(position int ,s *string)()  {
	length := d.DataSlotReadUint32(position)
	if 0!=length {
		*s = string(d.SGroove[position+4:position+4+int(length)])
	}else {
		*s =  ""
	}
}