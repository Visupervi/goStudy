package main

import (
	"fmt"
)

/**
家庭收支记账软件
收入
支出
输出收入支出表

1 收支明细
2 登记收入
3 登记支出
4 退出

*/

func main() {
	var key string
	var flag bool = true
	balance := 10000.0
	note := ""
	money := 0.00
	details := "收支\t账户金额\t收支金额\t说明"
	for {
		fmt.Println("\n---------------------------收支软件已启动---------------------------")
		fmt.Println("                             输入1显示收支明细")
		fmt.Println("                             输入2登记收入")
		fmt.Println("                             输入3登记支出")
		fmt.Println("                             输入4退出软件")
		fmt.Println("-----------------------------收支软件已启动---------------------------")
		fmt.Println("请输入<1-4>：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("----------------------收支明细----------------------")
			fmt.Println(details)
		case "2":
			fmt.Println("请输入收入")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("请输入说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("输入支出")
			//fmt.Println("请输入收入")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("请输入说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
		case "4":
			fmt.Println("你确定要退出么？y/n")
			choice := ""

			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}

				fmt.Println("输入有误")
			}

			if choice == "y" {
				flag = false
			}
		}
		if !flag {
			fmt.Println("退出收支软件")
			break
		}
	}
}
