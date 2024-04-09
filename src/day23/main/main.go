package main

import "fmt"

type Point struct {
	X string
	Y string
}

func jumpType(items ...interface{}) {
	for i, item := range items {
		switch item.(type) {
		case bool:
			fmt.Printf("地%v个参数是bool, 值是%v\n", i, item)
		case float32:
			fmt.Printf("地%v个参数是float32, 值是%v\n", i, item)
		case float64:
			fmt.Printf("地%v个参数是float64, 值是%v\n", i, item)
		case int, int32, int64:
			fmt.Printf("地%v个参数是int, 值是%v\n", i, item)
		case string:
			fmt.Printf("地%v个参数是string, 值是%v\n", i, item)
		default:
			fmt.Printf("地%v个参数不确定, 值是%v\n", i, item)
		}

	}
}
func main() {

	a := 1
	b := "hello"
	c := 1.1
	var d float64 = 1.22
	e := true
	jumpType(a, b, c, d, e)

	//var heroes Hero.HeroSlice
	//
	//for i := 0; i < 10; i++ {
	//	hero := Hero.Hero{Age: rand.Intn(100), Name: fmt.Sprintf("英雄～%d", rand.Intn(100))}
	//	heroes = append(heroes, hero)
	//}
	//fmt.Println(heroes)
	//sort.Sort(heroes)
	//fmt.Println(heroes)

	//monkey := Monkey.LittleMonkey{Monkey: Monkey.Monkey{Name: "孙悟空"}}
	//monkey.Climbing()
	//monkey.Flying()
	//monkey.Swimming()
	//
	//var usbs [3]_interface.IUsb
	//usbs[0] = Camera.Camera{
	//	Name: "nikon",
	//}
	//usbs[1] = Camera.Camera{
	//	Name: "Canon",
	//}
	//usbs[2] = Mobile.Mobile{
	//	Name:  "小米",
	//	Color: "black",
	//}
	//fmt.Println(usbs)
	//
	//var a interface{}
	//
	//var point Point = Point{"1", "2"}
	//a = point
	//var b Point
	//b = a.(Point) // 类型断言 assert
	//
	//fmt.Println(b)
	//var computer Computer.Computer
	//for _, usb := range usbs {
	//	computer.Working(usb)
	//	if usb, ok := usb.(Camera.Camera); ok { // 断言
	//		usb.Photo()
	//	}
	//
	//	if usb, ok := usb.(Mobile.Mobile); ok { // 断言
	//		usb.Call()
	//	}
	//}

	// 判断类型

}

// 接口体现多态的两种方式
// 多态参数
// 多态数组
