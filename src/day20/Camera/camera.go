package Camera

import "fmt"

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}

func (c Camera) Stop() {
	fmt.Println("相机停止作。。。")
}

func (c *Camera) Photo() {
	fmt.Println("我可以照相")
}
