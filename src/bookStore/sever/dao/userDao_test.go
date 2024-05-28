package dao

import (
	"bookStore/sever/db"
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
	//t.Run("插入Session测试", testAddSession)
	//t.Run("删除Session测试", testDeleteSession)
	//t.Run("查找Session测试", testGetSessionById)
	//t.Run("购物车测试", testCartAdd)
	//t.Run("购物车项测试", testAddCartsItem)
	//t.Run("测试获取购物车项", testGetCartItem)
	t.Run("购物车获取测试", testGetCartItems)
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
	sess := "45d1b156-e635-4b5c-aaa5-fcf6d56fd6d8"

	DeleteSession(sess)
}

func testGetSessionById(t *testing.T) {
	sess := "8d8f47dd-0ac0-49b1-8428-9ccee45bdcbd"

	GetSessionById(sess)
}

// 5 | 苏东坡传                  | 林语堂              |    19.00 |     100 |     100 | static/img/苏东坡传.jpg
func testCartAdd(t *testing.T) {

	var items []*model.CartItem

	book := &model.Book{
		ID:    5,
		Price: 19.00,
	}

	book1 := &model.Book{
		ID:    6,
		Price: 29.00,
	}
	uuidValue := uuid.New()
	item := &model.CartItem{
		CartId: uuidValue.String(),
		//Amount: 12,
		Count: 12,
		Book:  book,
	}
	items = append(items, item)
	item1 := &model.CartItem{
		CartId: uuidValue.String(),
		//Amount: 12,
		Count: 10,
		Book:  book1,
	}
	items = append(items, item1)

	//item := append()
	cart := &model.Cart{
		CartId: uuidValue.String(),
		//TotalAmount: 128.11,
		//TotalCount: 22,
		UserId: 1,
		Items:  items,
	}
	AddCart(cart)
}

func testAddCartsItem(t *testing.T) {
	db, _ := db.ConnectDB()
	//if error != nil {
	//	//return nil
	//}
	defer db.Close()
	//uuidValue := uuid.New()
	book := &model.Book{
		ID:    5,
		Price: 19.00,
	}
	item := &model.CartItem{
		CartId: "5b4eaa8f-2661-464c-9b55-af5ab3ce70cf",
		Count:  10,
		//Amount: 100,
		Book: book,
	}

	AddCartsItem(item)
}

func testGetCartItem(t *testing.T) {
	item, _ := GetCartItem(5)

	fmt.Println("item", item)
}

func testGetCartItems(t *testing.T) {
	items, _ := GetCartItems("72a40d8a-a9a0-43b0-9db8-4d8d97d77343")

	fmt.Println("items", items)

}
