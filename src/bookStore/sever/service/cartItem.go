package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

func AddCartItem(item *model.CartItem) error {
	return dao.AddCartsItem(item)
}

func DeleteCartItemById(id string) error {
	return dao.DeleteCartItemById(id)
}

func GetCartItem(bookId int) (*model.CartItem, error) {
	return dao.GetCartItem(bookId)
}
func GetCartItemByBookIdAndCartId(cartId string, bookId int) (*model.CartItem, error) {
	return dao.GetCartItemByUIdAndBookId(cartId, bookId)
}
func GetCartItems(cardId string) ([]*model.CartItem, error) {
	return dao.GetCartItems(cardId)
}

func UpdateCartItem(item *model.CartItem) error {
	return dao.UpdateCartItem(item)
}
