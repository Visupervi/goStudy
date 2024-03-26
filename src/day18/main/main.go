package main

import (
	"day18/account"
)

//type Animal struct {
//	Name   string `json:"name"`
//	Age    int    `json:"age"`
//	Gender string `json:"gender"`
//	Voice  string `json:"voice"`
//	Color  string `json:"color"`
//	Color  string `json:"color"`
//}
//
//
//
//func (a Animal) getFiles() {
//	fmt.Printf("姓名=%v\n年龄=%v\n性别=%v\n叫声=%v\n颜色=%v\n", a.Name, a.Age, a.Gender, a.Voice, a.Color)
//}
//
//func (a Animal) getSum(n1 int, n2 int) int {
//	return n1 + n2
//}
//
//func (c *Animal) getSum1() float64 {
//	return 3.14 * c.Radius * c.Radius
//}
//
//// 如果实现了*Animal 的String方法， 就会自动执行， 如果没实现就会输出地址
//func (animal *Animal) String() string {
//	str := fmt.Sprintf("name=[%v] age=[%v]", animal.Name, animal.Age)
//	return str
//}

//type Student struct {
//	Name string
//	Gender string
//	Age int
//	Id int
//	Score float64
//}
//
//func (student *Student) get() string  {
//	str := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
//		student.Name, student.Gender, student.Age, student.Id, student.Score)
//
//	return str
//}
//type Person struct {
//	Name    string `json:"name"`
//	Age     int    `json:"age"`
//	Address string `json:"address"`
//}
//
//
//type Box struct {
//	width float64
//	height float64
//	len float64
//}
//
//func (b *Box) getVolumn() float64  {
//	return b.height * b.width * b.len
//}
//
//type Visitor struct {
//	Age int
//	Name string
//}
//
//func (v *Visitor) buyTollet()   {
//	if v.Age > 90 || v.Age < 8 {
//		fmt.Println("太危险了，你就不要玩了\n")
//		return
//	}
//
//	if v.Age < 18 {
//		fmt.Printf("姓名 %v 年龄 %v 可以免费\n", v.Name, v.Age)
//	}
//
//	if v.Age >= 18 {
//		fmt.Printf("姓名 %v 年龄 %v 收费100\n", v.Name, v.Age)
//	}
//
//}
func main() {
	//dog := Animal{"Tom", 3, "female", "汪汪汪", "black", 12.2}
	//dog.getFiles()
	//cat := Animal{"Jerry", 2, "male", "喵喵喵", "white", 11.1}
	//cat.getFiles()
	//res := cat.getSum(10, 20)
	//
	//fmt.Println("res=", res)
	//
	//res1 := cat.getSum1()
	//
	//fmt.Println("res1=", res1)

	//fmt.Println(&dog)

	//stu := Student{"铁锤","女", 18,112344444,61.5}
	//res := stu.get()
	//
	//fmt.Println(res)
	//
	//box := Box{
	//	11.1,
	//	21.22,
	//	13.33,
	//}
	//
	//fmt.Printf("体积=%.2f", box.getVolumn())

	//var v Visitor
	//
	//for  {
	//	fmt.Println("请输入名字")
	//	fmt.Scanln(&v.Name)
	//
	//	if v.Name == "n" {
	//		fmt.Println("正在退出。。。")
	//		break
	//	}
	//	fmt.Println("请输入年龄")
	//
	//	fmt.Scanln(&v.Age)
	//
	//	v.buyTollet()
	//}
	//stu := utils.FactoryStudent("李四",20, 88.88)
	//fmt.Println(*stu)
	//fmt.Println(stu.GetScore())

	account := account.Account{"浦发", "123", 10}

	account.LookMoney("123")
	account.SaveMoney(10, "123")
	account.GetMoney(10, "123")
}

// 面向对象

// 定义结构体

// 定义字段

// 定义方法

// 构造函数

// 工厂模式
