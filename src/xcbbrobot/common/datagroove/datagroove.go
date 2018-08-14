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
	LenRemove     int
	LenData       int
	shortFlag     int
}
// I/O init buffer
func (d *DataBuff)BufferInit() () {
	d.sGroove = make([]byte,  GROOVE_CAP )
	d.lenGroove = GROOVE_CAP
	d.LenRemove = 0
	d.LenData = 0
	d.shortFlag = 0

}
//I/O get buffer address 只用于调试
func (d *DataBuff)GetBufferAddress() unsafe.Pointer {
	return unsafe.Pointer(&(d.sGroove[0]))
}
//I/O get data address 只用于调试
func (d *DataBuff)GetDataAddress(pos int) unsafe.Pointer {
	return unsafe.Pointer(&(d.sGroove[d.LenRemove + pos]))
}

//I/O Add a cup of data. 向后开辟一杯数据
func (d *DataBuff)AddDataACup(maxCup int) {
	if 0 == d.LenData {
		d.LenRemove = 0
	}
	if GROOVE_CAP != d.lenGroove  {
		if d.LenData + maxCup < GROOVE_CAP{
			d.shortFlag += 1
		}else {
			d.shortFlag = 0
		}
		if d.shortFlag >= 5{
			d.crossBufferMove(1)
		}
	}
	if d.lenGroove-d.LenRemove-d.LenData < maxCup {
		if d.lenGroove-d.LenData > maxCup{
			d.bufferMove()
		}else {
			page := (d.LenData + maxCup)/GROOVE_CAP + 1
			d.crossBufferMove(page)
		}
	}
}

//I/O get data address
func (d *DataBuff)DataAppend(addData []byte)() {
	length := len(addData)
	if GROOVE_CAP != d.lenGroove  {
		if d.LenData + length < GROOVE_CAP{
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
	if (d.lenGroove - d.LenData - length ) >=  0{
		if (d.lenGroove - d.LenRemove - d.LenData ) >=  length {
			d.bufferAppend(addData)
		}else{
			d.bufferMove()
			d.bufferAppend(addData)
		}
	}else{
		page := (d.LenData + length)/GROOVE_CAP + 1
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
	if d.LenData < 4 {
		popData = nil
	}
	length, _ := datastream.GetUint32(d.sGroove[d.LenRemove:])
	popData = d.bufferPop(int(length))
	return
}

//////////////////////////////////////////////////////////////////////////////////////////
// package function move buffer
func (d *DataBuff)bufferMove()  {
	copy(d.sGroove[0:d.LenData], d.sGroove[d.LenRemove:(d.LenRemove + d.LenData)])
	d.LenRemove = 0
}
// package function move buffer
func (d *DataBuff)crossBufferMove(page int)  {
	groove := make([]byte, GROOVE_CAP * page)
	copy(groove[0:d.LenData], d.sGroove[d.LenRemove:(d.LenRemove + d.LenData)])
	d.lenGroove = GROOVE_CAP * page
	d.LenRemove = 0
	d.shortFlag = 0
	d.sGroove = groove
}
//package function buffer append slices
func (d *DataBuff)bufferAppend(addData []byte)  {
	copy(d.sGroove[(d.LenRemove + d.LenData):(d.LenRemove + d.LenData)+ len(addData)], addData)
	d.LenData += len(addData)
}
//package function buffer pop slices
func (d *DataBuff)bufferPop(length int)(popData []byte)  {
	if length<= d.LenData {
		popData = d.sGroove[d.LenRemove:(d.LenRemove + length)]
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
//DataSlotWriteString16
func (d *DataBuff)DataSlotWriteString16(position int , wData string)()  {
	lenDat :=uint16(len(wData))
	d.DataSlotWriteUint16(position , lenDat)
	copy(d.sGroove[position+2:position+2+int(lenDat)],wData)
}
//DataSlotReadString16
func (d *DataBuff)DataSlotReadString16(position int )(string)  {
	length := d.DataSlotReadUint16(position)
	if 0!=length {
		return string(d.sGroove[position+2:position+2+int(length)])
	}else {
		return ""
	}
}
//DataSlotWriteString32
func (d *DataBuff)DataSlotWriteString32(position int , wData string)()  {
	lenDat :=uint32(len(wData))
	d.DataSlotWriteUint32(position , lenDat)
	copy(d.sGroove[position+4:position+4+int(lenDat)],wData)
}
//DataSlotReadString32
func (d *DataBuff)DataSlotReadString32(position int )(string)  {
	length := d.DataSlotReadUint32(position)
	if 0!=length {
		return string(d.sGroove[position+4:position+4+int(length)])
	}else {
		return ""
	}
}