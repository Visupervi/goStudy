package login

import (
	"day35-1/message"
	binary "encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func Login(userId int, pwd string) (err error) {

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

	conn, err := net.Dial("tcp", "localhost:9889")
	defer conn.Close()
	if err != nil {
		fmt.Println("链接失败,error=", err)
		return
	}

	var dataLen uint32
	dataLen = uint32(len(sendMessage))
	var byt [4]byte

	//fmt.Println("dataLen=", dataLen)
	binary.BigEndian.PutUint32(byt[0:4], dataLen)
	fmt.Println("byt=", byt)
	n, err := conn.Write(byt[:4])

	if err != nil || n != 4 {
		fmt.Println("发送失败", err)
		return
	}

	fmt.Println("客户端消息长度发送成功,长度=\n", n)

	fmt.Println("开始发送消息")

	//_, err = conn.Write(sendMessage)
	//
	//if err != nil {
	//	fmt.Println("发送消息失败")
	//	return err
	//}
	fmt.Println("发送消息结束", len(sendMessage))

	return nil
}
