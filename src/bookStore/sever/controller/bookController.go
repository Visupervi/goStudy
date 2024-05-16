package controller

import (
	"bookStore/sever/model"
	"bookStore/sever/service"
	"bookStore/sever/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetBookSlice(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	var pages model.PageResult
	res := &model.BookStoreResult{}
	err := json.NewDecoder(r.Body).Decode(&pages)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		//fmt.Println("err", err)
		return
	}
	//fmt.Println("err", pages)
	bookRes, err := service.GetBooksSlice(pages.Page, pages.Limit)

	if err != nil {
		str, err := res.ResponseMsg(w, bookRes, http.StatusServiceUnavailable)
		if err != nil {
			fmt.Fprintf(w, "处理失败")
			return
		}
		fmt.Fprintf(w, str)
		return
	}
	str, _ := res.ResponseMsg(w, bookRes, http.StatusOK)
	fmt.Fprintf(w, str)
}
