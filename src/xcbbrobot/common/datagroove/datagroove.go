package datagroove

import (
	"unsafe"
	"xcbbrobot/common/datastream"
	"encoding/binary"
)

const GROOVE_CAP = 10

type DataBuff struct {
	sGroove       []byte
	lenGroove     int
	lenRemove     int
	lenData       int
	shortFlag     int
}
// I/O init buffer
func (d *DataBuff)BufferInit() () {
	d.sGroove = make([]byte,  GROOVE_CAP )
	d.lenGroove = GROOVE_CAP
	d.lenRemove = 0
	d.lenData = 0
	d.shortFlag = 0

}
//I/O get buffer address 只用于调试
func (d *DataBuff)GetBufferAddress() unsafe.Pointer {
	return unsafe.Pointer(&(d.sGroove[0]))
}
//I/O get data address 只用于调试
func (d *DataBuff)GetDataAddress(pos int) unsafe.Pointer {
	return unsafe.Pointer(&(d.sGroove[d.lenRemove + pos]))
}

//I/O Add a cup of data. 向后开辟一杯数据
func (d *DataBuff)AddDataACup(maxCup int) {
	if GROOVE_CAP != d.lenGroove  {
		if d.lenData + maxCup < GROOVE_CAP{
			d.shortFlag += 1
		}else {
			d.shortFlag = 0
		}
		if d.shortFlag >= 5{
			d.crossBufferMove(1)
		}
	}
	if d.lenGroove-d.lenRemove-d.lenData < maxCup {
		if d.lenGroove-d.lenData > maxCup{
			d.bufferMove()
		}else {
			page := (d.lenData + maxCup)/GROOVE_CAP + 1
			d.crossBufferMove(page)
		}
	}
}

//I/O get data address
func (d *DataBuff)DataAppend(addData []byte)() {
	length := len(addData)
	if GROOVE_CAP != d.lenGroove  {
		if d.lenData + length < GROOVE_CAP{
			d.shortFlag += 1
		}else {
			d.shortFlag = 0
		}
		if d.shortFlag >= 5{
			d.crossBufferMove(1)
			d.bufferAppend(addData)
			return
		}
	}
	if (d.lenGroove - d.lenData - length ) >=  0{
		if (d.lenGroove - d.lenRemove - d.lenData ) >=  length {
			d.bufferAppend(addData)
		}else{
			d.bufferMove()
			d.bufferAppend(addData)
		}
	}else{
		page := (d.lenData + length)/GROOVE_CAP + 1
		d.crossBufferMove(page)
		d.bufferAppend(addData)
	}
}
//I/O get data address
func (d *DataBuff)DataPop(length int)(popData []byte) {
	popData = d.bufferPop(length)
	return
}
//I/O get data address
func (d *DataBuff)MessagePop()(popData []byte) {
	if d.lenData < 4 {
		popData = nil
	}
	length, _ := datastream.GetUint32(d.sGroove[d.lenRemove:])
	popData = d.bufferPop(int(length))
	return
}

//////////////////////////////////////////////////////////////////////////////////////////
// package function move buffer
func (d *DataBuff)bufferMove()  {
	copy(d.sGroove[0:d.lenData], d.sGroove[d.lenRemove:(d.lenRemove + d.lenData)])
	d.lenRemove = 0
}
// package function move buffer
func (d *DataBuff)crossBufferMove(page int)  {
	groove := make([]byte, GROOVE_CAP * page)
	copy(groove[0:d.lenData], d.sGroove[d.lenRemove:(d.lenRemove + d.lenData)])
	d.lenGroove = GROOVE_CAP * page
	d.lenRemove = 0
	d.shortFlag = 0
	d.sGroove = groove
}
//package function buffer append slices
func (d *DataBuff)bufferAppend(addData []byte)  {
	copy(d.sGroove[(d.lenRemove + d.lenData):(d.lenRemove + d.lenData)+ len(addData)], addData)
	d.lenData += len(addData)
}
//package function buffer pop slices
func (d *DataBuff)bufferPop(length int)(popData []byte)  {
	if length<= d.lenData {
		popData = d.sGroove[d.lenRemove:(d.lenRemove + length)]
		d.lenRemove += length
		d.lenData -= length
	}else {
		popData = nil
	}
	if 0 == d.lenData {
		d.lenRemove = 0
	}
	return
}
////////////////////////////////////////////////////////////////////////////////////////
//Data slot processing
////////////////////////////////////////////////////////////////////////////////////////
//DataSlotWriteUint8
func (d *DataBuff)DataSlotWriteByte(position int , wData byte)  {
	d.sGroove[position] = wData
}
//DataSlotReadUint8
func (d *DataBuff)DataSlotReadByte(position int ) ( byte) {
	return d.sGroove[position]
}//DataSlotWriteUint8
func (d *DataBuff)DataSlotWriteUint8(position int , wData uint8)  {
	d.sGroove[position] = wData
}
//DataSlotReadUint8
func (d *DataBuff)DataSlotReadUint8(position int ) ( uint8) {
	return d.sGroove[position]
}
//DataSlotWriteInt8
func (d *DataBuff)DataSlotWriteInt8(position int , wData int8) {
	if wData < 0 {
		d.sGroove[position] = byte(2<<(8-1) + int(wData))
	} else {
		d.sGroove[position] = byte(wData)
	}
}
//DataSlotReadInt8
func (d *DataBuff)DataSlotReadInt8(position int ) ( int8 ) {
	return int8(d.sGroove[position])
}
//DataSlotWriteUint16
func (d *DataBuff)DataSlotWriteUint16(position int , wData uint16)  {
	binary.LittleEndian.PutUint16(d.sGroove[position:], wData)
}
//DataSlotReadUint16
func (d *DataBuff)DataSlotReadUint16(position int ) ( uint16 ) {
	return binary.LittleEndian.Uint16(d.sGroove[position:])
}
//DataSlotWriteUint32
func (d *DataBuff)DataSlotWriteUint32(position int , wData uint32)  {
	binary.LittleEndian.PutUint32(d.sGroove[position:], wData)
}
//DataSlotReadUint32
func (d *DataBuff)DataSlotReadUint32(position int ) ( uint32 ) {
	return binary.LittleEndian.Uint32(d.sGroove[position:])
}
//DataSlotWriteUint64
func (d *DataBuff)DataSlotWriteUint64(position int , wData uint64)  {
	binary.LittleEndian.PutUint64(d.sGroove[position:], wData)
}
//DataSlotReadUint64
func (d *DataBuff)DataSlotReadUint64(position int ) ( uint64 ) {
	return binary.LittleEndian.Uint64(d.sGroove[position:])
}
//DataSlotWriteUint64
func (d *DataBuff)DataSlotWriteString16(position int , wData string)(int)  {
	lenDat :=uint16(len(wData))
	d.DataSlotWriteUint16(position , lenDat)
	copy(d.sGroove[position+2:],wData)
 	return int(2+lenDat)
}