package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
	"fmt"
)

// AddOder 向数据库中添加一个订单
func AddOder(o *model.Order) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()

	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := db.Exec(sqlStr, o.OrderId, o.CreateTime, o.TotalCount, o.TotalAmount, o.State, o.UserId)

	//fmt.Println("eee", err)
	if err != nil {
		return err
	}
	return nil
}

// GetOrderByUid 根据userId获取订单
func GetOrderByUid(uId int) ([]*model.Order, error) {
	db, error := db.ConnectDB()
	if error != nil {
		return nil, error
	}
	defer db.Close()
	str := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"

	rows, err := db.Query(str, uId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		//books = append(books,)
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		orders = append(orders, order)
	}
	return orders, nil
	return orders, nil
}

// UpdateOrder 更新购物车
func UpdateOrder(o *model.Order) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()
	//id | count | amount | book_id | cart_id
	sqlStr := "update orders set total_count=?,total_amount=?, state=? where id =?"
	_, err := db.Exec(sqlStr, o.TotalCount, o.TotalAmount, o.State)

	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}

	return nil
}

//func DeleteOrder(uId int, orderId string) error {
//	db, error := db.ConnectDB()
//	if error != nil {
//		return error
//	}
//	defer db.Close()
//	err := DeleteCartItem(cartId)
//	if err != nil {
//		return err
//	}
//	sqlStr := "delete from carts where user_id = ? and id =?"
//	_, err = db.Exec(sqlStr, uId, cartId)
//	if err != nil {
//		return err
//	}
//	return nil
//}

// GetOrders 获取所有订单
func GetOrders() ([]*model.Order, error) {
	db, error := db.ConnectDB()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	str := "select id,create_time,total_count,total_amount,state,user_id from orders"
	rows, err := db.Query(str)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		//books = append(books,)
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		orders = append(orders, order)
	}
	return orders, nil
}

// UpdateOrderState 更新订单状态
func UpdateOrderState(state int, od string) error {
	db, error := db.ConnectDB()

	if error != nil {
		return error
	}
	defer db.Close()
	sqlStr := "update orders set state=? where id =?"
	_, err := db.Exec(sqlStr, state, od)

	if err != nil {
		return err
	}

	return nil
}