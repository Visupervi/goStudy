package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

// AddCart 添加购物车
func AddCart(c *model.Cart) error {
	return dao.AddCart(c)
}

// GetCartByUserId 根据userId获取购物车
func GetCartByUserId(userId int) (*model.Cart, error) {
	return dao.GetCartByUid(userId)
}

// UpdateCart 更新购物车
func UpdateCart(cart *model.Cart) error {
	return dao.UpdateCart(cart)
}

// ClearCart 清空购物车
func ClearCart(userId int, cartId string) error {
	return dao.DeleteCart(userId, cartId)
}
