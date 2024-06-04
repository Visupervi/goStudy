package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
)

// AddOderItem 向数据库中添加一个订单项
func AddOderItem(oi *model.OrderItem) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()

	sqlStr := "insert into order_items(count,amount,title,author,price,img_path, order_id) values(?,?,?,?,?,?,?)"

	_, err := db.Exec(sqlStr, oi.Count, oi.Amount, oi.Title, oi.Author, oi.Price, oi.ImgPath, oi.OrderId)

	//fmt.Println("eee", err)
	if err != nil {
		return err
	}
	return nil
}

// GetOrdersByOrderId 根据订单号获取订单项
func GetOrdersByOrderId(od string) ([]*model.OrderItem, error) {
	db, error := db.ConnectDB()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	str := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id=?"
	rows, err := db.Query(str, od)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		//books = append(books,)
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemId, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderId)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
