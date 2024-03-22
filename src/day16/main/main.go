package main

import (
	"encoding/json"
	"fmt"
)

// 结构体

// 在内存中如何存在的, 结构体是值类型

// 属性

// 面向对象
type student struct {
	name    string ""
	age     int
	address string ""
}

type Cat struct {
	Name  string
	Age   int
	Color string
}

type Person struct {
	Name string
	Age  int
}

type Point struct {
	x int
	y int
}

type Circle struct {
	left, right Point
}

type A struct {
	Num int
}

type B struct {
	Num int
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Stu Student

func (s Student) test() {
	fmt.Println("调用了====", s.Name)
}

func main() {
	//cat1 := Cat{
	//	Name:  "小花",
	//	Age:   3,
	//	Color: "black",
	//}

	p := Person{"Mary", 20}
	fmt.Println(p)

	//fmt.Println(cat1.Name)

	var p1 *Person = new(Person)

	(*p1).Name = "李四"

	p1.Age = 19

	// 内存的分配机制
	fmt.Println(*p1)
	var person *Person = &Person{"张三", 20}

	fmt.Println(*person)

	var p2 Person

	p2.Age = 10
	p2.Name = "令狐冲"
	var p3 *Person = &p2

	fmt.Println((*p3).Age)
	fmt.Println(p3.Age)

	p3.Name = "风清扬"

	fmt.Printf("p2.name=%v, p3.name=%v\n", p2.Name, p3.Name)
	fmt.Printf("p2.name=%v, p3.name=%v\n", (*p3).Name, p2.Name)
	circle := Circle{Point{1, 2}, Point{3, 4}}

	fmt.Println(circle)

	// 结构体是用户单独定义的类型，和其他类型转换时需要有完全相同的字段（名字，个数和类型）

	var a A
	var b B

	a = A(b)

	fmt.Println(a)

	// 结构体进行type重新定义（相当于重新命名）golang认为是新的数据类型，但是相互间可以强转
	s1 := Student{"小明", 20}

	jsonStr, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))

	//s1.test()
	//var s2 Stu

	//s1 = Student(s2)

	// struct 每个字段可以写一个tag

	// 方法和函数

	// 方法是指作用在制定的数据类型上，自定义的数据类型都可以有方法
	s1.test()

}
