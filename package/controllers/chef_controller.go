package controllers

import (
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/middlewares"
	"github/aryan-go/food_ordering_go/package/models"
	// "log"
	"net/http"
	// "strconv"
	// "mux"
	// "golang.org/x/crypto/bcrypt"
)

func Chef_render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	if !state {
		fmt.Fprintf(w, "Your jwt token has expired please login again")
	} else if role != "chef" {
		fmt.Fprintf(w, "This is a protected route and you are not allowed")
	}else{
		fmt.Fprintf(w, "Welcome chef")
	}
}

type Com_item struct{
	Food_id int `json:"food_id"`
	Order_id int `json:"order_id"`
}

func Complete_order(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	if !state {
		fmt.Fprintf(w, "Your jwt token has expired please login again")
		return
	} else if role != "chef" {
		fmt.Fprintf(w, "This is a protected route and you are not allowed")
		return
	}else{
		var order_id Com_item
		err := json.NewDecoder(r.Body).Decode(&order_id)
		if(err != nil){
			fmt.Fprintf(w , "There is some error in the data sents : ")
			fmt.Println(err)
		}
		fmt.Println(order_id.Order_id , order_id.Food_id)
		if(err != nil){
			fmt.Fprintf(w , "There is some error in getting the order id")
		} else{
			fmt.Println(order_id.Order_id)
			check_item := models.Complete_order_item(order_id.Order_id , "completed" , order_id.Food_id)
			// check_item := models.Complete_order_item(6 , "left" , 1)
			check_order := models.Complete_order(order_id.Order_id)
			if(check_item && !check_order){
				fmt.Fprintf(w , "The order item has been completed")
			}
			if(check_order){
				fmt.Fprintf(w , "The order has been completed")
			}
		}
	}
}