package smProcess

import (
	"day39/client/model"
	"day39/message"
	"day39/utils"
	"encoding/json"
	"fmt"
)

type SmProcess struct {
}

func (sp *SmProcess) SendGroupMsg(content string) (err error) {
	// 创建一个消息
	var msg message.Message
	msg.Type = message.SmsMsgType
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = model.CurrentUser.UserId

	data, err := json.Marshal(smsMsg)

	if err != nil {
		fmt.Println("SendGroupMsg error", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("SendGroupMsg error", err)
		return
	}
	pip := &utils.Pipeline{
		Conn: model.CurrentUser.Conn,
	}

	err = pip.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMsg error", err)
		return
	}
	return
}
