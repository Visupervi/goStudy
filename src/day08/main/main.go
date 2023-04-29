package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 还需要设置一个种子，如果没有设置的话，永远是一个数
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
		//fmt.Println(n)
	}
label2:
	for i := 0; i < 5; i++ {
		//label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label2
			}
			fmt.Println("j=", j)
		}
	}

	var sum int = 0

	for i := 0; i < 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("i=", i)
			fmt.Println("sum=", sum)
			break
		}
	}
	var name string = ""
	var pwd string = ""
	var remain int = 3
	for i := 0; i < 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&pwd)
		if name == "张三" && pwd == "888" {
			fmt.Println("恭喜你登陆成功")
			break
		} else {
			remain--
			fmt.Printf("你还有%v次机会, 请珍惜", remain)
		}
		if remain == 0 {
			fmt.Println("登陆失败")
		}
	}
	//fmt.Println(count)
	// break 会默认跳出最近的for循环
	// break 配合label使用，制定跳转到哪里
}
