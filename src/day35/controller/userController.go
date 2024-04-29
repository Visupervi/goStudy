package controller

import (
	"day35/model"
	"day37/message"
	"day37/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserController struct {
	Conn net.Conn
	Msg  message.Message
}

func (u *UserController) LoginHandle(msg *message.Message) (err error) {
	var msgData message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &msgData)

	if err != nil {
		fmt.Println("LoginHandle 反序列化失败", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.LoginResType
	var res message.LoginResult

	// 从redis获取数据
	// dao.GetUser(msgData.UserId)
	// 假设用户id =100, 密码等于123456就是成功的否则失败
	//r, err := redis.String(c.Do("get", "name"))

	// 用户输入的用户名和密码
	user := model.User{}

	// 接口返回的
	result, err1 := user.Login(msgData.UserId, msgData.Pwd)

	if err1 != nil {
		fmt.Println("LoginHandle-----登陆失败-----", result)
		res.Code = 500
		res.Error = err1.Error()
	} else {
		res.Code = 200
		res.Error = "0"

		// 序列化操作
		strData, e := json.Marshal(result)

		if e != nil {
			fmt.Println("son.Marshal(result) 失败", e)
			return
		}
		res.Result = string(strData)
	}

	//fmt.Println("result==>", result)
	//json.Unmarshal([]byte(result), &user)
	// 将res序列化
	data, err := json.Marshal(res)

	if err != nil {
		fmt.Println("LoginHandle res json.Marshal 失败", err)
		return
	}
	resMsg.Data = string(data)

	// 将resMsg序列化
	strResMsg, err := json.Marshal(resMsg)

	if err != nil {
		fmt.Println("LoginHandle strResMsg json.Marshal 失败", err)
		return
	}

	pipe := &utils.Pipeline{
		Conn: u.Conn,
	}
	err = pipe.WritePkg(strResMsg)

	if err != nil {
		fmt.Println("服务端写失败", err)
		return
	}

	return
}

func RegistryProcess(conn net.Conn, msg *message.Message) (err error) {
	return
}