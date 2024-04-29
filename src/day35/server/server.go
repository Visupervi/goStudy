package server

import (
	"day35/controller"
	"fmt"
	"net"
)

func Start() {

	fmt.Println("开始监听")
	listener, err := net.Listen("tcp", "localhost:9889")

	if err != nil {
		fmt.Println("启动服务失败", err)
		return
	}

	defer listener.Close()

	for {
		fmt.Println("等待链接。。。")
		conn, err := listener.Accept() // 返回的是interface

		if err != nil {
			fmt.Println("Accept Error:", err)
		}
		p := &controller.Process{
			Conn: conn,
		}
		p.Process()
		//go controller.Process()

	}

	fmt.Println("监听对象", listener)
}
