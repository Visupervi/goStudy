package Camera

import "fmt"

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}

func (c Camera) Stop() {
	fmt.Println("相机停止作。。。")
}
