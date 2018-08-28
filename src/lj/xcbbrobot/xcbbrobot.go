package main

import (
	"fmt"
	"lj/xcbblinktest/datastream"
	"net"
	"os"
	"time"
)

func main() {
	server := "59.110.125.134:30301"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	conn, err1 := net.DialTCP("tcp", nil, tcpAddr)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	var outbyte []byte
	sendstream := make([]byte, 0)
	length := uint32(13 + 4 + 4 + 2)
	outbyte = datastream.AddUint32(length, sendstream)

	var (
		Uri     uint32
		Sid     uint16
		Rescode uint16
		Tag     uint8
	)
	Uri = (101 << 8) | 23
	Sid = 0
	Rescode = 200
	Tag = 1

	outbyte = datastream.AddUint32(Uri, outbyte)
	outbyte = datastream.AddUint16(Sid, outbyte)
	outbyte = datastream.AddUint16(Rescode, outbyte)
	outbyte = datastream.AddUint8(Tag, outbyte)

	var (
		id     uint32
		PIType uint32
		PIPass string
	)
	id = 0
	PIType = 64
	PIPass = ""

	outbyte = datastream.AddUint32(id, outbyte)
	outbyte = datastream.AddUint32(PIType, outbyte)
	outbyte = datastream.AddString16(PIPass, outbyte)

	fmt.Println("send stream", outbyte[:length])

	go func() {
		time.Sleep(3E9)
		fmt.Println("send start")
		conn.Write([]byte(outbyte))
	}()

	recevdata := make([]uint8, 4096)
	fmt.Println("rece start 1")
	for {
		fmt.Println("rece start")
		count, err := conn.Read(recevdata[0:256])
		fmt.Println("rece over")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			break
		}
		if count != 0 {
			fmt.Println("rec Binary stream", recevdata[:count])
			fmt.Println("connect close success LocalAddr:", conn.LocalAddr(), "RemoteAddr", conn.RemoteAddr(), "timenow:", time.Now().Format("2006-01-02 15:04:05"), "\r\n")
		}
	}

}
func recvdata(conn net.Conn)  {
	recevdata := make([]uint8, 4096)
	for {
		count, err := conn.Read(recevdata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			break
		}
		if count != 0 {
			fmt.Println("rec Binary stream", recevdata[:count])
			fmt.Println("connect close success LocalAddr:", conn.LocalAddr(), "RemoteAddr", conn.RemoteAddr(), "timenow:", time.Now().Format("2006-01-02 15:04:05"), "\r\n")
		}
	}

}