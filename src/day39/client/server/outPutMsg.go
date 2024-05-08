package server

import (
	"day39/message"
	"encoding/json"
	"fmt"
)

func OutGroupMsg(msg *message.Message) {
	var smsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsg)
	if err != nil {
		fmt.Println("OutGroupMsg Unmarshal error", err.Error())
		return
	}
	//content := smsg.Content
	// 显示信息
	info := fmt.Sprintf("用户：%d对大家说：%s", smsg.UserId, smsg.Content)

	fmt.Println(info)
}
