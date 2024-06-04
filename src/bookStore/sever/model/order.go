package model

import "time"

// Order 订单实体类
type Order struct {
	OrderId     string    `json:"id"`
	CreateTime  time.Time `json:"createTime"`
	TotalCount  int64     `json:"totalCount"`
	TotalAmount float64   `json:"totalAmount"`
	State       int64     `json:"state"` //订单状态， 0 未发货 1 已发货 2 交易完成 3 退货中 4 订单关闭
	UserId      int64     `json:"userId"`
}
