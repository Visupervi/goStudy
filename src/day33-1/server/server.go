package server

import (
	"fmt"
	"net"
)

func Start() {
	fmt.Println("服务器开始监听了")
	ln, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败")
		return
	}
	defer ln.Close()

	for {

		fmt.Println("等待链接。。。")
		conn, err := ln.Accept() // 返回的是interface

		if err != nil {
			fmt.Println("Accept Error:", err)
		} else {
			fmt.Printf("链接对象=%v,客户端IP地址=%v\n", conn, conn.RemoteAddr().String())
		}

	}

	//fmt.Println("监听成功", ln)
}
