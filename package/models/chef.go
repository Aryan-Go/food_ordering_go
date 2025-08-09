package models

import (
	"database/sql"
	"fmt"
	"log"
)

func Complete_order_item(order_id int, food_status string, food_id int) bool {
	fmt.Println(order_id ,food_status ,food_id)
	query := "UPDATE ordered_items SET food_status = (?) WHERE order_id = (?) AND food_id = (?);"
	result, err := DB.Exec(query, food_status, order_id, food_id)
	if err != nil {
		log.Fatal("There is some error in updating the status of the food : ", sql.ErrNoRows)
		return false
	} else {
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Error that no rows are getting affected", err)
			return false
		}
		if rowsAffected == 0 {
			fmt.Println("No rows updated — check if order_id and food_id match an existing row")
			return false
		}

		fmt.Println("The data of the food_id is updated in order_id")
		return true
	}
}

func Complete_order(order_id int) bool {
	query2 := "SELECT * FROM ordered_items WHERE order_id = ?"
	result, err := DB.Query(query2, order_id)
	if err != nil {
		log.Fatal("There is some error in gettinf data from ordered items", err)
		defer result.Close()
	} else {
		counter1 := 0
		counter2 := 0
		var food_item Food_added
		for result.Next() {
			fmt.Println("I am in")
			counter1++
			err := result.Scan(&food_item.Id, &food_item.Quant, &food_item.Instruct, &food_item.Order_status, &food_item.Food_status)
			if err != nil {
				log.Fatal("There is some error in scaiing for values for added food items")
				defer result.Close()
				return false
			} else {
				if food_item.Food_status == "left" {
					continue
				} else {
					counter2++
				}
			}
		}
		if counter1 == counter2 {
			food_status := "completed"
			query3 := "UPDATE order_table SET food_status = ? WHERE order_id = ? "
			_, err := DB.Exec(query3, food_status, order_id)
			if err != nil {
				log.Fatal("There is some error in updating the order table status", err)
				defer result.Close()
				return false
			} else {
				fmt.Println("This order is completed in order table")
			}
		}
	}
	defer result.Close()
	return true
}
