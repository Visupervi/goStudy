package main

import "fmt"

type Animal struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Voice  string `json:"voice"`
	Color  string `json:"color"`
	Radius float64
}

func (a Animal) getFiles() {
	fmt.Printf("姓名=%v\n年龄=%v\n性别=%v\n叫声=%v\n颜色=%v\n", a.Name, a.Age, a.Gender, a.Voice, a.Color)
}

func (a Animal) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func (c *Animal) getSum1() float64 {
	return 3.14 * c.Radius * c.Radius
}

// 如果实现了*Animal 的String方法， 就会自动执行， 如果没实现就会输出地址
func (animal *Animal) String() string {
	str := fmt.Sprintf("name=[%v] age=[%v]", animal.Name, animal.Age)
	return str
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	dog := Animal{"Tom", 3, "female", "汪汪汪", "black", 12.2}
	dog.getFiles()
	cat := Animal{"Jerry", 2, "male", "喵喵喵", "white", 11.1}
	cat.getFiles()
	res := cat.getSum(10, 20)

	fmt.Println("res=", res)

	res1 := cat.getSum1()

	fmt.Println("res1=", res1)

	fmt.Println(&dog)
}

// 方法和函数的区别

// 调用方式不同
// 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
// 对于方法，接收者为值类型时，可以直接使用指针类型的变量调用方法，反过来同样可以
