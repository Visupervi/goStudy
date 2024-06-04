package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
)

// AddCartsItem 添加购物项
func AddCartsItem(item *model.CartItem) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//fmt.Println(item.Count)
	//fmt.Println(item.GetAmount())
	//fmt.Println(item.Book.ID)
	//fmt.Println(item.CartId)
	_, err := db.Exec(sqlStr, item.Count, item.GetAmount(), item.Book.ID, item.CartId)
	//fmt.Println("res", err)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartItem 根据cartId删除购物项
func DeleteCartItem(cartId string) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()

	sqlStr := "delete from cart_items where cart_id = ?"
	_, err := db.Exec(sqlStr, cartId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartItemById 根据Id删除购物项
func DeleteCartItemById(id string) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()

	sqlStr := "delete from cart_items where id = ?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItem 根据bookId获取购物项
func GetCartItem(bookId int) (*model.CartItem, error) {
	db, error := db.ConnectDB()
	if error != nil {
		return nil, error
	}
	defer db.Close()
	str := "select id,count,amount,cart_id from cart_items where book_id = ?"

	row := db.QueryRow(str, bookId)

	item := &model.CartItem{}
	error = row.Scan(&item.Id, &item.Count, &item.Amount, &item.CartId)
	//fmt.Println("error", error)
	if error != nil {
		return nil, error
	}

	book, _ := GetBookById(bookId)
	item.Book = book
	return item, nil

}

// GetCartItemByUIdAndBookId 根据userId和bookId获取购物项
func GetCartItemByUIdAndBookId(cartId string, bookId int) (*model.CartItem, error) {
	db, error := db.ConnectDB()
	if error != nil {
		return nil, error
	}
	defer db.Close()
	str := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id=?"

	row := db.QueryRow(str, bookId, cartId)

	item := &model.CartItem{}
	error = row.Scan(&item.Id, &item.Count, &item.Amount, &item.CartId)
	//fmt.Println("error", error)
	if error != nil {
		return nil, error
	}

	book, _ := GetBookById(bookId)
	item.Book = book
	return item, nil

}

// GetCartItems 获取所有购物项
func GetCartItems(cId string) ([]*model.CartItem, error) {
	db, error := db.ConnectDB()
	if error != nil {
		return nil, error
	}
	defer db.Close()
	str := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"

	rows, err := db.Query(str, cId)

	if err != nil {
		return nil, err
	}
	var items []*model.CartItem
	for rows.Next() {
		//s

		var bookId int
		item := &model.CartItem{}
		error = rows.Scan(&item.Id, &item.Count, &item.Amount, &bookId, &item.CartId)

		if error != nil {
			return nil, error
		}
		//fmt.Println("bookId", bookId)
		book, _ := GetBookById(bookId)
		item.Book = book
		items = append(items, item)
	}
	return items, nil

}

// UpdateCartItem 更新购物项
func UpdateCartItem(item *model.CartItem) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()
	//id | count | amount | book_id | cart_id
	sqlStr := "update cart_items set count=?,amount=? where id =?"
	_, err := db.Exec(sqlStr, item.Count, item.GetAmount(), item.Id)

	if err != nil {
		return err
	}

	return nil
}
