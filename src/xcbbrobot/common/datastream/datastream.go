//This package is an operation that specializes in processing binary streams
//Add stream data, intercept and output stream data
package datastream

import (
	"encoding/binary"

)



///////////////////////////////////////////////////////////////////////////////////////////////////
//Add a uint8 to the data stream
func AddUint8(myuint8 uint8,inbyte []byte )(outbyte []byte) {
	bytes := make([]byte,1)
	bytes[0] = myuint8
	outbyte = append(inbyte,bytes...)
	return
}
//Take a uint8 from a data stream
func GetUint8(inbyte []byte)(myuint8 uint8,outbyte []byte){
	myuint8 = inbyte[0]
	outbyte = inbyte[1:]
	return
}
//Add a int8 to the data stream. But this is a rare operation
func AddInt8(myint8 int8,inbyte []byte) (outbyte []byte){
	bytes := make([]byte,1)
	if myint8 < 0 {
		bytes[0] = byte(2<<(8-1) + int(myint8))
	} else {
		bytes[0] = byte(myint8)
	}
	outbyte = append(inbyte,bytes...)
	return
}
//Add a uint16 to the data stream
func AddUint16(myuint16 uint16,inbyte []byte)(outbyte []byte){
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, myuint16)
	outbyte = append(inbyte,bytes...)
	return
}
//Take a uint16 from a data stream
func GetUint16(inbyte []byte)(myuint16 uint16,outbyte []byte){
	myuint16 = binary.LittleEndian.Uint16(inbyte)
	outbyte = inbyte[2:]
	return
}
//Add a uint32 to the data stream
func AddUint32(myuint32 uint32,inbyte []byte)(outbyte []byte){
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, myuint32)
	outbyte = append(inbyte,bytes...)
	return
}
//Take a uint32 from a data stream
func GetUint32(inbyte []byte)(myuint32 uint32,outbyte []byte){
	myuint32 = binary.LittleEndian.Uint32(inbyte)
	outbyte = inbyte[4:]
	return
}
//Add a uint64 to the data stream
func AddUint64(myuint64 uint64,inbyte []byte)(outbyte []byte){
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, myuint64)
	outbyte = append(inbyte,bytes...)
	return
}
//Take a uint64 from a data stream
func GetUint64(inbyte []byte)(myuint64 uint64,outbyte []byte){
	myuint64 = binary.LittleEndian.Uint64(inbyte)
	outbyte = inbyte[8:]
	return
}
//Add a []byte to the data stream
func AddBytes(inbyte []byte, addbyte[]byte)(outbyte []byte){
	outbyte = append(inbyte,addbyte...)
	return
}
//Add a string16 to the data stream
func AddString16(mystring string,inbyte []byte)(outbyte []byte){
	length :=uint16(len(mystring))
	outbyte = AddUint16(length,inbyte)
	outbyte = append(outbyte,mystring...)
	return
}
//Add a string32 to the data stream
func AddString32(mystring string,inbyte []byte)(outbyte []byte){
	length :=uint32(len(mystring))
	outbyte = AddUint32(length,inbyte)
	outbyte = append(outbyte,mystring...)
	return
}

//Take a String from a data stream
func GetString(inbyte []byte)(myString string,outbyte []byte){
	var length uint16
	datastring :=make([]byte,0)
	length,outbyte =GetUint16(inbyte)
	if 0!=length {
		datastring= outbyte[0:length]
		myString = string(datastring)
		outbyte = outbyte[length:]
	}else {
		myString = ""
	}
	return
}
