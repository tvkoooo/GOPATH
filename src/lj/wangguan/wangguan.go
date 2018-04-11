package main
import (
	"fmt"
	"net"
	"log"
	"os"
	//"encoding/json"
	"lj/messagehead"
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
	//建立socket，监听端口
	fnetListen, ferr := net.Listen("tcp", "localhost:6090")
	CheckError(ferr)
	defer fnetListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)

		fconn, ferr := netListen.Accept()
		if ferr != nil {
			continue
		}
		Log(fconn.RemoteAddr().String(), " tcp connect success")
		fhandleConnection(fconn)


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
		fmt.Println(err)
		fmt.Println(n)
		fmt.Print(buffer[:n])
		buffer_send := buffer[:n]
		g_len :=messagehead.Get_lenth(buffer)
		fmt.Println("\r\n")
		fmt.Println(g_len)
		g_pid :=messagehead.Get_pid(buffer)
		fmt.Println(g_pid)
		w_linkd :=uint32(332)
		messagehead.Change_linkd(buffer,w_linkd)

		//fmt.Print(string(buffer[:n]))
		//datarec := &sendword{}
		//err = json.Unmarshal([]byte(buffer[0:n]), &datarec)
		//if err != nil {
		//	Log(conn.RemoteAddr().String(), " connection error: ", err)
		//}
		//Log(conn.RemoteAddr().String(), "receive data string:\n",n ,datarec,string(buffer[:n]))

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
		//datasendj, err := json.Marshal(datarec)
		//if err != nil {
		//	fmt.Print(err.Error())
		//}
		//err = json.Unmarshal([]byte(datasendj), &datarec)
		fmt.Print(buffer[:n])
		conn.Write([]byte(buffer_send))
		fmt.Println("send over")
	}

}

//处理连接
func fhandleConnection(fconn net.Conn) {

	buffer := make([]byte, 2048)

	for {

		n, err := fconn.Read(buffer)
		if err != nil {
			Log(fconn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		fmt.Println(err)
		fmt.Println(n)
		fmt.Print(buffer[:n])
		buffer_send := buffer[:n]
		g_len :=messagehead.Get_lenth(buffer)
		fmt.Println("\r\n")
		fmt.Println(g_len)
		g_pid :=messagehead.Get_pid(buffer)
		fmt.Println(g_pid)
		w_linkd :=uint32(332)
		messagehead.Change_linkd(buffer,w_linkd)
		fmt.Print(buffer_send[:n])

		//fmt.Print(string(buffer[:n]))
		//datarec := &sendword{}
		//err = json.Unmarshal([]byte(buffer[0:n]), &datarec)
		//if err != nil {
		//	Log(conn.RemoteAddr().String(), " connection error: ", err)
		//}
		//Log(conn.RemoteAddr().String(), "receive data string:\n",n ,datarec,string(buffer[:n]))

		//转发包裹
		//server := "127.0.0.1:9092"
		//tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		//	os.Exit(1)
		//}
		//conn, err := net.DialTCP("tcp", nil, tcpAddr)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		//	os.Exit(1)
		//}
		//fmt.Println("connect success")




		//datasendj, err := json.Marshal(datarec)
		//if err != nil {
		//	fmt.Print(err.Error())
		//}
		//err = json.Unmarshal([]byte(datasendj), &datarec)


		//
		//fmt.Print(buffer[:n])
		//conn.Write([]byte(buffer_send))
		//fmt.Println("send over")
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
