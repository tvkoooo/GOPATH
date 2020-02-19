package message

import (
	"encoding/binary"
	"sync"
	"xcbbrobot/common/datagroove"
)

type UriDecodeHandler func(d *datagroove.DataBuff, i interface{}, length int)
type MapUriFuncDecode struct {
	M map[uint32]UriDecodeHandler
	L *sync.RWMutex
}

func (p *MapUriFuncDecode) ZhuCe(uri uint32, i UriDecodeHandler) {
	p.L.Lock()
	p.M[uri] = i
	p.L.Unlock()
}

func (p *MapUriFuncDecode) UriDecodeHandlerInit() {
	p.M = make(map[uint32]UriDecodeHandler)
	p.L = new(sync.RWMutex)

}

//pack body
type PackBody struct {
	Body   []byte
	OffSet int
	Len    int
}

type MesIo interface {
	WriteBody(pB *PackBody)
	ReadBody(pB *PackBody)
}
type MesFun interface {
	kkkkkk(pB *PackBody)
}
type MesUri struct {
	UriMap map[uint32]MesIo
}

//StreamWriteByte
func (pB *PackBody) StreamWriteByte(wData byte) {
	pB.Body[pB.OffSet] = wData
	pB.OffSet += 1
}

//StreamReadByte
func (pB *PackBody) StreamReadByte() (out byte) {
	out = pB.Body[pB.OffSet]
	pB.OffSet += 1
	return out
}

//StreamWriteUint8
func (pB *PackBody) StreamWriteUint8(wData uint8) {
	pB.Body[pB.OffSet] = wData
	pB.OffSet += 1
}

//StreamReadUint8
func (pB *PackBody) StreamReadUint8() (out uint8) {
	out = pB.Body[pB.OffSet]
	pB.OffSet += 1
	return out
}

//StreamWriteInt8
func (pB *PackBody) StreamWriteInt8(wData int8) {
	if wData < 0 {
		pB.Body[pB.OffSet] = byte(2<<(8-1) + int(wData))
	} else {
		pB.Body[pB.OffSet] = byte(wData)
	}
}

//StreamReadInt8
func (pB *PackBody) StreamReadInt8() (out int8) {
	out = int8(pB.Body[pB.OffSet])
	pB.OffSet += 1
	return out
}

//StreamWriteUint16
func (pB *PackBody) StreamWriteUint16(wData uint16) {
	binary.LittleEndian.PutUint16(pB.Body[pB.OffSet:], wData)
	pB.OffSet += 2
}

//StreamReadUint16
func (pB *PackBody) StreamReadUint16() (out uint16) {
	out = binary.LittleEndian.Uint16(pB.Body[pB.OffSet:])
	pB.OffSet += 2
	return out
}

//StreamWriteUint32
func (pB *PackBody) StreamWriteUint32(wData uint32) {
	binary.LittleEndian.PutUint32(pB.Body[pB.OffSet:], wData)
	pB.OffSet += 4
}

//StreamReadUint32
func (pB *PackBody) StreamReadUint32() (out uint32) {
	out = binary.LittleEndian.Uint32(pB.Body[pB.OffSet:])
	pB.OffSet += 4
	return out
}

//StreamWriteUint64
func (pB *PackBody) StreamWriteUint64(wData uint64) {
	binary.LittleEndian.PutUint64(pB.Body[pB.OffSet:], wData)
	pB.OffSet += 8
}

//StreamReadUint64
func (pB *PackBody) StreamReadUint64() (out uint64) {
	out = binary.LittleEndian.Uint64(pB.Body[pB.OffSet:])
	pB.OffSet += 8
	return out
}

//StreamWriteString16
func (pB *PackBody) StreamWriteString16(wData *string) {
	lenDat := uint16(len(*wData))
	pB.StreamWriteUint16(lenDat)
	copy(pB.Body[pB.OffSet:], *wData)
	pB.OffSet += int(lenDat)
}

//StreamReadString16
func (pB *PackBody) StreamReadString16(out *string) {
	length := pB.StreamReadUint16()
	if 0 != length {
		*out = string(pB.Body[pB.OffSet : pB.OffSet+int(length)])
		pB.OffSet += int(length)
	} else {
		*out = ""
	}
}

//DataSlotWriteString32
func (pB *PackBody) StreamWriteString32(wData *string) {
	lenDat := uint32(len(*wData))
	pB.StreamWriteUint32(lenDat)
	copy(pB.Body[pB.OffSet:], *wData)
	pB.OffSet += int(lenDat)
}

//DataSlotReadString32
func (pB *PackBody) StreamReadString32(out *string) {
	length := pB.StreamReadUint32()
	if 0 != length {
		*out = string(pB.Body[pB.OffSet : pB.OffSet+int(length)])
		pB.OffSet += int(length)
	} else {
		*out = ""
	}
}
