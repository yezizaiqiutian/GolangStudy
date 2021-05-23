package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	//链接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net Dial error:", err)
		return nil
	}

	client.conn = conn
	return client

}

//处理server回应的消息,直接显示到标准输出
func (client *Client) DealResponse() {
	//一旦client.conn有数据,就会直接copy到stdout标准输出上,永久阻塞监听
	io.Copy(os.Stdout, client.conn)

}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1:公聊模式")
	fmt.Println("2:私聊模式")
	fmt.Println("3:更新用户名")
	fmt.Println("0:退出")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("请输入合法范围内的数字")
		return false
	}
}

func (client *Client) ReName() bool {
	fmt.Println(">>>>>>请输入新用户名")

	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	} else {
		return true
	}

}

func (client *Client) PublicChat() {

	var chatMsg string
	//提示用户输入消息,
	fmt.Println(">>>>>>请输入聊天内容,exit表示退出")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		//发给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break

			}
		}
		chatMsg = ""
		fmt.Println(">>>>>>请输入聊天内容,exit表示退出")
		fmt.Scanln(&chatMsg)
	}
}

func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write error ", err)
		return
	}
}

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println(">>>>>>请输入聊天对象[用户名],exit 退出")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
		fmt.Println(">>>>>>请输入消息内容,exit 退出")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			//消息为空则不发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write error ", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>>>>请输入消息内容,exit 退出")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		fmt.Println(">>>>>>请输入聊天对象[用户名],exit 退出")
		fmt.Scanln(&remoteName)
	}
}

func (client Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		//根据不同的type处理不同的类型
		switch client.flag {
		case 1:
			//
			fmt.Println("公聊模式")
			client.PublicChat()
			break
		case 2:
			fmt.Println("私聊模式")
			client.PrivateChat()
			break
		case 3:
			fmt.Println("改名模式")
			client.ReName()
			break

		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器地址(默认127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认8888)")
}

func main() {
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>连接服务器失败")
		return
	}

	//单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>>>>>连接服务器成功")
	//启动客户端的业务
	client.Run()
}
