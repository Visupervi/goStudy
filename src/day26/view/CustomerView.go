package view

import "fmt"

type CustomerView struct {
	Key  string
	Loop bool
}

func (cv *CustomerView) GetMainMenu() {
	for {
		fmt.Println("---------------------------客户信息管理软件---------------------------")
		fmt.Println("                           1 添 加 用 户")
		fmt.Println("                           2 修 改 用 户")
		fmt.Println("                           3 删 除 用 户")
		fmt.Println("                           4 客 户 列 表")
		fmt.Println("                           5 退      出")
		fmt.Print("请选择<1-5>:")
		fmt.Scanln(&cv.Key)
		switch cv.Key {
		case "1":
			fmt.Println("添 加 用 户")
		case "2":
			fmt.Println("修 改 用 户")
		case "3":
			fmt.Println("删 除 用 户")
		case "4":
			fmt.Println("客 户 列 表")
		case "5":
			cv.Loop = false
		default:
			fmt.Println("请输入正确菜单")
		}

		if !cv.Loop {
			fmt.Println("退出")
			break
		}
	}
}
