package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

func AddOrderItems(oi *model.OrderItem) error {
	return dao.AddOderItem(oi)
}
