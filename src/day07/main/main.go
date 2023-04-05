package main

import (
	"fmt"
)

/**
 * @Author visupervi
 * @Description
 * @Date 15:49 2023/4/5
 * @Param
 * @return
 */

// for循环
//
func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Hello Kitty")
	}
	//	for循环的第二种写法

	j := 0
	for j < 10 {
		fmt.Println("Hello Wold")
		j++
	}

	//	for循环的第三种写法
	//等价于 for ;;{} 等价于无限循环
	k := 0
	for {
		if k < 10 {
			fmt.Println("Hello Jack")
		} else {
			break
		}
		k++
	}

	//	遍历字符串

	str := "hello wold!北京上海"
	str1 := []rune(str)
	// 若果字符串含有中文，遍历会出现乱码
	// 传统的遍历是按照字节遍历的，中文正好UTF8是三个字节。
	// 解决方案就是将数组转换成切片 str 转成[]rune切片
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c \n", str1[i])
	}

	// 是按照字符遍历的，所以遇到中文也不会乱码
	for i, i2 := range str {
		fmt.Printf("index = %d, val= %c \n", i, i2)
	}

	// 打印1~100间9的倍数的个数以及他们的和

	num := 0

	sum := 0

	for i := 0; i < 100; i++ {
		if i%9 == 0 {
			num++
			sum += i
		}
	}

	fmt.Printf("是9的倍数个数=%d, 和=%d", num, sum)

	// go 没有while 和 do while
}
