package structures

import "gopkg.in/guregu/null.v3"

type Complete_payment_item struct {
	Order_id int `json:"order_id"`
	Tip      int `json:"tip"`
}

type Payment_details struct {
	Final_payment float64
	Tip           int
}

type Payment_table struct {
	Total_price    float64    `json:"total_price"`
	Tip            null.Float `json:"tip"`
	Payment_status string     `json:"payment_status"`
	Order_id       int        `json:"order_id"`
	Customer_id    int        `json:"customer_id"`
	Payment_id     int        `json:"payment_id"`
}