package datagroove

import (
	"unsafe"
	"encoding/binary"
	"errors"
	"gotest/common/logfile"
)

const GROOVE_CAP = 1024

type DataBuff struct {
	SGroove       []byte
	LenGroove     uint16
	LenRemove     uint16
	LenData       uint16
	shortFlag     uint8
}
// I/O init buffer
func (d *DataBuff)BufferInit() () {
	d.LenRemove = 0
	d.LenData = 0
	d.LenGroove = GROOVE_CAP
	d.SGroove = make([]byte,  d.LenGroove )
}
//I/O get buffer address 只用于调试
func (d *DataBuff)GetBufferAddress() unsafe.Pointer {
	return unsafe.Pointer(&(d.SGroove[0]))
}
//I/O get data address 只用于调试
func (d *DataBuff)GetDataAddress(pos uint16) unsafe.Pointer {
	return unsafe.Pointer(&(d.SGroove[d.LenRemove + pos]))
}
//获取数据槽内数据
func (d *DataBuff)PrintDataGroove(){
	logfile.GlobalLog.Infoln("PrintDataGroove::SGroove:" , d.SGroove)
}
//跳过数据槽内数据
func (d *DataBuff)DataJump(length uint16) (error){
	if d.LenData >= length {
		d.LenRemove += length
		d.LenData -= length
		return nil
	}else {
		return	errors.New("ERROR : Data slot does not have enough data")
	}

}

//本函数只用于变更基本数据槽规格和后续扩容方式
//用于调整数据槽基本槽大小（默认情况最好是1024等常用计算机存储单位）
//用于调整数据槽基本页（数据槽增加是以基本槽1页为单位，但是默认数据槽的页可以  不是 1页）
//因此实际数据槽大小是基本页 和 基本槽大小 的 乘积，增页是按照基本页增加，不是槽大小倍数增加
func (d *DataBuff)SetDataGroove(page uint16 )  {
	d.crossBufferMove(page)
}

//I/O Add a cup of data. 向后开辟一杯数据
func (d *DataBuff)AddDataACup(maxCup uint16) {
	if 0 == d.LenData {
		d.LenRemove = 0
	}
	if GROOVE_CAP != d.LenGroove  {
		if d.LenData + maxCup <= GROOVE_CAP{
			d.shortFlag += 1
		}else {
			d.shortFlag = 0
		}
		if d.shortFlag >= 5{
			d.crossBufferMove(GROOVE_CAP)
		}
	}
	if d.LenGroove-d.LenData-d.LenRemove < maxCup {
		if d.LenGroove-d.LenData > maxCup{
			d.bufferMove()
		}else {
			page := (d.LenData + maxCup)/(GROOVE_CAP) + 1
			d.crossBufferMove(page)
		}
	}
}
//////////////////////////////////////////////////////////////////////////////////////////
// package function move buffer
func (d *DataBuff)bufferMove()  {
	if 0 != d.LenRemove {
		copy(d.SGroove[0:d.LenData], d.SGroove[d.LenRemove:(d.LenRemove + d.LenData)])
		d.LenRemove = 0
	}
}
// package function move buffer
func (d *DataBuff)crossBufferMove(page uint16)  {
	groove := make([]byte, GROOVE_CAP * page)
	if 0 != d.LenData {
		copy(groove[0:d.LenData], d.SGroove[d.LenRemove:(d.LenRemove + d.LenData)])
	}
	d.LenGroove = GROOVE_CAP * page
	d.LenRemove = 0
	d.shortFlag = 0
	d.SGroove = groove
}

////////////////////////////////////////////////////////////////////////////////////////
//Data slot processing
////////////////////////////////////////////////////////////////////////////////////////
//DataSlotWriteUint8
func (d *DataBuff)DataSlotWriteByte(position uint16 , wData byte)  {
	d.SGroove[position] = wData
}
//DataSlotReadUint8
func (d *DataBuff)DataSlotReadByte(position uint16 ) ( byte) {
	return d.SGroove[position]
}//DataSlotWriteUint8
func (d *DataBuff)DataSlotWriteUint8(position uint16 , wData uint8)  {
	d.SGroove[position] = wData
}
//DataSlotReadUint8
func (d *DataBuff)DataSlotReadUint8(position uint16 ) ( uint8) {
	return d.SGroove[position]
}
//DataSlotWriteInt8
func (d *DataBuff)DataSlotWriteInt8(position uint16 , wData int8) {
	if wData < 0 {
		d.SGroove[position] = byte(2<<(8-1) + uint16(wData))
	} else {
		d.SGroove[position] = byte(wData)
	}
}
//DataSlotReadInt8
func (d *DataBuff)DataSlotReadInt8(position uint16 ) ( int8 ) {
	return int8(d.SGroove[position])
}
//DataSlotWriteUint16
func (d *DataBuff)DataSlotWriteUint16(position uint16 , wData uint16)  {
	binary.LittleEndian.PutUint16(d.SGroove[position:], wData)
}
//DataSlotReadUint16
func (d *DataBuff)DataSlotReadUint16(position uint16 ) ( uint16 ) {
	return binary.LittleEndian.Uint16(d.SGroove[position:])
}
//DataSlotWriteUint32
func (d *DataBuff)DataSlotWriteUint32(position uint16 , wData uint32)  {
	binary.LittleEndian.PutUint32(d.SGroove[position:], wData)
}
//DataSlotReadUint32
func (d *DataBuff)DataSlotReadUint32(position uint16 ) ( uint32 ) {
	return binary.LittleEndian.Uint32(d.SGroove[position:])
}
//DataSlotWriteUint64
func (d *DataBuff)DataSlotWriteUint64(position uint16 , wData uint64)  {
	binary.LittleEndian.PutUint64(d.SGroove[position:], wData)
}
//DataSlotReadUint64
func (d *DataBuff)DataSlotReadUint64(position uint16 ) ( uint64 ) {
	return binary.LittleEndian.Uint64(d.SGroove[position:])
}
//DataSlotWriteString16
func (d *DataBuff)DataSlotWriteString16(position uint16 , s *string)()  {
	lenDat :=uint16(len(*s))
	d.DataSlotWriteUint16(position , lenDat)
	copy(d.SGroove[position+2:position+2+uint16(lenDat)],*s)
}
//DataSlotReadString16
func (d *DataBuff)DataSlotReadString16(position uint16 ,s *string)()  {
	length := d.DataSlotReadUint16(position)
	if 0!=length {
		*s = string(d.SGroove[position+2:position+2+uint16(length)])

	}else {
		*s =  ""
	}
}
//DataSlotWriteString32
func (d *DataBuff)DataSlotWriteString32(position uint16 , s *string)()  {
	lenDat :=uint32(len(*s))
	d.DataSlotWriteUint32(position , lenDat)
	copy(d.SGroove[position+4:position+4+uint16(lenDat)],*s)
}
//DataSlotReadString32
func (d *DataBuff)DataSlotReadString32(position uint16 ,s *string)()  {
	length := d.DataSlotReadUint32(position)
	if 0!=length {
		*s = string(d.SGroove[position+4:position+4+uint16(length)])
	}else {
		*s =  ""
	}
}