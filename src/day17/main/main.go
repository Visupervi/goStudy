package main

import "fmt"

type Animal struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Voice  string `json:"voice"`
	Color  string `json:"color"`
}

func (a Animal) getFiles() {
	fmt.Printf("姓名=%v\n年龄=%v\n性别=%v\n叫声=%v\n颜色=%v\n", a.Name, a.Age, a.Gender, a.Voice, a.Color)
}
func main() {
	dog := Animal{"Tom", 3, "female", "汪汪汪", "black"}
	dog.getFiles()
	cat := Animal{"Jerry", 2, "male", "喵喵喵", "white"}
	cat.getFiles()
}
