package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type sendword struct {
	Pid       uint64
	Mid       uint64
	Datawords string
}

func sender(conn net.Conn) {
	datasends := &sendword{
		10086,
		800122001,
		"hello world!",
	}
	datasendj, err := json.Marshal(datasends)
	if err != nil {
		fmt.Print(err.Error())
	}

	datarec := &sendword{}
	err = json.Unmarshal([]byte(datasendj), &datarec)


	fmt.Print(datasends)
	fmt.Print(string(datasendj),len(datasendj))
	fmt.Print(datarec)
	conn.Write([]byte(datasendj))
	fmt.Println("send over")

}

func main() {
	server := "127.0.0.1:9090"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)

}
