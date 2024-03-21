package main

import "fmt"

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

func main() {
	cat1 := Cat{
		Name:  "小花",
		Age:   3,
		Color: "black",
	}
	fmt.Println(cat1.Name)
}
