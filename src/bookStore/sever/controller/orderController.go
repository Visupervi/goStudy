package controller

import (
	"bookStore/sever/model"
	"bookStore/sever/service"
	"bookStore/sever/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	res := &model.BookStoreResult{}
	var order *model.Order
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(r.Body)
	// 用于存放参数key=value数据
	var params map[string]string

	// 解析参数 存入map
	err := decoder.Decode(&params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cart, err := service.GetCartByUserId(utils.String2Int(params["userId"]))

	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	if cart != nil {
		uuidValue := uuid.New()
		o := &model.Order{
			OrderId:     uuidValue.String(),
			CreateTime:  time.Now(),
			TotalAmount: cart.TotalAmount,
			TotalCount:  cart.TotalCount,
			State:       0,
			UserId:      utils.String2Int64(params["userId"]),
		}

		items := cart.Items
		err = service.AddOrder(o)

		if err != nil && err != sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		order = o
		for _, v := range items {
			oi := &model.OrderItem{
				Count:   v.Count,
				Amount:  v.Amount,
				OrderId: uuidValue.String(),
				Book: model.Book{
					Title:   v.Book.Title,
					Author:  v.Book.Author,
					Price:   v.Book.Price,
					ImgPath: v.Book.ImgPath,
				},
			}
			err = service.AddOrderItems(oi)
			if err != nil && err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			}

			book := v.Book
			book.Sales = book.Sales + int(v.Count)
			book.Stock = book.Stock - int(v.Count)

			// 更新图书
			service.UpdateBook(book)

		}
		// 清空购物车
		service.ClearCart(utils.String2Int(params["userId"]), params["cartId"])
	} else {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	str, _ := res.ResponseMsg(w, order, http.StatusOK)
	fmt.Fprintf(w, str)
}

func GetOrderList(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	res := &model.BookStoreResult{}

	orders, err := service.GetOrderList()

	if err != nil && orders == nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	str, _ := res.ResponseMsg(w, orders, http.StatusOK)
	fmt.Fprintf(w, str)
}
