package main

import (
	"fmt"
	"goWeb01/db"
	"goWeb01/model"
	"net/http"
)

type MyHandle struct {
}

func (mh *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	fmt.Fprintln(w, "请求地址为：", r.URL)
	//w.Write([]byte("Hello, cross-origin!\n"))
}
func main() {
	// 适配器
	//handle := &MyHandle{}
	//http.Handle("/", handle)
	//
	//// 监听端口
	//
	//server := http.Server{
	//	Addr:        ":9988",
	//	Handler:     handle,
	//	ReadTimeout: 3 * time.Minute,
	//}
	//server.ListenAndServe()
	//http.ListenAndServe(":9988", nil)
	db.InitDB()
	user := &model.User{
		UserName: "admin",
		Password: "123456",
		Email:    "123456@qq.com",
	}
	err := user.AddUser(user)
	if err != nil {
		fmt.Println("err=", err)
	}

}

func handler(rw http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(rw, "hello Wold", request.URL.Path)

}

//报文格式

//请求首行
//请求头信息
//空行
//请求体
