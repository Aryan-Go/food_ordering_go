package structures

import "gopkg.in/guregu/null.v3"

type Items_added struct {
	Items_added          []int    `json:"item_add"`
	Special_instructions []string `json:"instructions"`
	Id_arr               []int    `json:"id"`
}

type Order_id struct {
	Id int `json:"id"`
}

type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
	Role       string `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Com_item struct {
	Food_id  int `json:"food_id"`
	Order_id int `json:"order_id"`
}

type Com_item2 struct {
	Order_id int `json:"order_id"`
	Tip      int `json:"tip"`
}

type Payment_details struct {
	Final_payment float64
	Tip           int
}

type Incomplete struct {
	Order_id_order   []int
	Order_id_payment []int
	Customer_chef_id []int
}

type User2 struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Id       int    `json:"id"`
}

type Food struct {
	Food_id     int     `json:"id"`
	Name        string  `json:"name"`
	Desc        string  `json:"description"`
	Price       float64 `json:"price"`
	Category_id int     `json:"c_id"`
}

type Payment_table struct {
	Total_price    float64    `json:"total_price"`
	Tip            null.Float `json:"tip"`
	Payment_status string     `json:"payment_status"`
	Order_id       int        `json:"order_id"`
	Customer_id    int        `json:"customer_id"`
	Payment_id     int        `json:"payment_id"`
}

type Order struct {
	Order_id    int    `json:"order_id"`
	Customer_id int    `json:"customer_id"`
	Food_status string `json:"food_status"`
	Chef_id     int    `json:"chef_id"`
}

type Food_added struct {
	Id           int    `json:"food_id"`
	Quant        int    `json:"quant"`
	Instruct     string `json:"instructions"`
	Food_status  string `json:"status"`
	Order_status int    `json:"order_id"`
	Food_name    string `json:"food_name"`
}

type Error struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

type Customer_id struct {
	Id int `json:"id"`
}
