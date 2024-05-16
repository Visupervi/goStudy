package dao

import (
	"bookStore/sever/model"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	//db.InitDB()
	m.Run()
}
func TestUser(t *testing.T) {
	fmt.Println()
	//t.Run("添加用户", testAddUser)
	//t.Run("检查用户", testGetUsers)
	//t.Run("根据用户名检查用户", testGetUserByName)
	//t.Run("查询所有图书", testGetBooks)
	t.Run("分页查询图书", testGetBooksByPage)
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

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()

	for k, v := range books {
		//fmt.Println("k=", k)
		fmt.Printf("k=%v,v=%v\n", k, v)
	}

	//fmt.Println("books", books)
}

func testGetBooksByPage(t *testing.T) {
	//db, _ := db.ConnectDB()

	//if error != nil {
	//	return nil, error
	//}
	//defer db.Close()
	res, _ := GetBooksByPage(1, 10)

	fmt.Println("res", res)
}
