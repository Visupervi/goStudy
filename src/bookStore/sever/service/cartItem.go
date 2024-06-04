package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

// AddCartItem 添加购物项
func AddCartItem(item *model.CartItem) error {
	return dao.AddCartsItem(item)
}

// DeleteCartItemById 删除购物项
func DeleteCartItemById(id string) error {
	return dao.DeleteCartItemById(id)
}

// GetCartItem 获取购物项
func GetCartItem(bookId int) (*model.CartItem, error) {
	return dao.GetCartItem(bookId)
}

// GetCartItemByBookIdAndCartId 根据用图书id和购物车id获取购物项
func GetCartItemByBookIdAndCartId(cartId string, bookId int) (*model.CartItem, error) {
	return dao.GetCartItemByUIdAndBookId(cartId, bookId)
}
func GetCartItems(cardId string) ([]*model.CartItem, error) {
	return dao.GetCartItems(cardId)
}

// UpdateCartItem 更新购物项
func UpdateCartItem(item *model.CartItem) error {
	return dao.UpdateCartItem(item)
}
