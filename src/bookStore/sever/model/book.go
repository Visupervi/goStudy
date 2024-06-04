package model

// Book 图书实体类
type Book struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Author  string  `json:"author"`
	Price   float64 `json:"price"`
	Sales   int     `json:"sales"`
	Stock   int     `json:"stock"`
	ImgPath string  `json:"imgPath"`
}
