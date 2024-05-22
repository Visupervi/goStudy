package controller

import (
	"bookStore/sever/model"
	"bookStore/sever/service"
	"bookStore/sever/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	var user model.User
	res := &model.BookStoreResult{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//// 响应成功消息
	u, err := service.Login(user.UserName, user.Password)
	if err != nil || u == nil {
		str, err := res.ResponseMsg(w, u, http.StatusServiceUnavailable)
		if err != nil {
			fmt.Fprintf(w, "处理失败")
			return
		}
		fmt.Fprintf(w, str)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:     "user",
		Value:    user.UserName,
		MaxAge:   10000,
		HttpOnly: true,
	}

	//w.Header().Set("Set-Cookie", cookie.String())
	http.SetCookie(w, &cookie)
	str, err := res.ResponseMsg(w, u, http.StatusOK)
	if err != nil {
		fmt.Fprintf(w, "处理失败")
		return
	}
	fmt.Fprintf(w, str)
}
