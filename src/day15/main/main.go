package main

import "fmt"

// 二维数组

func main() {
	//var arr [2][3]int = [2][3]int {{12,3,4},{33,44,55}}
	//
	//fmt.Printf("arr[0]=%p\n", &arr[0])
	//fmt.Printf("arr[0][0]=%p\n", &arr[0][0])
	//
	//fmt.Printf("arr[1]=%p\n", &arr[1])
	//fmt.Printf("arr[1][0]=%p", &arr[1][0])
	//
	//fmt.Println(arr)
	//fmt.Println(arr[1][1])
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr[i]); j++ {
	//		fmt.Printf("%v\t", arr[i][j])
	//	}
	//	fmt.Println()
	//}
	var a map[string]string
	// 使用map前先make
	a = make(map[string]string)
	a["no1"] = "张三"
	a["no2"] = "李四"
	a["no1"] = "王五"
	a["no4"] = "赵六"
	a["tes"] = "周七"
	fmt.Println(a)

	cities := make(map[string]string)

	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "重庆"
	cities["no4"] = "天津"

	fmt.Println(cities)

	heroes := map[string]string{
		"hero1": "熏悟空",
		"hero2": "花果山",
	}

	fmt.Println(heroes)

	students := make(map[string]map[string]string)

	students["stu1"] = make(map[string]string)
	students["stu1"]["name"] = "Tom"
	students["stu1"]["gender"] = "males"
	students["stu1"]["age"] = "20"

	students["stu2"] = make(map[string]string)
	students["stu2"]["name"] = "Mary"
	students["stu2"]["gender"] = "females"
	students["stu2"]["age"] = "18"

	fmt.Println(students)
}

// map 数据结构
// 基本介绍用法，:var 变量名 map[keyType]valueType
// 基本使用

var a map[string]string // 声明是不分配内存的，需要make，分配内存后才可以赋值和使用
