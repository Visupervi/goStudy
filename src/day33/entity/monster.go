package entity

import "fmt"

type Monster struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Score  float64 `json:"score"`
	Gender string  `json:"gender"`
}

func (m Monster) Print() {
	fmt.Println("print start")
	fmt.Println(m)
	fmt.Println("print end")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, score float64, gender string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Gender = gender

}
