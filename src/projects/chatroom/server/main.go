package main

import (
	"net"
	"log"
		"projects/chatroom/server/process"
)

func main(){
	listen,err :=net.Listen("tcp",":8888")
	if err!=nil{
		log.Println("net.Listen err=",err)
		return
	}
	defer listen.Close()

	for {
		conn,err:=listen.Accept()
		if err !=nil{
			log.Println("listen.Accept err=",err)
			return
		}
		go processConn(conn)
	}
}

func processConn(conn net.Conn)  {
	defer conn.Close()
	process := &process.Process{
		Conn : conn,
	}

	err := process.ProcessData()
	if err != nil {
		log.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}