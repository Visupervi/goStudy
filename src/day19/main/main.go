package main

import (
	"day19/person"
	"day19/student"
	"fmt"
)

func main() {

	//account := account.Account{"浦发", "123", 10}

	p := person.NewPerson("美队")

	p.SetSalary(1200)
	p.SetAge(120)

	fmt.Println(*p)

	p1 := person.NewPerson("灭霸")

	p1.SetAge(100)
	p1.SetSalary(12.22)
	fmt.Println(*p1)

	stu := student.Student{Name: "小学生", Age: 8, Score: 99, No: 202409100}
	pupil := student.Pupil{Student: stu}
	fmt.Println(pupil)

	pupil.Reading("自然")

	stu1 := student.Student{Name: "大学生", Age: 18, Score: 99, No: 202409111}

	gra := student.Graduate{Student: stu1}

	gra.Reading("英语四级")
	//account.LookMoney("123")
	//account.SaveMoney(10, "123")
	//account.GetMoney(10, "123")
}

// 继承
