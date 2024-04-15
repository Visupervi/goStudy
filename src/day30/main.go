package main

import (
	"encoding/json"
	"fmt"
)

// json
//

// 反序列化

// goroutine 协程
// channel 管道

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

	//myHandler := Handler{}

	// 反序列化
	var a1 Animal
	str := string(data)
	fmt.Println("\nstr=", str)
	err1 := json.Unmarshal([]byte(str), &a1)

	if err != nil {
		fmt.Println(err1)
	}

	fmt.Printf("反序列化后=%v", a1)
	//http.Handle("/", &myHandler)
	//
	//http.ListenAndServe(":8888", nil)
	// 反序列化map

	var b map[string]interface{}

	str11 := string(result)
	fmt.Println("\nstr11=", str11)
	err11 := json.Unmarshal([]byte(str), &b)

	if err != nil {
		fmt.Println(err11)
	}

	fmt.Printf("反序列化后=%v", b)

	// 单元测试框架 testing

}
