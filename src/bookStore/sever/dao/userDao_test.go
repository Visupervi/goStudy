package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	db.InitDB()
	m.Run()
}
func TestUser(t *testing.T) {
	fmt.Println()
	t.Run("添加用户", testAddUser)
	t.Run("检查用户", testGetUsers)
	t.Run("根据用户名检查用户", testGetUserByName)
}

func testAddUser(t *testing.T) {
	u := &model.User{
		UserName: "admin",
		Password: "123456",
		Email:    "visupervi@outlook.com",
	}
	AddUser(u)
}

func testGetUserByName(t *testing.T) {
	GetUserByName("admin")
}

func testGetUsers(t *testing.T) {
	GetUsers("admin", "123456")
}
