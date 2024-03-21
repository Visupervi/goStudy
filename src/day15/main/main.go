package main

import (
	"fmt"
	"sort"
)

// 二维数组

// 结构体
type student struct {
	name    string
	age     int
	address string
}

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

	//students := make(map[string]map[string]string)
	//
	//students["stu1"] = make(map[string]string)
	//students["stu1"]["name"] = "Tom"
	//students["stu1"]["gender"] = "males"
	//students["stu1"]["age"] = "20"
	//
	//students["stu2"] = make(map[string]string)
	//students["stu2"]["name"] = "Mary"
	//students["stu2"]["gender"] = "females"
	//students["stu2"]["age"] = "18"
	//
	//fmt.Println(students)

	//for key, value := range students {
	//	fmt.Println("\n",key)
	//	fmt.Println("\n",value)
	//	for k1, v1 := range value {
	//		//fmt.Println("\n",k1)
	//		//fmt.Println("\n",v1)
	//		fmt.Printf("key=%v, value=%v\n",k1,v1)
	//	}
	//}

	// map 切片

	//var monsters []map[string]string
	//monsters = make([]map[string]string,2)

	// 第二种
	//monsters := map[string]string{
	//	"name": "Tom",
	//	"age": "20",
	//}

	map1 := make(map[int]int)
	map1[1] = 1
	map1[2] = 2
	map1[3] = 3
	map1[4] = 4
	map1[5] = 5
	map1[6] = 6
	fmt.Println(map1)
	var keys []int
	for k, _ := range map1 {
		//fmt.Printf("k=%v,v=%v\n", k, v)
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("k=%v, value=%v\n", key, map1[key])
	}
	students := make(map[string]student)

	stu1 := student{
		name:    "张三",
		age:     18,
		address: "上海市浦东新区",
	}

	stu2 := student{
		name:    "李四",
		age:     18,
		address: "上海市浦东新区",
	}

	students["stu001"] = stu1
	students["stu002"] = stu2

	fmt.Println(students)

	persons := make(map[string]map[string]string)

	//info := make(map[string]string)
	//
	//persons["张三"] = info

	updatePassWord(persons, "李四")
	updatePassWord(persons, "张三")

	fmt.Println(persons)
}

// map 数据结构
// 基本介绍用法，:var 变量名 map[keyType]valueType
// 基本使用

var a map[string]string // 声明是不分配内存的，需要make，分配内存后才可以赋值和使用

// map的使用细节

// 1. 是引用类型 map的value也可以是结构体
// 2. map到达容量后，会自动扩容

func updatePassWord(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["pwd"] = "xxxxxx"
		users[name]["age"] = "20"
		users[name]["gender"] = "female"
	}
}
