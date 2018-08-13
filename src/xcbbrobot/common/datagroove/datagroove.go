package datagroove

import (
	"unsafe"
	"xcbbrobot/common/datastream"
)

const GROOVE_CAP = 1024

type DataBuff struct {
	sGroove       []byte
	lenGroove     int
	lenBaseGroove int
	lenRemove     int
	lenData       int
	shortFlag     byte
}
// I/O init buffer
func (d *DataBuff)BufferInit() () {
	groove := make([]byte, 0, GROOVE_CAP )
	d.lenGroove = GROOVE_CAP
	d.lenBaseGroove = GROOVE_CAP
	d.lenRemove = 0
	d.lenData = 0
	d.shortFlag = 0
	d.sGroove = groove[0:d.lenGroove]
}
//I/O get buffer address
func (d *DataBuff)GetBufferAddress() unsafe.Pointer {
	return unsafe.Pointer(&(d.sGroove[0]))
}
//I/O get data address
func (d *DataBuff)GetDataAddress(pos int) unsafe.Pointer {
	return unsafe.Pointer(&(d.sGroove[d.lenRemove + pos]))
}
//I/O get data address
func (d *DataBuff)DataAppend(addData []byte)() {
	length := len(addData)
	if d.lenBaseGroove != d.lenGroove  {
		if d.lenData + length < d.lenBaseGroove{
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
		page := (d.lenData + length)/d.lenBaseGroove + 1
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
	groove := make([]byte, 0, GROOVE_CAP * page)
	copy(groove[0:d.lenData], d.sGroove[d.lenRemove:(d.lenRemove + d.lenData)])
	d.lenGroove = GROOVE_CAP * page
	d.lenRemove = 0
	d.shortFlag = 0
	d.sGroove = groove[0:d.lenGroove]
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