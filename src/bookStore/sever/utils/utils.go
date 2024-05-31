package utils

import (
	"bookStore/sever/model"
	"bookStore/sever/service"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func SetCorsHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域
	//w.Header().Set("Content-Length", "65536") // 允许所有域
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
}

// String2Int string转int
func String2Int(str string) int {

	int, err := strconv.Atoi(str)

	if err != nil {
		return 0
	}
	return int
}

// String2Int64 string转int64
func String2Int64(str string) int64 {

	int64, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return 0
	}
	return int64
}

// String2Float64 string 转 float64
func String2Float64(str string) float64 {
	float, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0
	}
	return float
}

func UpdateCartAddItem(params map[string]string, cart *model.Cart) error {
	book := &model.Book{
		ID:    String2Int(params["bookId"]),
		Price: String2Float64(params["price"]),
	}
	//fmt.Println("cart", cart)
	cartItem := &model.CartItem{
		CartId: cart.CartId,
		Count:  1,
		Book:   book,
	}
	err := service.AddCartItem(cartItem)

	if err != nil {
		fmt.Println("加入购物车失败", err.Error())
		return err
	}
	// 更新购物车
	cart.Items = append(cart.Items, cartItem)
	err = service.UpdateCart(cart)
	if err != nil {
		fmt.Println("error=", err.Error())
		return err
	}
	return nil
}

func UpdateCartAndItem(item *model.CartItem, cart *model.Cart) error {
	// 更新item, 更新carts
	item.Count += 1
	err := service.UpdateCartItem(item)
	if err != nil {
		return err
	}
	item.Amount += item.Book.Price
	for i, v := range cart.Items {
		if v.Id == item.Id {
			cart.Items[i] = item
		}
	}
	err = service.UpdateCart(cart)
	if err != nil {
		return err
	}
	return nil
}

func CreateNewCart(params map[string]string) error {
	var items []*model.CartItem
	uuidValue := uuid.New()
	book := &model.Book{
		ID:    String2Int(params["bookId"]),
		Price: String2Float64(params["price"]),
	}

	item := &model.CartItem{
		CartId: uuidValue.String(),
		//Amount: 12,
		Count: 1,
		Book:  book,
	}
	items = append(items, item)
	cart := &model.Cart{
		CartId: uuidValue.String(),
		UserId: String2Int(params["userId"]),
		Items:  items,
	}
	err := service.AddCart(cart)

	if err != nil {
		return err
	}

	return nil
}
