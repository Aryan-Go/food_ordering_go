package controllers

import (
	"encoding/json"
	"fmt"

	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/structures"

	// "log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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
		return
	}
	fmt.Println(order_id.Order_id)
	check_item := models.CompleteOrderItem(order_id.Order_id, "completed", order_id.Food_id)
	// check_item := models.Complete_order_item(6 , "left" , 1)
	check_order := models.CompleteOrder(order_id.Order_id)
	if check_item && !check_order {
		var succ structures.Error
		succ.Code = http.StatusAccepted
		succ.Message = "The order item has been completed"
		json.NewEncoder(w).Encode(succ)
		return
	}
	if check_order {
		var succ structures.Error
		succ.Code = http.StatusAccepted
		succ.Message = "The order has been completed"
		json.NewEncoder(w).Encode(succ)
		return
	}
}

func GetChefOrderedItems(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value("props")
	claims, ok := props.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid claims type", http.StatusInternalServerError)
		return
	}
	email := claims["email"].(string)
	fmt.Println(email)
	role := claims["role"].(string)
	if role != "chef" {
		var err2 structures.Error
		err2.Code = http.StatusBadRequest
		err2.Message = "Please login as chef"
		json.NewEncoder(w).Encode(err2)
		return
	}
	chef_id := models.FindChefId(email)
	fmt.Println(chef_id)
	incomp_order_id := models.FindChefOrders(chef_id)
	var food_slices []structures.Food_added
	for _, value := range incomp_order_id {
		fmt.Println(value)
		food_slices = append(food_slices, models.GetOrders(value)...)
	}
	if len(food_slices) == 0 {
		var err2 structures.Error
		err2.Code = http.StatusBadRequest
		err2.Message = "No items to be made"
		json.NewEncoder(w).Encode(err2)
		return
	}
	json.NewEncoder(w).Encode(food_slices)
}
