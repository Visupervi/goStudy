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

func UpdateCart(cart *model.Cart) error {
	return dao.UpdateCart(cart)
}

func ClearCart(userId int, cartId string) error {
	return dao.DeleteCart(userId, cartId)
}
