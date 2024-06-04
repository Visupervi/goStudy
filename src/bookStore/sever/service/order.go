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

func UpdateOrderState(state int, od string) error {
	return dao.UpdateOrderState(state, od)
}

func GetOrdersByUserId(id int) ([]*model.Order, error) {
	return dao.GetOrderByUid(id)
}
