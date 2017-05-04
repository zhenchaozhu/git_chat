package main

import (
	"fmt"
	"os"
	"net"
)

func chatSend(conn net.Conn){

	var input string
	username := conn.LocalAddr().String()
	for {
		fmt.Scanln(&input)
		if input == "/quit"{
			fmt.Println("ByeBye..")
			conn.Close()
			os.Exit(0)
		}

		lens,err :=conn.Write([]byte(username + " Say :::" + input))
		fmt.Println(lens)
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			break
		}
	}
}

func StartClient(tcpaddr string){

	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	if err != nil {
		fmt.Println(err.Error())
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
	}
	//启动客户端发送线程
	go chatSend(conn)

	//开始客户端轮训
	buf := make([]byte,1024)
	for{

		length, _ := conn.Read(buf)
		fmt.Println(string(buf[0:length]))

	}
}

func main(){

	if len(os.Args) != 2 {
		fmt.Println("Wrong pare")
		os.Exit(0)
	}
	StartClient(os.Args[1])
}