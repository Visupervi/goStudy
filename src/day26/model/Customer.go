package model

type Customer struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	PhoneNo string `json:"phone_no"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
}

func NewCustomer(name string, age int, phone string, gender string, id string, email string) Customer {
	return Customer{
		Name:    name,
		Age:     age,
		PhoneNo: phone,
		Gender:  gender,
		Id:      id,
		Email:   email,
	}
}
