package model

import (
	"fmt"
	"goWeb01/db"
	"testing"
)

func TestMain(m *testing.M) {
	db.InitDB()
	m.Run()
}

func TestUser(t *testing.T) {
	//t.Run("测试添加", testUser_AddUser)
	t.Run("查找添加", testUser_GetUserById)
	t.Run("查找所有", testUser_GetUsers)
}

func testUser_AddUser(t *testing.T) {
	user := &User{
		UserName: "admin",
		Password: "123456",
		Email:    "123456@qq.com",
	}
	err := user.AddUser(user)

	fmt.Println("err=", err)
}

func testUser_AddUser2(t *testing.T) {
	user := &User{
		UserName: "admin1",
		Password: "123456",
		Email:    "1234577@qq.com",
	}
	user.AddUser1(user)
}

func testUser_GetUserById(t *testing.T) {

	user := &User{}

	user, err := user.GetUserById(1)

	if err != nil {
		fmt.Println("查找出错", err)
		return
	}

	fmt.Println("user", user)
}

func testUser_GetUsers(t *testing.T) {
	user := &User{}

	users, err := user.GetUsers()
	if err != nil {
		fmt.Println("查找出错", err)
		return
	}
	for _, u := range users {
		fmt.Println("user=", u)
	}
	//fmt.Println("user", users)
}
