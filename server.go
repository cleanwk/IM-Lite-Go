package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

//创建server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//todo 处理业务代码
	fmt.Println("链接建立成功")
}

//启动服务器接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	//close listen socket
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err : ", err)
			continue
		}
		//do handler
		go this.Handler(conn)

	}

}
