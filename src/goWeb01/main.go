package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandle struct {
}

func (mh *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "请求地址为：", r.URL)
}
func main() {
	// 适配器
	handle := &MyHandle{}
	http.Handle("/", handle)

	// 监听端口

	server := http.Server{
		Addr:        ":9988",
		Handler:     handle,
		ReadTimeout: 3 * time.Minute,
	}
	server.ListenAndServe()
	//http.ListenAndServe(":9988", nil)
}

func handler(rw http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(rw, "hello Wold", request.URL.Path)

}
