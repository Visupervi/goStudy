package model

// CartItem
///**

type CartItem struct {
	Id     int64   `json:"id"`     // 购物项Id
	Book   *Book   `json:"book"`   // 购物项图书信息
	Count  int64   `json:"count"`  // 购物项图书数量
	Amount float64 `json:"amount"` // 购物项金额
	CartId string  `json:"cartId"`
}

// GetAmount 计算购物项的金额
func (ci *CartItem) GetAmount() float64 {
	price := ci.Book.Price
	return float64(ci.Count) * price
}
