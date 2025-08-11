package controllers

import (
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/structures"

	// "log"
	"net/http"
	// "strconv"
	// "mux"
	// "golang.org/x/crypto/bcrypt"
)

func ChefHandler(w http.ResponseWriter, r *http.Request) {
	var succ structures.Error
	succ.Code = http.StatusAccepted
	succ.Message = "Welcome chef"
	json.NewEncoder(w).Encode(succ)
}

func CompleteOrder(w http.ResponseWriter, r *http.Request) {
	var order_id structures.Complete_item
	err := json.NewDecoder(r.Body).Decode(&order_id)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is some error in the data sents : "
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
	}
	fmt.Println(order_id.Order_id, order_id.Food_id)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is some error in getting the order id"
		json.NewEncoder(w).Encode(err)
	} else {
		fmt.Println(order_id.Order_id)
		check_item := models.CompleteOrderItem(order_id.Order_id, "completed", order_id.Food_id)
		// check_item := models.Complete_order_item(6 , "left" , 1)
		check_order := models.CompleteOrder(order_id.Order_id)
		if check_item && !check_order {
			var succ structures.Error
			succ.Code = http.StatusAccepted
			succ.Message = "The order item has been completed"
			json.NewEncoder(w).Encode(succ)
		}
		if check_order {
			var succ structures.Error
			succ.Code = http.StatusAccepted
			succ.Message = "The order has been completed"
			json.NewEncoder(w).Encode(succ)
		}
	}
}
