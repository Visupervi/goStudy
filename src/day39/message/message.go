package message

import (
	"day39/client/model"
)

const (
	LoginMsgType        = "LoginMsg"
	LoginResType        = "LoginResult"
	RegistryType        = "RegistryType"
	NotifyUserStatusMsg = "NotifyUserStatus"
)
const (
	Online = iota
	Offline
	BusyLine
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type IMessageHandle interface {
	LoginHandle()
	RegistryHandle()
}

type LoginMsg struct {
	UserId   int    `json:"userId"`
	Pwd      string `json:"pwd"`
	UserName string `json:"userName"`
}

type LoginResult struct {
	Code        int         `json:"code"`
	Error       string      `json:"error"`
	Result      *model.User `json:"result"`
	OnlineUsers []int       `json:"onlineUsers"`
}

type RegistryMsg struct {
	User model.User `json:"user"`
}

type RegistryResult struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
type NotifyUserStatus struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
