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
	customer := model.NewCustomer("张三", 18, "15522339988", "女", "1234@gmail.com")
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return cs
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

func (cs *CustomerService) InsertCustomer(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

func (cs *CustomerService) UpdateCustomer() bool {

	return true
}

func (cs *CustomerService) DeleteCustomer(id int) bool {
	index := cs.FindById(id)
	if index > 0 {
		cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	} else {
		return false
	}
	return true
}

func (cs *CustomerService) FindById(id int) int {
	index := -1
	for _, customer := range cs.customers {
		if customer.Id == id {
			index = customer.Id
		}
	}
	return index
}
