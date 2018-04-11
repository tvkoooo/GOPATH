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
	netListen, err := net.Listen("tcp", "localhost:9092")
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
		fmt.Print(n)
		fmt.Println(buffer[:n])
		g_data,ndata :=messagehead.Get_data(buffer,uint32(n))
		g_str_data := string(g_data[:ndata])
		fmt.Println(g_data[:ndata])
		fmt.Println(ndata)
		fmt.Println(g_str_data)
		//datarec := &sendword{}
		//err = json.Unmarshal([]byte(buffer[0:n]), &datarec)
		//if err != nil {
		//	Log(conn.RemoteAddr().String(), " connection error: ", err)
		//}
		//Log(conn.RemoteAddr().String(), "receive data string:\n",n ,datarec,string(buffer[:n]))

		//返回包裹
		server := "127.0.0.1:6090"
		tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			os.Exit(1)
		}
		fconn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println("connect success")

		f_send := "ni ma bi ya!"
		fsendbyte := []byte(f_send)
		flen_send := uint32(len(fsendbyte))
		fbuffer_send := buffer[:n]
		new_data,new_length :=messagehead.Change_data(fbuffer_send,fsendbyte,flen_send)
		fmt.Println(new_data)
		fmt.Println(new_length)
		f_say,f_say_l := messagehead.Get_data(new_data,new_length)
		fmt.Println(f_say)
		fmt.Println(f_say_l)
		fconn.Write([]byte(fbuffer_send))
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