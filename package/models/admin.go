package models

import (
	"fmt"
	"log"
	"gopkg.in/guregu/null.v3"
)

func Add_payment_details(price float64, order_id int , customer_id int){
	payment_status := "left";
	query := "INSERT INTO payment_table (total_price,payment_status,order_id,customer_id) VALUES (?,?,?,?)"
    _,err := DB.Query(query,price, payment_status, order_id, customer_id);
	if(err != nil){
		log.Fatal("Some err in adding details in payment table : " , err)
	}else{
		fmt.Println("Data added successfully")
	}
}

func Find_payment(quant []int, food_id []int) float64 {
	var total_price float64 = 0;
    for i := 0; i < len(quant) ; i++ {
		query := "SELECT * FROM food_menu WHERE food_id = ?"
        result,err := DB.Query(query , food_id[i])
		if(err != nil){
			log.Fatal("There is some error in finding the data for this is")
		} else{
			for result.Next(){
				var food_item Food
				err := result.Scan(&food_item.Food_id , &food_item.Name , &food_item.Desc , &food_item.Price , &food_item.Category_id)
				if(err != nil){
					log.Fatal("There is some error in scanning for details" , err)
				} else{
					total_price = total_price + (float64(quant[i]) * food_item.Price);

				}
			}
		}
    }
    return total_price;
}

type Payment_table struct{
	Total_price float64 `json:"total_price"`
    Tip null.Float `json:"tip"`
    Payment_status string `json:"payment_status"`
    Order_id int `json:"order_id"`
    Customer_id int `json:"customer_id"`
    Payment_id int `json:"payment_id"`
}

func Find_total_payment(order_id int , customer_id int)float64{
	pay_stat := "left"
	query := "SELECT * FROM payment_table WHERE order_id = ? AND customer_id = ? AND payment_status = ?"
	result,err := DB.Query(query , order_id , customer_id , pay_stat)
	var details Payment_table
	if(err != nil){
		log.Fatal("There is some error in getting the payment data")
	} else{
		for result.Next(){
			err := result.Scan(&details.Total_price , &details.Tip , &details.Payment_status , &details.Order_id , &details.Customer_id , &details.Payment_id)
			if(err != nil){
				log.Fatal("There is some error in scanning details for payment : " , err)
			}
		}
	}
	return details.Total_price
}