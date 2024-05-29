package controller

import (
	"bookStore/sever/model"
	"bookStore/sever/service"
	"bookStore/sever/utils"
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
	cart, err := service.GetCartByUserId(user.ID)
	if err != nil {
		str, err := res.ResponseMsg(w, "failed", http.StatusServiceUnavailable)
		if err != nil {
			fmt.Fprintf(w, "处理失败")
			return
		}
		fmt.Fprintf(w, str)
		return
	}
	str, _ := res.ResponseMsg(w, cart, http.StatusOK)
	fmt.Fprintf(w, str)
}
