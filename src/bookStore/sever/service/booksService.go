package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

func GetBooksSlice(page int, size int) (*model.PageResult, error) {
	//fmt.Println(page)
	return dao.GetBooksByPage(page, size)
}

func InsertBook(book *model.Book) error {
	return dao.InsertBook(book)
}

func GetBooks() ([]*model.Book, error) {
	return dao.GetBooks()
}

func DeleteBook(id int) error {
	return dao.DeleteBook(id)
}

func UpdateBook(book *model.Book) error {
	return dao.UpdateBook(book)
}
