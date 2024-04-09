package service

import (
	"day26/model"
)

type CustomerService struct {
	customers   []model.Customer `json:"customers"`
	customerNum int              `json:"customer_num"`
}

func InitSplice() *CustomerService {

	cs := &CustomerService{}
	cs.customerNum = 1
	//customer := model.Customer{
	//	Name: "张三",
	//	Id: "1",
	//	Age: 18,
	//	Gender: "女",
	//	PhoneNo: "15522339988",
	//	Email: "1234@gmail.com",
	//}

	customer := model.NewCustomer("张三", 18, "15522339988", "女", "1", "1234@gmail.com")

	cs.customers = append(cs.customers, customer)
	return cs
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}
