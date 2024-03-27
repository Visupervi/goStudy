package student

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
	No    int64
}

func (s *Student) Reading(book string) {
	fmt.Printf("%v正在reading%v", s.Name, book)
}

// 通过匿名结构体实现继承

type Graduate struct {
	Student
}

func (g *Graduate) testing() {

}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {

}

//func NewStudent(name string, age int, score float64) *student{
//	return &student{
//		Name: name,
//		Age: age,
//		Score: score,
//	}
//}
