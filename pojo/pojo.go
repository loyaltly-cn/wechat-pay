package pojo

type Pay struct {
	OpenId  string `json:"openId"`
	OrderId string `json:"orderId"`
	Desc    string `json:"desc"`
	Price   int64  `json:"price"`
}

type Status struct {
	OrderId string `json:"orderId"`
}

type Refund struct {
	OrderId      string `json:"orderId"`
	RefundAmount int64  `json:"refundAmount"`
	TotalAmount  int64  `json:"totalAmount"`
}
