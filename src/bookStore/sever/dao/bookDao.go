package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
	"fmt"
)

func GetBooksByPage(page int, limit int) (*model.PageResult, error) {
	db, error := db.ConnectDB()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	// 计算偏移量
	offset := (page - 1) * limit

	// 获取总条目数
	var total int
	err := db.QueryRow("SELECT COUNT(*) FROM books").Scan(&total)

	//fmt.Println("total", total)
	if err != nil {
		//fmt.Println("db.QueryRow", err)
		return nil, error
	}

	// 获取分页数据
	rows, err := db.Query("SELECT id,title,author,price,sales,stock,img_path FROM books LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		fmt.Println("db.Query", err)
		return nil, error
	}
	defer rows.Close()
	var books []*model.Book
	for rows.Next() {
		//books = append(books,)
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//if err != nil {
		//	return nil, err
		//}
		//fmt.Println("book", book)
		books = append(books, book)
	}
	//if err := rows.Err(); err != nil {
	//	return nil, err
	//}
	// 计算总页数
	totalPages := (total + limit - 1) / limit

	pageResponse := &model.PageResult{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		Books:      books,
	}
	return pageResponse, nil
}

func GetBooks() ([]*model.Book, error) {
	db, error := db.ConnectDB()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := db.Query(sqlStr)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		//books = append(books,)
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

func InsertBook(b *model.Book) (err error) {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)

	//fmt.Println("res", res)
	if err != nil {
		return err
	}
	return nil
}
