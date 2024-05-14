package service

import (
	"bookStore/sever/controller"
	"bookStore/sever/model"
	"bookStore/sever/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegistryService(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	var user model.User
	res := &model.BookStoreResult{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//// 响应成功消息
	u, err := controller.Registry(user.UserName, user.Password, user.Email)
	if err != nil || !u {
		str, err := res.ResponseMsg(w, u, http.StatusServiceUnavailable)
		if err != nil {
			fmt.Println("用户名称已存在")
			return
		}
		fmt.Fprintf(w, str)
		return
	}
	str, err := res.ResponseMsg(w, u, http.StatusOK)
	if err != nil {
		fmt.Println("请求出错")
		return
	}
	fmt.Fprintf(w, str)
}
