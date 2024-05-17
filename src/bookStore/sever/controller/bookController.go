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

func InsertBooks(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)

	var book model.Book

	res := &model.BookStoreResult{}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Println("book==", book)
	err = service.InsertBook(&book)
	if err != nil {
		str, _ := res.ResponseMsg(w, "", http.StatusServiceUnavailable)
		fmt.Fprintf(w, str)
		return
	}

	str, _ := res.ResponseMsg(w, "success", http.StatusOK)

	fmt.Fprintf(w, str)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	res := &model.BookStoreResult{}

	bookRes, err := service.GetBooks()
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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeader(w, r)
	var book model.Book
	res := &model.BookStoreResult{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	//fmt.Println("book==", book)
	err = service.DeleteBook(book.ID)
	if err != nil {
		str, err := res.ResponseMsg(w, "failed", http.StatusServiceUnavailable)
		if err != nil {
			fmt.Fprintf(w, "处理失败")
			return
		}
		fmt.Fprintf(w, str)
		return
	}
	str, _ := res.ResponseMsg(w, "success", http.StatusOK)
	fmt.Fprintf(w, str)
}
