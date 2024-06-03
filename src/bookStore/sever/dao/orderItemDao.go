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
