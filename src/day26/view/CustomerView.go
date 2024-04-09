package view

import (
	"day26/model"
	"day26/service"
	"fmt"
)

type CustomerView struct {
	Key             string
	Loop            bool
	CustomerService *service.CustomerService
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
			cv.Add()
		case "2":
			fmt.Println("修 改 用 户")
		case "3":
			cv.Delete()
		case "4":
			cv.GetCustomerList()
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

func (cv *CustomerView) GetCustomerList() {
	slice := cv.CustomerService.List()

	fmt.Println("------------------------------------客户列表-----------------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")

	for _, customer := range slice {
		fmt.Println(customer.GetInfo())
	}
	fmt.Println("----------------------------------客户列表完成----------------------------------")
}

func (cv *CustomerView) Add() {
	fmt.Println("-------------------------------------添加客户-------------------------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话：")
	Tel := ""
	fmt.Scanln(&Tel)

	fmt.Println("邮件地址：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(name, age, Tel, gender, email)
	if cv.CustomerService.InsertCustomer(customer) {
		fmt.Println("------------------------------------添加完成------------------------------------")

	} else {
		fmt.Println("------------------------------------添加失败------------------------------------")

	}

}

func (cv *CustomerView) Delete() {
	fmt.Println("-------------------------------------删除客户-------------------------------------")

	fmt.Println("请输入用户编号：")
	id := 0
	fmt.Scanln(&id)

	if id == -1 {
		return
	} else {
		fmt.Println("确认是否删除 y/n：")
		choice := ""
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			if cv.CustomerService.DeleteCustomer(id) {
				fmt.Println("------------------------------------删除完成------------------------------------")

			} else {
				fmt.Println("------------------------------------添加失败------------------------------------")
			}
		}
	}

}
