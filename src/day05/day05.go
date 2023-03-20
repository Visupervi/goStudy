package main

import "fmt"

var a string
var (
	n1 = 90
	n2 = "Rose"
	n3 = "Jerry"
)

func main() {

	// 指针变量存储的是地址，这个地址指向的空间才是值
	//i := 10
	//fmt.Println("i的地址=", &i)
	//
	//var point *int = &i
	//*point = 100
	//fmt.Println("i的值=", i)
	//fmt.Println("point指向的值=", *point)

	// var num int = 1
	// var j = 999
	// var ptr *int = &j
	// 引用类型
	// 指针、slice切片、map、管道chan、interface等都是引用类型
	// 值类型与引用类型的区别
	// 值类型：存放在栈中
	// 引用类型：存放在堆中
	// go编译会有逃逸分析，基本类型也可能存在堆中，引用类型也可能在栈中
	// go中根据变量的存活时间，具体分析存放在内存的哪一个地方

	// 标识符概念

	// 命名规则
	// 由26个字母大小写，0-9，_组成
	// 数字不可以开头
	// 严格区分大小写
	// 标识符不能包含空格
	// 下划线'_'本身在go中是一个特殊字符，称之为空标识符。可以代表任何其他的标识符
	// 但是它对应的值会被忽略，所以仅能够被作为占位符使用，不能作为标识符使用
	// 不能以系统的保留关键字作为标识符
	// 如果变量名、函数名、常量名首字母大写，则可以被其他的包访问，如果首字母小写则只能在本包
	// 中使用（大小写区分私有化）
	// 算数运算符

	// a % b = a- a / b * b
	// ++和--的使用
	// 只能独立使用
	//++ 或者--完成后再使用
	// 只有后++和后--
	//

	var name string
	var age byte
	var sal float32
	var isPass bool
	var gender string

	//fmt.Println("请输入姓名")
	//fmt.Scanln(&name)
	//
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&sal)
	//
	//fmt.Println("是否通过考试")
	//fmt.Scanln(&isPass)

	//fmt.Printf("姓名：%v \n 年龄： %v \n 薪水是：%v \n 是否通过考试：%v \n", name, age, sal, isPass)

	fmt.Println("请输入姓名，性别，年龄，薪水，是否通过考试")
	fmt.Scanf("%s %s %d %f %t", &name, &gender, &age, &sal, &isPass)
	fmt.Printf("姓名：%v \n 性别：%v \n 年龄： %v \n 薪水是：%v \n 是否通过考试：%v \n", name, gender, age, sal, isPass)

}
