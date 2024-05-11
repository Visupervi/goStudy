package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MyHandle struct {
}

type User struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

func (mh *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域
	//w.Header().Set("Content-Length", "65536") // 允许所有域
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	//fmt.Fprintln(w, "请求地址为：", r.URL)
	//fmt.Fprintln(w, "请求地址为：", r.URL.RawQuery)
	//fmt.Fprintln(w, "请求Header：", r.Header)
	//len := r.ContentLength
	////
	//body := make([]byte, len)
	////
	//r.Body.Read(body)
	//fmt.Fprintln(w, "请求body：")
	//w.Write([]byte("Hello, cross-origin!\n"))
	//fmt.Println("json.NewDecoder(r.Body)", json.NewDecoder(r.Body))

	// json.NewDecodejson.NewDecoder 和 json.Unmarshal 都是 Go 中用于处理 JSON 数据的方法，它们之间有一些区别和适用场景：
	//
	//json.NewDecoder：
	//
	//json.NewDecoder 是一个构造函数，用于创建一个 Decoder 对象，用于从输入流中读取 JSON 数据。
	//通常情况下，json.NewDecoder 结合 Decoder 对象的方法如 Decode 可以用于解析大型或者流式的 JSON 数据。
	//json.NewDecoder 返回的 Decoder 对象可以从 io.Reader 中读取 JSON 数据，而不需要一次性将所有数据加载到内存中。
	//json.NewDecoder 主要适用于需要逐步处理大型 JSON 数据的场景。
	//json.Unmarshal：
	//
	//json.Unmarshal 函数用于将 JSON 数据解析为 Go 中的数据结构，如结构体、切片、map 等。
	//当你已经有了一个 JSON 字节切片或者字符串，且想要将其直接解析为 Go 中的数据结构时，可以使用 json.Unmarshal。
	//json.Unmarshal 将完整的 JSON 数据一次性解析为指定的 Go 数据结构，适用于小型或者已经完全接收到的 JSON 数据。
	//总的来说，如果你需要逐步处理大型 JSON 数据或者从网络流中读取 JSON 数据，可以使用 json.NewDecoder；而如果你已经有了完整的 JSON 数据并且想要将其解析为 Go 中的数据结构，可以使用 json.Unmarshal。
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Printf("Received JSON data: %+v\n", user)
	u := User{
		Name: "tom",
		Pwd:  "12333",
	}
	data, _ := json.Marshal(u)

	//fmt.Println("data==", string(data))
	//fmt.Println("body==", string(body))
	//err := json.Unmarshal(body, &user)

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	fmt.Println("user", user.Name)
	// 响应成功消息
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	//fmt.Fprintf(w, "Received JSON data successfully\n")

}
func main() {
	// 适配器
	handle := &MyHandle{}
	http.Handle("/", handle)
	//
	//// 监听端口
	//
	server := http.Server{
		Addr:        ":9988",
		Handler:     handle,
		ReadTimeout: 3 * time.Minute,
	}
	server.ListenAndServe()
	//http.ListenAndServe(":9988", nil)
	//db.InitDB()
	//user := &model.User{
	//	UserName: "admin",
	//	Password: "123456",
	//	Email:    "123456@qq.com",
	//}
	//err := user.AddUser(user)
	//if err != nil {
	//	fmt.Println("err=", err)
	//}

}

func handler(rw http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(rw, "hello Wold", request.URL.Path)

}

//报文格式

//请求首行
//请求头信息
//空行
//请求体
