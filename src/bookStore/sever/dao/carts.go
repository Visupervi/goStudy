package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
	"fmt"
)

// AddCart 添加购物车
func AddCart(c *model.Cart) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()

	sqlStr := "insert into carts(id ,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := db.Exec(sqlStr, c.CartId, c.GetCount(), c.GetAmount(), c.UserId)

	//fmt.Println("eee", err)
	if err != nil {
		return err
	}
	cartItems := c.Items
	// 将购物项添加到数据库中
	for _, item := range cartItems {
		//fmt.Println("items", item.Book)
		AddCartsItem(item)
	}
	return nil
}

// GetCartByUid 根据userId 获取购物车
func GetCartByUid(uId int) (*model.Cart, error) {
	db, error := db.ConnectDB()
	if error != nil {
		return nil, error
	}
	defer db.Close()
	str := "select id,total_count,total_amount, user_id from carts where user_id = ?"

	row := db.QueryRow(str, uId)

	cart := &model.Cart{}
	error = row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if error != nil {
		//fmt.Println("error", error)
		return nil, error
	}
	items, err := GetCartItems(cart.CartId)
	//fmt.Println("error", err)
	if err != nil {
		return nil, err
	}

	cart.Items = items

	return cart, nil
}

// UpdateCart 更新购物车
func UpdateCart(c *model.Cart) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()
	//id | count | amount | book_id | cart_id
	sqlStr := "update carts set total_count=?,total_amount=? where id =?"
	_, err := db.Exec(sqlStr, c.GetCount(), c.GetAmount(), c.CartId)

	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}

	return nil
}

// DeleteCart 删除购物车
func DeleteCart(uId int, cartId string) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()
	err := DeleteCartItem(cartId)
	if err != nil {
		return err
	}
	sqlStr := "delete from carts where user_id = ? and id =?"
	_, err = db.Exec(sqlStr, uId, cartId)
	if err != nil {
		return err
	}
	return nil
}