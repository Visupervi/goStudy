package model

type Cart struct {
	TotalCount  int64       `json:"totalCount"`
	TotalAmount float64     `json:"totalAmount"`
	Items       []*CartItem `json:"items"`
	CartId      string      `json:"cartId"`
	UserId      int         `json:"userId"`
}

// GetAmount 计算总价格
func (c *Cart) GetAmount() float64 {
	var totalAmount float64
	for _, v := range c.Items {
		//fmt.Println("totalAmount---book", v.Book)
		//fmt.Println("totalAmount---v", v)
		amount := v.Book.Price * float64(v.Count)

		totalAmount += amount
	}

	return totalAmount
}

// GetCount 计算总数量
func (c *Cart) GetCount() int64 {
	var totalCount int64
	for _, v := range c.Items {
		totalCount += v.Count
	}

	return totalCount
}
