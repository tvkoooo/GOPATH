package typechange

import (
	"encoding/binary"
	"strconv"
)

/////////////////////////////////////////////////////////////////////////////////
func String2Int(stringIn *string, intOut *int) (ferr error) {
	*intOut, ferr = strconv.Atoi(*stringIn)
	return ferr
}
func String2IntRe(stringIn string) (intOut int) {
	intOut, _ = strconv.Atoi(stringIn)
	return intOut
}

/////////////////////////////////////////////////////////////////////////////////
func String2Int64(stringIn *string, intOut *int64) (ferr error) {
	*intOut, ferr = strconv.ParseInt((*stringIn), 10, 64)
	return ferr
}
func String2Int64Re(stringIn string) (intOut int64) {
	intOut, _ = strconv.ParseInt(stringIn, 10, 64)
	return intOut
}

/////////////////////////////////////////////////////////////////////////////////
func Int2String(intIn *int, stringOut *string) {
	*stringOut = strconv.Itoa(*intIn)
}
func Int2StringRe(intIn int) (stringOut string) {
	stringOut = strconv.Itoa(intIn)
	return stringOut
}

/////////////////////////////////////////////////////////////////////////////////
func Int642String(intIn *int64, stringOut *string) {
	*stringOut = strconv.FormatInt(*intIn, 10)
}
func Int642StringRe(intIn int64) (stringOut string) {
	stringOut = strconv.FormatInt(intIn, 10)
	return stringOut
}

/////////////////////////////////////////////////////////////////////////////////
func Slice2String(sliceIn *[]byte, stringOut *string) {
	*stringOut = string((*sliceIn)[:])
}
func Slice2StringRe(sliceIn []byte) (stringOut string) {
	stringOut = string(sliceIn[:])
	return stringOut
}

/////////////////////////////////////////////////////////////////////////////////
func String2Slice(stringIn *string, sliceOut *[]byte) {
	*sliceOut = []byte(*stringIn)
}
func String2SliceRe(stringIn string) (sliceOut []byte) {
	sliceOut = []byte(stringIn)
	return sliceOut
}

////////////////////////////////////////////////////////////////////
func Uint8_2_Slice(myuint8 uint8) []byte {
	bytes := make([]byte, 1)
	bytes[0] = myuint8
	return bytes
}

func Slice_2_Uint8(s []byte) uint8 {
	return s[0]
}

func Int8_2_Slice(myint8 int8) []byte {
	bytes := make([]byte, 1)
	if myint8 < 0 {
		bytes[0] = byte(2<<(8-1) + int(myint8))
	} else {
		bytes[0] = byte(myint8)
	}
	return bytes
}

func Uint16_2_Slice(myuint16 uint16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, myuint16)
	return bytes
}
func Slice_2_uint16(s []byte) uint16 {
	return binary.LittleEndian.Uint16(s)
}

func Uint32_2_Slice(myuint32 uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, myuint32)
	return bytes
}
func Slice_2_uint32(s []byte) uint32 {
	return binary.LittleEndian.Uint32(s)
}

func Uint64_2_Slice(myuint64 uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, myuint64)
	return bytes
}
func Slice_2_uint64(s []byte) uint64 {
	return binary.LittleEndian.Uint64(s)
}

func string_2_Slice(mystring string) (outbyte []byte) {
	inbyte := make([]byte, 0)
	outbyte = append(inbyte, mystring...)
	return outbyte
}
func Slice_2_string(s []byte) (mystring string) {
	if 0 != len(s) {
		mystring = string(s)
	} else {
		mystring = ""
	}
	return
}
