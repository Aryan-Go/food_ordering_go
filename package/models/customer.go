package models

import (
	"fmt"
	"log"
)

type Food struct {
	Food_id     int     `json:"id"`
	Name        string  `json:"name"`
	Desc        string  `json:"description"`
	Price       float64 `json:"price"`
	Category_id int     `json:"c_id"`
}

func Customer_to_chef(email string) {
	query := "UPDATE user SET role = (?) WHERE email = (?)"
	_, err := DB.Query(query, "chef", email)
	if err != nil {
		log.Fatal("There is some problem in converting customer to chef", err)
	}
	fmt.Println("The task is done")

}

var menu_data []Food

func Get_menu() []Food {
	query := "SELECT * FROM food_menu"
	result, err := DB.Query(query)
	if err != nil {
		log.Fatal("There is some error in bringing the menu from the database : ", err)
	} else {
		for result.Next() {
			var food Food
			if err := result.Scan(&food.Food_id, &food.Name, &food.Desc, &food.Price, &food.Category_id); err != nil {
				log.Fatal(err)
			}
			menu_data = append(menu_data, food)
		}
	}
	return menu_data
}

func Find_free_chef() int {
	fmt.Println("Find free chef")
	role := "chef"
  	comp_lef  := "left"
	query := "SELECT * FROM user WHERE role = ? AND NOT EXISTS (SELECT *  FROM  order_table WHERE order_table.food_status = ? AND order_table.chef_id = user.user_id);"
	result,err := DB.Query(query , role , comp_lef)
	var user User
	if(err != nil){
		log.Fatal("There is some error in finding a free chef : " , err)
		return -1
	} else{
		if(!result.Next()){
			return -1
		} else{
			if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
				log.Fatal(err)
				}
		}
	}
	return user.Id
}

func Add_order_table(cutsomer int , status string , chef int) int{
	fmt.Println("Add order table")
	query := "INSERT INTO order_table (customer_id , food_status , chef_id) VALUES (?,?,?)"
	result,err := DB.Exec(query , cutsomer , status , chef)
	if(err != nil){
		log.Fatal("There is some error in adding data inside order table : " , err)
		return -1
	}else{
		fmt.Println("The data has been added successfully")
		id,err := result.LastInsertId()
		if(err != nil){
			log.Fatal("There is some error in getting the order id : " , err)
			return -1
		}else{
			return int(id) 
		}
	}
}
func Find_customer_id(email string) int {
	fmt.Println("Find customer id")
	role := "customer"
	query := "SELECT * FROM user WHERE role = ? AND email = ?"
	result,err := DB.Query(query , role , email)
	var user User
	if(err != nil){
		log.Fatal("There is some error in finding a customer : " , err)
		return -1
	} else{
		if(!result.Next()){
			return -1
		} else{
			if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
				log.Fatal(err)
				}
		}
	}
	return user.Id
}

func Find_email_id(email string) bool {
	fmt.Println("Find customer id")
	role := "customer"
	query := "SELECT * FROM user WHERE role = ? AND email = ?"
	result,err := DB.Query(query , role , email)
	if(err != nil){
		log.Fatal("There is some error in finding a customer : " , err)
		return false
	} else{
		if(!result.Next()){
			return false
		} else{
			return true
		}
	}
}

func Add_ordered_items(food_id int , quant int, instructions string , order_id int){
	fmt.Println("Add ordered items")
	food_status := "left";
    query := "INSERT INTO ordered_items VALUES (?,?,?,?,?)"
	_,err := DB.Query(query ,food_id, quant, instructions, order_id, food_status);
	if(err != nil){
		log.Fatal("There is some error in adding data inside the ordered items table : " , err)
	}else{
		fmt.Println("The data has been successfully added")
	}
}

type Food_added struct{
	Id int `json:"food_id"`
	Quant int `json:"quant"`
	Instruct string `json:"instructions"`
	Food_status string `json:"status"`
	Order_status int `json:"order_id"`
	Food_name string `json:"food_name"`
}

var food_slice []Food_added

func Get_orders(order_id int)([]Food_added){
	food_status := "left";
	query := `SELECT * FROM ordered_items WHERE order_id = ? AND food_status = ?`
    result,err := DB.Query(query, order_id, food_status);
	if(err != nil){
		log.Fatal("There is some error in getting data from the ordered_items : " , err)
		} else{
			for result.Next(){
				var food_item Food_added
				err := result.Scan(&food_item.Id , &food_item.Quant , &food_item.Instruct ,&food_item.Order_status, &food_item.Food_status)
				food_item.Food_name = Get_food_name(food_item.Id)
				if(err != nil){
					fmt.Println(err.Error())
					continue
				}
				food_slice = append(food_slice, food_item)
			}
		}
		return food_slice;
}

func Get_food_name(food_id int)string{
	query := `SELECT * FROM food_menu WHERE food_id = ?`
	result,err := DB.Query(query , food_id)
	var food_item Food
	if(err != nil){
		log.Fatal("There is some error getting the food name" , err)
	} else{
		for result.Next(){
			err := result.Scan(&food_item.Food_id , &food_item.Name , &food_item.Desc , &food_item.Price , &food_item.Category_id)
			if(err != nil){
				log.Fatal("There is some error in scanning the data" , err)
			}
		}
	}
  return food_item.Name;
}
