package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

//创建用户
func NewUser(conn net.Conn, server *Server) *User {

	userAddress := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddress,
		Addr:   userAddress,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "已上线")
}
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "下线")
}
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//who
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//rename|张三
		newName := strings.Split(msg, "|")[1]
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName

			this.SendMsg("您已更新用户名:" + newName + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//to|zhangsan|哈哈哈哈
		toNmae := strings.Split(msg, "|")[1]
		if toNmae == "" {
			this.SendMsg("对方用户名不能为空\n")
			return
		}
		toUser, ok := this.server.OnlineMap[toNmae]
		if !ok {
			this.SendMsg("对方用户名不存在\n")
			return
		}
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容\n")
			return
		}
		toUser.SendMsg(this.Name + "对你说" + content)

		//toUser := this.server.OnlineMap[toNmae]
		//toUser.C <- "1111111"

	} else {
		this.server.BroadCast(this, msg)
	}
}

//监听当前user channel的方法,一旦有消息就发送客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
