package models

import (
	"log"

	"github/aryan-go/food_ordering_go/package/structures"
)

func AddPaymentDetails(price float64, order_id int, customer_id int) {
	payment_status := "left"
	query := "INSERT INTO payment_table (total_price,payment_status,order_id,customer_id) VALUES (?,?,?,?)"
	_, err := DB.Exec(query, price, payment_status, order_id, customer_id)
	if err != nil {
		log.Fatal("Some err in adding details in payment table : ", err)
	} else {
	}
}

func EditMenu(name string, desc string, price float64, category_id int) {
	query := "INSERT INTO food_menu (food_name,description,price,category_id) VALUES (?,?,?,?)"
	_, err := DB.Exec(query, name, desc, price, category_id)
	if err != nil {
		log.Fatal("Some err in adding details in menu table : ", err)
	}
}

func FindPayment(quant []int, food_id []int) float64 {
	var total_price float64 = 0
	for i := 0; i < len(quant); i++ {
		query := "SELECT * FROM food_menu WHERE food_id = ?"
		result, err := DB.Query(query, food_id[i])
		if err != nil {
			log.Fatal("There is some error in finding the data for this is")
		} else {
			for result.Next() {
				var food_item structures.Food
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

func FindTotalPayment(order_id int, customer_id int) float64 {
	pay_stat := "left"
	query := "SELECT * FROM payment_table WHERE order_id = ? AND customer_id = ? AND payment_status = ?"
	result, err := DB.Query(query, order_id, customer_id, pay_stat)
	var details structures.Payment_table
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

func IncompleteOrderId() []structures.Order {
	query := "SELECT * FROM order_table"
	var ids []structures.Order 
	var orderdetails structures.Order
	result, err := DB.Query(query)
	if err != nil {
		log.Fatal("There is some error in getting the required details")
	} else {
		for result.Next() {
			err := result.Scan(&orderdetails.Order_id, &orderdetails.Customer_id, &orderdetails.Food_status, &orderdetails.Chef_id)
			if err != nil {
				log.Fatal("There is some error in scanning the details for incomplete order id")
			} else {
				ids = append(ids, orderdetails)
			}
		}
	}
	defer result.Close()
	return ids
}

func UnpaidPaymentId() []structures.Payment_table {
	query := "SELECT * FROM payment_table"
	var ids []structures.Payment_table
	var details structures.Payment_table
	result, err := DB.Query(query, status)
	if err != nil {
		log.Fatal("There is some error in getting the required details")
	} else {
		for result.Next() {
			err := result.Scan(&details.Total_price, &details.Tip, &details.Payment_status, &details.Order_id, &details.Customer_id, &details.Payment_id)
			if err != nil {
				log.Fatal("There is some error in scanning details for payment : ", err)
			} else {
				ids = append(ids, details)
			}
		}
	}
	defer result.Close()
	return ids
}

func GetPaymentId(payment_id int) (float64, int) {
	query := "SELECT * FROM payment_table WHERE payment_id = ?"
	var details structures.Payment_table
	result, err := DB.Query(query, payment_id)
	if err != nil {
		log.Fatal("There is some error in getting the required details")
	} else {
		for result.Next() {
			err := result.Scan(&details.Total_price, &details.Tip, &details.Payment_status, &details.Order_id, &details.Customer_id, &details.Payment_id)
			if err != nil {
				log.Fatal("There is some error in scanning details for payment : ", err)
			}
		}
	}
	defer result.Close()
	return details.Total_price, details.Order_id
}

func UpdatePaymentTable(order_id int, customer_id int) {
	payment_status_2 := "completed"
	payment_status_1 := "left"
	query := "UPDATE payment_table SET payment_status = ? WHERE customer_id = ? AND payment_status = ? AND order_id = ?"
	_, err := DB.Exec(query, payment_status_2, customer_id, payment_status_1, order_id)
	if err != nil {
		log.Fatal("There is some error in completing the payment : ", err)
	} else {
	}
}

func UpdatePaymentId(payment_id int) {
	payment_status_2 := "completed"
	payment_status_1 := "left"
	query := "UPDATE payment_table SET payment_status = ? WHERE payment_id = ? AND payment_status = ?"
	_, err := DB.Exec(query, payment_status_2, payment_id, payment_status_1)
	if err != nil {
		log.Fatal("There is some error in completing the payment : ", err)
	} else {
	}
}
