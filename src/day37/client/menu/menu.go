package menu

import (
	"day37/client/model"
	"day37/client/userProcess"
	"fmt"
)

func InitMenu() {
	var key int
	var userId int
	var pwd string
	var uName string
	loop := true
	for loop {
		fmt.Println("------------------------------欢迎登陆多人聊天系统------------------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）:")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &pwd)
			//ul := login.UserLogin{
			//	UserId: userId,
			//	Pwd: pwd,
			//}
			up := userProcess.UserProcess{}
			err := up.Login(userId, pwd)

			if err != nil {
				fmt.Println("********登陆失败********")
			}

			loop = false

		case 2:
			fmt.Println("********开始注册用户********")

			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &pwd)
			fmt.Println("请输入用户名")
			fmt.Scanf("%s\n", &uName)
			user := model.User{
				UserId:   userId,
				Pwd:      pwd,
				UserName: uName,
			}

			up := userProcess.UserProcess{}

			up.RegistryProcess(user)
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误")
		}

	}
	//
	//if key == 1 {
	//	fmt.Println("请输入用户ID")
	//	fmt.Scanf("%d\n", &userId)
	//	fmt.Println("请输入用户密码")
	//	fmt.Scanf("%s\n", &pwd)
	//	ul := login.UserLogin{
	//		UserId: userId,
	//		Pwd: pwd,
	//	}
	//	err := ul.Login(userId, pwd)
	//
	//	if err != nil {
	//		fmt.Println("********登陆失败********")
	//	} else {
	//		fmt.Println("********登陆成功********")
	//	}
	//
	//} else if key == 2 {
	//	fmt.Println("请注册用户")
	//}
}
