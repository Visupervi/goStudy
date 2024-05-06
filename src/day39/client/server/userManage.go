package server

import (
	"day39/client/model"
	"day39/message"
	"fmt"
)

var ClientOnlineUsers map[int]*model.User = make(map[int]*model.User, 100)

func UpdateOnlineUsers(msg *message.NotifyUserStatus) {
	user := &model.User{
		UserId: msg.UserId,
		Status: msg.Status,
	}
	ClientOnlineUsers[msg.UserId] = user
	OutPutOnline()
}

func OutPutOnline() {
	fmt.Println("当前在线用户列表:")
	for id, _ := range ClientOnlineUsers {
		fmt.Println("用户Id=", id)
		//fmt.Println("用户=", user)
		//fmt.Println("")
	}
}
