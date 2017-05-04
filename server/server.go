package main

import (
	"fmt"
	"os"
	"net"
)

func echoHandler(conns *map[string]net.Conn,messages chan string){
	for{
		msg:= <- messages
		fmt.Println(msg)

		for key,value := range *conns {
			fmt.Println("connection is connected from ...",key)
			_,err :=value.Write([]byte(msg))
			if(err != nil){
				fmt.Println(err.Error())
				delete(*conns,key)
			}

		}
	}
}

func Handler(conn net.Conn, messages chan string) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for{
		length, _ := conn.Read(buf)
		if length > 0 {
			buf[length] = 0
		}
		//fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))
		reciveStr :=string(buf[0:length])
		messages <- reciveStr
	}
}

func StartServer(port string){
	service:=":"+port //strconv.Itoa(port);
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println(err.Error())
	}
	l,err := net.ListenTCP("tcp",tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
	}

	conns := make(map[string]net.Conn)
	messages := make(chan string, 10)

	//启动服务器广播线程
	go echoHandler(&conns,messages)

	for  {
		fmt.Println("Listening ...")
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Accepting ...")
		conns[conn.RemoteAddr().String()] = conn
		go Handler(conn, messages)

	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong pare")
		os.Exit(0)
	}

	StartServer(os.Args[1])
}
