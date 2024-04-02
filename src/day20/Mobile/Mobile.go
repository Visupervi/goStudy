package Mobile

import "fmt"

type Mobile struct {
	Name  string
	Color string
}

func (m Mobile) Start() {
	fmt.Println("手机开始工作。。。")
}

func (m Mobile) Stop() {
	fmt.Println("手机停止作。。。")
}

func (m *Mobile) Call() {
	fmt.Println("我是手机我可以打电话")
}
