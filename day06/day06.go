package main

import (
	"fmt"
)

func main() {
	//  主要学些go的语法

	//	顺序控制
	//	分支控制
	//	if else
	//	if 条件表达式 {}
	//	if else if else
	//	循环控制

	var age byte
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你已经大于18，请为你的行为负责")
	} else {
		fmt.Println("请在监护人的陪同下观看")
	}

}
