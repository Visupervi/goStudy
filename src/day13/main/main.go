package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//arrTest1()

	//arr := [...]int{1, 2, 4}

	//test01(arr)
	//
	//fmt.Println(arr)
	//test02(&arr)
	//fmt.Println(arr)
	//test1()
	max := findMax()

	fmt.Println(max)

	//sum()
	//randomNum()
	//sliceTest()
	sliceTTest()
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

//
func test01(arr [3]int) {
	arr[0] = 99
}

// 通过指针来修改数组
func test02(arr *[3]int) {
	(*arr)[0] = 100
}

func test1() {
	var arr [26]byte
	arr[0] = 'A'
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + 1
	}
	for _, v := range arr {
		fmt.Printf("%c ", v)
	}
}

func findMax() int {
	numArr := [...]int{1, -1, 9, 90, 11}

	max := numArr[0]

	for _, v := range numArr {
		if v > max {
			max = v
		}
	}
	return max
}

func sum() {
	sum := 0
	numArr := [...]int{1, -1, 9, 90, 12}
	for _, v := range numArr {
		sum += v
	}

	fmt.Println("sum=", sum)
	fmt.Println("avg=", float64(sum)/float64(len(numArr))) // /
}

// 随机生成五个数， 反转打印

func randomNum() {
	var randoms [6]int
	lenth := len(randoms)
	rand.Seed(time.Now().UnixNano()) // 设置seed
	for i := 0; i < 6; i++ {
		n := rand.Intn(100) + 1
		randoms[i] = n
	}
	//types.Array{}
	fmt.Println(randoms)

	// 反转数组

	for i := 0; i < lenth/2; i++ {
		var temp = randoms[i]
		randoms[i] = randoms[lenth-i-1]
		randoms[lenth-i-1] = temp
	}

	fmt.Println(randoms)

}

// 切片

// 切片是数组的一个引用， 因此切片一个引用类型， 在进行传递的时候，遵循引用传递的机制
// 切片的使用和数组类似，遍历切片、访问切片的元素和🍂切片的长度都一样
// 切片的长度是可以变化的，因此切片是一个可以动态的数组
// 切片的语法 var 变量名 []类型------》var a []int

// [1:3] 就是类似 Javascript中的slice函数，返回一个数组

func sliceTest() {
	//var sliceDemo []int
	var intArr [5]int = [5]int{1, 2, 3, 4, 5}

	sliceArr := intArr[1:3] //

	//fmt.Println(sliceArr)
	//fmt.Println(sliceDemo)
	fmt.Printf("%p\n", &intArr[1])
	fmt.Printf("%v\n", sliceArr)
	fmt.Printf("%v\n", sliceArr[0])
	fmt.Printf("%v", sliceArr[1])
	//fmt.Printf("%v", sliceArr[2])

}

// 切片在内存中形式， 在内存中是有三部分组成，一个是指向的数据的地址，一个是长度，一个是容量

// 切片的使用

// 第一种

// 第二种

func makeSlice() {
	var slice []int = make([]int, 4, 10)
	slice[0] = 10
}

// 通过make创建的切片可以制定切片的大小和容量
// 如果没有给切片的各个元素赋值，就会使用默认值
// 通过make创建的切片的对应的数组，对外不可见

//第三种 定义一个切片， 直接制定具体数组

var sliceStr []string = []string{"tom", "jack", "jerry"}

//

func sliceTTest() {
	arr := [...]int{10, 20, 30, 40, 50}
	slice := arr[:]

	for _, v := range slice {
		fmt.Println(v)
	}
}
