package models

import (
	"fmt"
	"github/aryan-go/food_ordering_go/package/structures"
	"log"
)



func CustomerToChef(email string) {
	query := "UPDATE user SET role = (?) WHERE email = (?)"
	_, err := DB.Query(query, "chef", email)
	if err != nil {
		log.Fatal("There is some problem in converting customer to chef", err)
	}
	fmt.Println("The task is done")

}

var menu_data []structures.Food

func GetMenu() []structures.Food {
	query := "SELECT * FROM food_menu"
	result, err := DB.Query(query)
	if err != nil {
		log.Fatal("There is some error in bringing the menu from the database : ", err)
	} else {
		for result.Next() {
			var food structures.Food
			if err := result.Scan(&food.Food_id, &food.Name, &food.Desc, &food.Price, &food.Category_id); err != nil {
				log.Fatal(err)
			}
			menu_data = append(menu_data, food)
		}
	}
	defer result.Close()
	return menu_data
}

func FindFreeChef() int {
	fmt.Println("Find free chef")
	role := "chef"
  	comp_lef  := "left"
	query := "SELECT * FROM user WHERE role = ? AND NOT EXISTS (SELECT *  FROM  order_table WHERE order_table.food_status = ? AND order_table.chef_id = user.user_id);"
	result,err := DB.Query(query , role , comp_lef)
	var user structures.User2
	if(err != nil){
		log.Fatal("There is some error in finding a free chef : " , err)
		defer result.Close()
		return -1
	} else{
		if(!result.Next()){
			defer result.Close()
			return -1
		} else{
			if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
				log.Fatal(err)
				}
		}
	}
	defer result.Close()
	return user.Id
}

func AddOrderTable(cutsomer int , status string , chef int) int{
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
func FindCustomerId(email string) int {
	fmt.Println("Find customer id")
	role := "customer"
	query := "SELECT * FROM user WHERE role = ? AND email = ?"
	result,err := DB.Query(query , role , email)
	var user structures.User2
	if(err != nil){
		log.Fatal("There is some error in finding a customer : " , err)
		defer result.Close()
		return -1
	} else{
		if(!result.Next()){
			defer result.Close()
			return -1
		} else{
			if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
				defer result.Close()
				log.Fatal(err)
				}
		}
	}
	defer result.Close()
	return user.Id
}

func FindEmailId(email string) bool {
	fmt.Println("Find customer id")
	role := "customer"
	query := "SELECT * FROM user WHERE role = ? AND email = ?"
	result,err := DB.Query(query , role , email)
	if(err != nil){
		log.Fatal("There is some error in finding a customer : " , err)
		defer result.Close()
		return false
	} else{
		if(!result.Next()){
			defer result.Close()
			return false
		} else{
			defer result.Close()
			return true
		}
	}
}

func AddOrderedItems(food_id int , quant int, instructions string , order_id int){
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



var food_slice []structures.Food_added

func GetOrders(order_id int)([]structures.Food_added){
	food_status := "left";
	query := `SELECT * FROM ordered_items WHERE order_id = ? AND food_status = ?`
    result,err := DB.Query(query, order_id, food_status);
	if(err != nil){
		log.Fatal("There is some error in getting data from the ordered_items : " , err)
		defer result.Close()
		} else{
			for result.Next(){
				var food_item structures.Food_added
				err := result.Scan(&food_item.Id , &food_item.Quant , &food_item.Instruct ,&food_item.Order_status, &food_item.Food_status)
				food_item.Food_name = GetFoodName(food_item.Id)
				if(err != nil){
					fmt.Println(err.Error())
					continue
				}
				food_slice = append(food_slice, food_item)
			}
		}
		defer result.Close()
		return food_slice;
}

func GetFoodName(food_id int)string{
	query := `SELECT * FROM food_menu WHERE food_id = ?`
	result,err := DB.Query(query , food_id)
	var food_item structures.Food
	if(err != nil){
		defer result.Close()
		log.Fatal("There is some error getting the food name" , err)
	} else{
		for result.Next(){
			err := result.Scan(&food_item.Food_id , &food_item.Name , &food_item.Desc , &food_item.Price , &food_item.Category_id)
			if(err != nil){
				log.Fatal("There is some error in scanning the data" , err)
			}
		}
	}
	defer result.Close()
  return food_item.Name;
}

