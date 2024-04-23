package main

import "day33-1/server"

// tcp

// 网络编程
// tcp/ip 协议， 物联网输会示用// 物理层、链路层、网络层、传输层、会话层、表示层和应用层
// C/S

// ip地址
// 电脑上一共有65535个端口
// 22 ssh 23 telnet 21 ftp 25 smtp 80 iis 7 echo

// 1-1024是固定

// 1025-65535是动态

// 监听端口8888
// 接收客户端的tcp链接，建立客户端与服务端的链接
//
func main() {
	server.Start()
}
