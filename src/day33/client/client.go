package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("client链接失败", err)
		return
	}

	fmt.Println("client Success", conn)

	for {
		reader := bufio.NewReader(os.Stdin)

		str, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("err", err)
		}
		str = strings.Trim(str, " \r\n")
		if str == "exit" {
			fmt.Println("客户端退出")
			break
		}

		n, err := conn.Write([]byte(str + "\n"))

		if err != nil {
			fmt.Println("Write-err", err)
		}

		fmt.Printf("客户端发送的了%d个数据：", n)
	}

}
