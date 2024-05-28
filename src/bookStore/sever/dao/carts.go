package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
)

//TotalCount  int64       `json:"totalCount"`
//TotalAmount float64     `json:"totalAmount"`
//Items       []*CartItem `json:"items"`
//CartId      string      `json:"cartId"`
//UserId      int

func AddCart(c *model.Cart) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()

	sqlStr := "insert into carts(id ,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := db.Exec(sqlStr, c.CartId, c.GetCount(), c.GetAmount(), c.UserId)

	//fmt.Println("res", res)
	if err != nil {
		return err
	}
	cartItems := c.Items
	// 将购物项添加到数据库中
	for _, item := range cartItems {
		//fmt.Println("items", item)
		AddCartsItem(item)
	}
	return nil
}

//func GetCartItems(cId string)([]*model.CartItem, error) {
//	db, error := db.ConnectDB()
//	if error != nil {
//		return nil, error
//	}
//	defer db.Close()
//	str := "select * from cart_items where id = ?"
//
//	row := db.QueryRow(str, )
//
//	item := &model.CartItem{}
//	error = row.Scan(&item.Id, &item.Amount, &item.Amount, &item.CartId)
//
//	if error != nil {
//		return nil, error
//	}
//
//	return item, nil
//}
