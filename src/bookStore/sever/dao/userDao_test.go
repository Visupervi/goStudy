package dao

import (
	"bookStore/sever/model"
	"fmt"
	"github.com/google/uuid"
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
	//t.Run("分页查询图书", testGetBooksByPage)
	//t.Run("插入图书测试", testInsertBook)
	t.Run("插入Session测试", testAddSession)
	//t.Run("删除Session测试", testDeleteSession)
	t.Run("查找Session测试", testGetSessionById)
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

func testInsertBook(t *testing.T) {
	book := &model.Book{
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   20,
		Sales:   10000,
		Stock:   10000,
		ImgPath: "static/img/西游记.png",
	}

	InsertBook(book)
}

func testAddSession(t *testing.T) {
	uuidValue := uuid.New()
	sess := &model.Session{
		SessionId: uuidValue.String(),
		UserName:  "admin",
		UserId:    1,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	sess := "e851e923-0820-41d4-bc46-dd5a0a9cf3e0"

	DeleteSession(sess)
}

func testGetSessionById(t *testing.T) {
	sess := "8d8f47dd-0ac0-49b1-8428-9ccee45bdcbd"

	GetSessionById(sess)
}
