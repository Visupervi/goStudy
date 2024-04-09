package model

import "fmt"

type Customer struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	PhoneNo string `json:"phone_no"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
}

func NewCustomer(name string, age int, phone string, gender string, email string) Customer {
	return Customer{
		Name:    name,
		Age:     age,
		PhoneNo: phone,
		Gender:  gender,
		Email:   email,
	}
}

func (c *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", c.Id, c.Name, c.Gender, c.Age, c.PhoneNo, c.Email)
	return info
}
