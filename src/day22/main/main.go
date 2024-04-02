package main

import (
	"day20/Camera"
	"day20/Computer"
	"day20/Mobile"
	_interface "day20/interface"
	"day22/Monkey"
	"fmt"
)

type Point struct {
	X string
	Y string
}

func main() {

	//var heroes Hero.HeroSlice
	//
	//for i := 0; i < 10; i++ {
	//	hero := Hero.Hero{Age: rand.Intn(100), Name: fmt.Sprintf("英雄～%d", rand.Intn(100))}
	//	heroes = append(heroes, hero)
	//}
	//fmt.Println(heroes)
	//sort.Sort(heroes)
	//fmt.Println(heroes)

	monkey := Monkey.LittleMonkey{Monkey: Monkey.Monkey{Name: "孙悟空"}}
	monkey.Climbing()
	monkey.Flying()
	monkey.Swimming()

	var usbs [3]_interface.IUsb
	usbs[0] = Camera.Camera{
		Name: "nikon",
	}
	usbs[1] = Camera.Camera{
		Name: "Canon",
	}
	usbs[2] = Mobile.Mobile{
		Name:  "小米",
		Color: "black",
	}
	fmt.Println(usbs)

	var a interface{}

	var point Point = Point{"1", "2"}
	a = point
	var b Point
	b = a.(Point) // 类型断言 assert

	fmt.Println(b)
	var computer Computer.Computer
	for _, usb := range usbs {
		computer.Working(usb)
		if usb, ok := usb.(Camera.Camera); ok { // 断言
			usb.Photo()
		}

		if usb, ok := usb.(Mobile.Mobile); ok { // 断言
			usb.Call()
		}
	}

}

// 接口体现多态的两种方式
// 多态参数
// 多态数组
