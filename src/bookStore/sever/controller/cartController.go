package controller

import (
	"bookStore/sever/model"
	"bookStore/sever/service"
	"bookStore/sever/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	var user model.User
	res := &model.BookStoreResult{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	//fmt.Println("book==", book)
	cart, _ := service.GetCartByUserId(user.ID)
	str, _ := res.ResponseMsg(w, cart, http.StatusOK)
	fmt.Fprintf(w, str)
}

// AddCart 的逻辑
// 先看cart存在不存在
// 然后在看item存在不存在
// 都不存在，新建添加
// cart存在，item不存在，item插入cart_items中
// cart存在， item存在， 更新cart_items中的count,amount以及cart中的totalCount和totalAmount
// 不会出现cart不存在，items存在的情况

func AddCart(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	//var book *model.Book
	//var user model.User
	res := &model.BookStoreResult{}
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(r.Body)
	// 用于存放参数key=value数据
	var params map[string]string

	// 解析参数 存入map
	err := decoder.Decode(&params)
	cart, err := service.GetCartByUserId(utils.String2Int(params["userId"]))

	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	if cart != nil {
		item, _ := service.GetCartItemByBookIdAndCartId(cart.CartId, utils.String2Int(params["bookId"]))
		if item != nil {
			err = utils.UpdateCartAndItem(item, cart)
			if err != nil && err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			}
		} else {
			fmt.Println("come here")
			err = utils.UpdateCartAddItem(params, cart)
			if err != nil && err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			}
			//return
		}

	} else {
		err = utils.CreateNewCart(params)
		if err != nil && err != sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		//return
	}
	//fmt.Printf("POST json: userId=%s, bookId=%s\n", params["userId"], params["bookId"])
	str, _ := res.ResponseMsg(w, "success", http.StatusOK)
	fmt.Fprintf(w, str)
}
