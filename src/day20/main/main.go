package main

import (
	"day20/Camera"
	"day20/Computer"
	"day20/Mobile"
	_interface "day20/interface"
	"day20/student"
	"fmt"
)

func main() {

	//account := account.Account{"浦发", "123", 10}
	computer := Computer.Computer{}
	camera := Camera.Camera{}
	mobile := Mobile.Mobile{Name: "三星", Color: "black"}
	computer.Working(camera)
	computer.Working(mobile)

	stu := student.Stu{}

	var A _interface.AInterface = stu
	var B _interface.BInterface = stu
	A.Test02()
	A.Test01()
	fmt.Println("ok", A, B)

}

// 继承

// interface
// 在go中多态主要是依靠interface实现的

// 接口的快速入门
