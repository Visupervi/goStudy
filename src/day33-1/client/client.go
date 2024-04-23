package client

import (
	"fmt"
	"net"
)

func Client() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("client链接失败", err)
		return
	}

	fmt.Println("client Success", conn)

}
