package model

import "net"

// 在登陆成功后对其初始化
var CurrentUser CurUser

type CurUser struct {
	Conn net.Conn `json:"conn"`
	User `json:"user"`
}
