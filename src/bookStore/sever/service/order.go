package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

func AddOrder(o *model.Order) error {
	return dao.AddOder(o)
}

func GetOrderList() ([]*model.Order, error) {
	return dao.GetOrders()
}
