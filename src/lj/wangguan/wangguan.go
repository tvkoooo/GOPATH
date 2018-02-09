package main
import (
	"fmt"
	"net"
	"log"
	"os"
	"encoding/json"
)
type sendword struct {
	Pid       uint64
	Mid       uint64
	Datawords string
}

func main() {

	//建立socket，监听端口
	netListen, err := net.Listen("tcp", "localhost:9090")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}
//处理连接
func handleConnection(conn net.Conn) {

	buffer := make([]byte, 2048)

	for {

		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		fmt.Print(string(buffer[:n]))
		datarec := &sendword{}
		err = json.Unmarshal([]byte(buffer[0:n]), &datarec)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
		}
		Log(conn.RemoteAddr().String(), "receive data string:\n",n ,datarec,string(buffer[:n]))

		//转发包裹
		server := "127.0.0.1:9092"
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
		datasendj, err := json.Marshal(datarec)
		if err != nil {
			fmt.Print(err.Error())
		}
		err = json.Unmarshal([]byte(datasendj), &datarec)
		conn.Write([]byte(datasendj))
		fmt.Println("send over")
	}

}
func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
