package main

import (
	"fmt"
)

func main() {

	//arrTest1()

	arr := [...]int{1, 2, 4}

	test01(arr)

	fmt.Println(arr)
	test02(&arr)
	fmt.Println(arr)
}

// 数组的定义

func arrTest() {
	arr := []float64{3, 5, 1, 3.4, 2, 50}
	var sum float64 = 0.0
	for _, i2 := range arr {
		//println("i = ", i)
		//println("i2 = ", i2)
		sum += i2
	}

	fmt.Println("总=", sum)
	fmt.Println("平均", fmt.Sprintf("%.2f", sum/float64(len(arr))))

}

// 数组的定义 和内存分配
func arrTest1() {
	var intArr [4]int
	intArr[0] = 10
	intArr[1] = 30
	intArr[2] = 20
	fmt.Printf("intArr的地址=%p,intArr[0]的地址=%p, intArr[1]的地址=%p，intArr[2]的地址=%p\n", &intArr, &intArr[0], &intArr[1], &intArr[2])

	// 四种初始化数组的方法

	var numArr [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr", numArr)
	var numArr1 = [3]int{5, 6, 7}
	fmt.Println("numArr1", numArr1)
	var numArr2 = [...]int{11, 12, 13}
	fmt.Println("numArr2", numArr2)
	var numArr3 = [...]int{0: 12121, 1: 90, 2: 33}
	fmt.Println("numArr3", numArr3)

	// 常规遍历

	// for range
}

// 数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
// var arr[]int 这时arr就是一个slice切片
// 数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
// 数组创建后，如果没有赋值，有默认值 数值类型默认0， 字符串类型 ""， bool数组默认false
// 使用数组步骤，1。声明数组并开辟空间 2给各个元素赋值 3使用数组
// 数组下标必须在合理范围内，否则会报panic
// go的数组属于值类型， 在默认的情况下是值传递，因此会进行值拷贝，数组间不会相互影响
// 如果想再其他函数中去修改原来的数组，可以使用引用传递（指针的方式）。

func test01(arr [3]int) {
	arr[0] = 99
}

// 通过指针来修改数组
func test02(arr *[3]int) {
	(*arr)[0] = 100
}
