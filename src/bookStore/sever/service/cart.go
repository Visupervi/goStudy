package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

func AddCart(c *model.Cart) error {
	return dao.AddCart(c)
}

func GetCartByUserId(userId int) (*model.Cart, error) {
	return dao.GetCartByUid(userId)
}
