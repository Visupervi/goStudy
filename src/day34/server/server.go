package server

import (
	"fmt"
	"io"
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
		go ProcessHandle(conn)
	}

	//fmt.Println("监听成功", ln)
}

func ProcessHandle(conn net.Conn) {
	defer conn.Close()
	for {
		slice := make([]byte, 1024)

		fmt.Println("服务起在等待客户端输入信息")
		// 如果客户端美欧write，协程会一直堵塞在这里
		n, err := conn.Read(slice)

		if err != io.EOF {
			fmt.Println("客户端已退出", err)
			return
		}

		// 显示数据到终端
		fmt.Print("读取到的数据：", string(slice[:n]))
	}
}
