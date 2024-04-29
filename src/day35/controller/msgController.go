package controller

import (
	"day37/message"
	"day37/utils"
	"fmt"
	"net"
)

type Process struct {
	Conn net.Conn
}

// Process 读取客户断发送的消息
func (p *Process) Process() {
	defer p.Conn.Close()
	pipe := &utils.Pipeline{
		Conn: p.Conn,
	}
	for {
		msg, err := pipe.ReadPkg()

		if err != nil {
			fmt.Println("server reade error", err)
			return
		}
		p.MessageHandle(&msg)

		//fmt.Println("数据等于=", msg)
		fmt.Println("客户端发送的数据等于", msg)

	}

	return

}
func (p *Process) MessageHandle(msg *message.Message) (err error) {

	switch msg.Type {
	case message.LoginMsgType:
		//	处理登陆
		//(conn, msg)

		uc := UserController{
			Conn: p.Conn,
		}

		uc.LoginHandle(msg)
	case message.RegistryType:
		//	处理注册
		//RegistryProcess(, msg)
		uc := UserController{
			Conn: p.Conn,
		}
		uc.RegistryProcess(msg)
	case message.LoginResType:
		//LoginResultHandle(conn, msg)

	}
	return
}
