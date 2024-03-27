package person

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age < 0 || age > 200 {
		fmt.Println("请输入正确年龄")
		return
	} else {
		p.age = age
	}
}

func (p *person) GetAge(name string) int {
	var age int
	if name == p.Name {
		age = p.age
	}
	return age
}

func (p *person) SetSalary(salary float64) {
	if salary < 0 {
		fmt.Println("薪水不正确")
		return
	} else {
		p.salary = salary
	}
}

func (p *person) GetSalary(name string) float64 {
	var sal float64

	if name == p.Name {
		sal = p.salary
	}
	return sal
}
