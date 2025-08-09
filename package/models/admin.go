package models

import (
	"fmt"
	"log"

	"gopkg.in/guregu/null.v3"
)

func Add_payment_details(price float64, order_id int, customer_id int) {
	payment_status := "left"
	query := "INSERT INTO payment_table (total_price,payment_status,order_id,customer_id) VALUES (?,?,?,?)"
	_, err := DB.Query(query, price, payment_status, order_id, customer_id)
	if err != nil {
		log.Fatal("Some err in adding details in payment table : ", err)
	} else {
		fmt.Println("Data added successfully")
	}
}

func Find_payment(quant []int, food_id []int) float64 {
	var total_price float64 = 0
	for i := 0; i < len(quant); i++ {
		query := "SELECT * FROM food_menu WHERE food_id = ?"
		result, err := DB.Query(query, food_id[i])
		if err != nil {
			log.Fatal("There is some error in finding the data for this is")
		} else {
			for result.Next() {
				var food_item Food
				err := result.Scan(&food_item.Food_id, &food_item.Name, &food_item.Desc, &food_item.Price, &food_item.Category_id)
				if err != nil {
					log.Fatal("There is some error in scanning for details", err)
				} else {
					total_price = total_price + (float64(quant[i]) * food_item.Price)
					defer result.Close()
				}
			}
		}
	}
	return total_price
}

type Payment_table struct {
	Total_price    float64    `json:"total_price"`
	Tip            null.Float `json:"tip"`
	Payment_status string     `json:"payment_status"`
	Order_id       int        `json:"order_id"`
	Customer_id    int        `json:"customer_id"`
	Payment_id     int        `json:"payment_id"`
}

func Find_total_payment(order_id int, customer_id int) float64 {
	pay_stat := "left"
	query := "SELECT * FROM payment_table WHERE order_id = ? AND customer_id = ? AND payment_status = ?"
	result, err := DB.Query(query, order_id, customer_id, pay_stat)
	var details Payment_table
	if err != nil {
		log.Fatal("There is some error in getting the payment data")
	} else {
		for result.Next() {
			err := result.Scan(&details.Total_price, &details.Tip, &details.Payment_status, &details.Order_id, &details.Customer_id, &details.Payment_id)
			if err != nil {
				log.Fatal("There is some error in scanning details for payment : ", err)
			}
		}
	}
	defer result.Close()
	return details.Total_price
}

type Order struct {
	Order_id    int    `json:"order_id"`
	Customer_id int    `json:"customer_id"`
	Food_status string `json:"food_status"`
	Chef_id     int    `json:"chef_id"`
}

func Incomplete_order_id() []int {
	status := "left"
	query := "SELECT * FROM order_table WHERE food_status=?"
	var ids []int
	var orderdetails Order
	result, err := DB.Query(query, status)
	if err != nil {
		log.Fatal("There is some error in getting the required details")
	} else {
		for result.Next() {
			err := result.Scan(&orderdetails.Order_id, &orderdetails.Customer_id, &orderdetails.Food_status, &orderdetails.Chef_id)
			if err != nil {
				log.Fatal("There is some error in scanning the details for incomplete order id")
			} else {
				ids = append(ids, orderdetails.Order_id)
			}
		}
	}
	defer result.Close()
	return ids
}

func Unpaid_order_id() []int {
	status := "left"
	query := "SELECT * FROM payment_table WHERE payment_status=?"
	var ids []int
	var details Payment_table
	result, err := DB.Query(query, status)
	if err != nil {
		log.Fatal("There is some error in getting the required details")
	} else {
		for result.Next() {
			err := result.Scan(&details.Total_price, &details.Tip, &details.Payment_status, &details.Order_id, &details.Customer_id, &details.Payment_id)
			if err != nil {
				log.Fatal("There is some error in scanning details for payment : ", err)
			} else {
				ids = append(ids, details.Order_id)
			}
		}
	}
	defer result.Close()
	return ids
}

func Update_payment_table(order_id int, customer_id int) {
	payment_status_2 := "completed"
	payment_status_1 := "left"
	query := "UPDATE payment_table SET payment_status = ? WHERE customer_id = ? AND payment_status = ? AND order_id = ?"
	_, err := DB.Exec(query, payment_status_2, customer_id, payment_status_1, order_id)
	if(err != nil){
		log.Fatal("There is some error in completing the payment : " , err)
	} else{
		fmt.Println("The payment is completed")
	}
}
