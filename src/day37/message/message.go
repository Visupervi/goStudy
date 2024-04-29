package message

import (
	"day37/client/model"
)

const (
	LoginMsgType = "LoginMsg"
	LoginResType = "LoginResult"
	RegistryType = "RegistryType"
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
	Code   int    `json:"code"`
	Error  string `json:"error"`
	Result string `json:"result"`
}

type RegistryMsg struct {
	User model.User `json:"user"`
}

type RegistryResult struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
