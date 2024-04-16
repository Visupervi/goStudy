package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// json
// 将结构体、map、切片序列化

type Animal struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Color string   `json:"color"`
	Voice string   `json:"voice"`
	Food  []string `json:"food"`
}

func main() {
	animal := Animal{
		Name:  "大黄",
		Age:   10,
		Color: "黄",
		Voice: "汪汪汪",
	}

	foofs := []string{"猪肉", "鸡肉"}
	animal.Food = append(animal.Food, foofs...)

	data, err := json.Marshal(&animal)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("data=", string(data))

	var a map[string]interface{}

	// 使用map前，需要make一下
	a = make(map[string]interface{})

	a["name"] = "熏悟空"
	a["age"] = 500
	a["gender"] = "男"
	res, err := json.Marshal(a)

	fmt.Println("res=", string(res))

	var slice []map[string]interface{}

	for i := 0; i < 5; i++ {
		a = make(map[string]interface{})
		a["name"] = "熏悟空"
		a["age"] = 500 + i
		a["gender"] = "男"
		a["address"] = "花果山"

		slice = append(slice, a)
	}

	result, err := json.Marshal(slice)

	fmt.Println("result=", string(result))

	myHandler := Handler{}

	http.Handle("/", &myHandler)

	http.ListenAndServe(":8888", nil)

}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go World"))
}
