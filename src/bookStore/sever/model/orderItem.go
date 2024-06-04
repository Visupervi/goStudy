package model

// OrderItem 订单项实体类
type OrderItem struct {
	OrderItemId int64   `json:"orderItemId"`
	Count       int64   `json:"count"`
	Amount      float64 `json:"amount"`
	OrderId     string  `json:"orderId"`
	Book
}
