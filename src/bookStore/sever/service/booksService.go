package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

// GetBooksSlice 分页获取图书
func GetBooksSlice(page int, size int) (*model.PageResult, error) {
	//fmt.Println(page)
	return dao.GetBooksByPage(page, size)
}

// InsertBook 新增图书
func InsertBook(book *model.Book) error {
	return dao.InsertBook(book)
}

// GetBooks 获取图书
func GetBooks() ([]*model.Book, error) {
	return dao.GetBooks()
}

// DeleteBook 删除图书
func DeleteBook(id int) error {
	return dao.DeleteBook(id)
}

// UpdateBook 更新图书
func UpdateBook(book *model.Book) error {
	return dao.UpdateBook(book)
}
