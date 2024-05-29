package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

func AddCartItem(item *model.CartItem) error {
	return dao.AddCartsItem(item)
}

func DeleteCartItem(id int64) error {
	return dao.DeleteCartItem(id)
}

func GetCartItem(bookId int) (*model.CartItem, error) {
	return dao.GetCartItem(bookId)
}

func GetCartItems(cardId string) ([]*model.CartItem, error) {
	return dao.GetCartItems(cardId)
}
