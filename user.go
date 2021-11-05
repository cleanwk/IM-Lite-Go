package main

import "net"

type User struct {
	Name    string
	Address string
	C       chan string
	conn    net.Conn
}

//创建用户的api
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    userAddr,
		Address: userAddr,
		C:       make(chan string),
		conn:    conn,
	}

	go user.ListenMessage()

	return user
}

//监听当前user channel，一旦有消息，直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
