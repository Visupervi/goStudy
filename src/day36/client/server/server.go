// Package server
//
//保持场链接
//
///**

package server

import (
	"day36/utils"
	"fmt"
	"net"
	"os"
)

func ShowMenu(userId int) {
	var key int
	//var userId int
	//var pwd string
	loop := true
	for loop {
		fmt.Printf("------------------------------welcome %d to多人聊天系统------------------------------\n", userId)
		fmt.Println("\t\t\t 1 显示在线用户列表")
		fmt.Println("\t\t\t 2 发送消息")
		fmt.Println("\t\t\t 3 信息列表")
		fmt.Println("\t\t\t 4 退出系统")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
			loop = false
		case 2:
			fmt.Println("发送消息")
			loop = false
		case 3:
			fmt.Println("信息列表")
			loop = false
		case 4:
			fmt.Println("退出系统....")
			//loop = false
			os.Exit(0)

		default:
			fmt.Println("输入有误")
		}

	}

}

func LiveProcess(conn net.Conn) {

	p := &utils.Pipeline{
		Conn: conn,
	}
	for {
		fmt.Println("等待服务端的数据。。。")
		res, err := p.ReadPkg()
		if err != nil {
			fmt.Println("数据失败", err)
			return
		}

		fmt.Println("res=", res)
	}
}
