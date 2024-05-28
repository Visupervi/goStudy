package model

import "fmt"

type Cart struct {
	TotalCount  int64       `json:"totalCount"`
	TotalAmount float64     `json:"totalAmount"`
	Items       []*CartItem `json:"items"`
	CartId      string      `json:"cartId"`
	UserId      int         `json:"userId"`
}

func (c *Cart) GetAmount() float64 {
	var totalAmount float64
	for _, v := range c.Items {

		totalAmount += v.Amount
	}
	fmt.Println("totalAmount", totalAmount)
	return totalAmount
}

func (c *Cart) GetCount() int64 {
	var totalCount int64
	for _, v := range c.Items {
		totalCount += v.Count
	}

	return totalCount
}
