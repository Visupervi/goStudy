package main

import "fmt"

// 再调用一个函数时，会给该函数分配一个新的空间，编辑器会自身处里让这个地址和其他空间区分开来
// 在每个函数对应的栈中，数据空间是空间的，不会混淆
// 当一个函数执行完毕后，程序会销毁这个函数对应的空间
// go 可以返回多个值
// 可以使用_占位，忽略

func main() {
	_, res2 := getNumbs(123, 12)
	//fmt.Println("res1=",)
	fmt.Println("res2=", res2)
}

func getNumbs(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
