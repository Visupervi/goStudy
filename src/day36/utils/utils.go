package utils

import (
	"day36/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Pipeline struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (p *Pipeline) WritePkg(data []byte) (err error) {
	// 先发送长度
	var dataLen uint32
	dataLen = uint32(len(data))
	//var byt [4]byte

	//fmt.Println("dataLen=", dataLen)
	//  需要转成bytes,使用这个方式
	binary.BigEndian.PutUint32(p.Buf[0:4], dataLen)
	//fmt.Println("byt=", p.Buf)
	n, err := p.Conn.Write(p.Buf[:4])

	if err != nil || n != 4 {
		fmt.Println("发送失败", err)
		return
	}

	fmt.Println("消息长度发送成功,长度=", n)

	fmt.Println("********开始发送消息********")

	n, err = p.Conn.Write(data)

	if err != nil || n != int(dataLen) {
		fmt.Println("发送消息失败")
		return
	}
	fmt.Println("********发送消息结束********")
	return
}

// ReadPkg 读取包的内容
func (p *Pipeline) ReadPkg() (msg message.Message, error error) {
	//buf := make([]byte, 8096)

	fmt.Println("等待读取数据")
	_, err := p.Conn.Read(p.Buf[:4])
	if err != nil {
		error = errors.New("服务端读取失败")
		fmt.Println("服务端读取失败")
		return
	}
	//if n != 4 {
	//	fmt.Println("服务端读取失败", n)
	//	return
	//}
	//var pkgLen uint32
	pkgLen := binary.BigEndian.Uint32(p.Buf[0:4])
	n, err := p.Conn.Read(p.Buf[:pkgLen])

	if err != nil || uint32(n) != pkgLen {
		error = errors.New("conn read fail")
		fmt.Println("conn read fail", err)
		return
	}
	//fmt.Println("客户端发送消息OK", buf[:4])

	// 把pkg反序列化

	err = json.Unmarshal(p.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("反序列化失败", err)
		error = errors.New("反序列化失败")
		return
	}

	return
}
