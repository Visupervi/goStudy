package userProcess

import (
	"day37/client/server"
	"fmt"
)
import (
	"day37/message"
	"day37/utils"
	"encoding/json"
	"net"
)

type UserProcess struct {
	//Conn net.Conn
	//UserId int
	//Pwd string
}

func (up *UserProcess) Login(userId int, pwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:9889")
	defer conn.Close()
	if err != nil {
		fmt.Println("链接失败,error=", err)
		return
	}
	fmt.Println("userId\n", userId)
	fmt.Println("pwd\n", pwd)
	mg := message.LoginMsg{
		UserId:   userId,
		Pwd:      pwd,
		UserName: "",
	}
	data, err := json.Marshal(mg)
	if err != nil {
		fmt.Println("mg序列化失败")
		return
	}
	// 先给服务器发送消息长度
	// 再给服务器发送消息本体
	sendMsg := message.Message{
		Type: message.LoginMsgType,
		Data: string(data),
	}
	sendMessage, err := json.Marshal(sendMsg)

	if err != nil {
		fmt.Println("sendMsg序列化失败")
		return
	}

	p := &utils.Pipeline{
		Conn: conn,
	}
	err = p.WritePkg(sendMessage)

	if err != nil {
		fmt.Println("客户端写错误")
	}
	// 客户端读消息
	for {

		p := utils.Pipeline{
			Conn: conn,
		}
		serverMsg, error := p.ReadPkg()

		// 将数据反序列化

		//response

		if error != nil {
			fmt.Println("server reade error", error)
			return
		}

		var response message.LoginResult
		//_, err := json.Unmarshal([]byte(serverMsg.Data), &response)
		fmt.Println("sendMsg.Data=", serverMsg.Data)
		err = json.Unmarshal([]byte(serverMsg.Data), &response)

		if error != nil {
			fmt.Println("Unmarshal sendMsg.Data error", error)
			return
		}

		fmt.Println("response=", response)

		if response.Code == 200 {
			go server.LiveProcess(conn)
			server.ShowMenu(userId)
			fmt.Println("login Success")
		} else {
			fmt.Println(response.Error)
		}

		//fmt.Println("serverMsg", serverMsg)

	}

	// 需要处理服务器返回的来消息
	//LoginHandle(conn, )

	return
}
