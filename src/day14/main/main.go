package main

import "fmt"

// 切片二
// 数组是同一种数据类型元素的集合，数组在定义时需要指定长度和元素类型。

// 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

func main() {
	//var slice []int = []int {10,20,30}
	//slice1 := append(slice, 100)
	//fmt.Println(slice1)
	//fmt.Println(slice)
	//var slice2 = make([]int, 10)
	//
	//copy(slice2, slice) // 返回的是slice的长度
	//
	//fmt.Println(slice2)

	sliceTest()
}

//append
// copy 要求参数都是切片
func sliceTest() {
	//var slice []int

	var arr []int = []int{1, 2, 3, 4, 5}
	test(arr)

	fmt.Println(arr)
	//
	//slice = arr[:]
	//var slice2 = slice
	//slice2[0] = 20
	//fmt.Println(slice)
	//fmt.Println(slice2)
	//fmt.Println(arr)
}

func test(slice []int) {
	slice[0] = 100
	fmt.Println("slice=", slice)
}

// string 和 slice的关系

// string的底层还一个byte数组
