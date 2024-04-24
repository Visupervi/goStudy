package server

import (
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
		go Process(conn)

	}

	fmt.Println("监听对象", listener)
}

// Process 读取客户断发送的消息
func Process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 8096)

		fmt.Println("等待读取数据")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("服务端读取失败")
			return
		}
		//if n != 4 {
		//	fmt.Println("服务端读取失败", n)
		//	return
		//}
		fmt.Println("客户端发送消息OK", buf[:4])
	}

	return

}
