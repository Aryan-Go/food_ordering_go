package structures

type Items_added struct {
	Items_added          []int    `json:"item_add"`
	Special_instructions []string `json:"instructions"`
	Id_arr               []int    `json:"id"`
}

type Order_id struct {
	Id int `json:"id"`
}

type Complete_item struct {
	Food_id  int `json:"food_id"`
	Order_id int `json:"order_id"`
}

type Incomplete struct {
	Order_id_order   []Order `json:"incomplete_order"`
	Payment_id       []Payment_table `json:"incomplete_payment"`
	Customer_chef_id []int `json:"chef_customer"`
}

type Food struct {
	Food_id     int     `json:"id"`
	Name        string  `json:"name"`
	Desc        string  `json:"description"`
	Price       float64 `json:"price"`
	Category_id int     `json:"c_id"`
}

type Order struct {
	Order_id    int    `json:"order_id"`
	Customer_id int    `json:"customer_id"`
	Food_status string `json:"food_status"`
	Chef_id     int    `json:"chef_id"`
}

type Food_added struct {
	Id          int    `json:"food_id"`
	Quant       int    `json:"quant"`
	Instruct    string `json:"instructions"`
	Food_status string `json:"status"`
	Order_id    int    `json:"order_id"`
	Food_name   string `json:"food_name"`
}
