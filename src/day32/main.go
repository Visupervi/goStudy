package main

import (
	"day32/entity"
	"fmt"
	"reflect"
)

// 反射学习
// reflect
// 基本介绍
// 反射可以在运行时动态的获取变量的各种信息，比如变量的类型、类别
// 如果时结构体变量，还可以获取到结构体本身的信息（包括结构体字段、方法）
// 通过反射修改变量的值，可以调用关联的方法
// 使用反射 需要import("reflect")
// reflect包实现了运行时反射，允许程序操作任意类型的对象。
// 典型用法是用静态类型interface{}保存一个值，
// 通过调用TypeOf获取其动态类型信息，该函数返回一个Type类型值。
// 调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。
// Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。
// 变量、interface、和reflect.Value是可以相互转换的

func main() {
	stu := entity.Student{
		Name:  "张三",
		Age:   18,
		Grade: 99.9,
	}

	test(stu)

	//const r = 3.14
}

func test(i interface{}) {
	//	将interface转成reflect.Value
	rVal := reflect.ValueOf(i)

	fmt.Printf("type=%T", rVal)
	fmt.Println(rVal.Kind())
	//	将reflect.Value转成interface
	iVal := rVal.Interface()
	fmt.Printf("type111=%T\n", iVal)

	rr := iVal.(entity.Student)
	fmt.Printf("type122=%v", rr.Name)

	//rr = rVal.(entity.Student) // 转回最开始的类型

}
