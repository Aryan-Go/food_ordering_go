package controllers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/structures"
	"log"
	"strconv"

	// "log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	// "strconv"
	// "mux"
	// "golang.org/x/crypto/bcrypt"
)

var Customer_chef_arr []int

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var succ structures.Error
	succ.Code = http.StatusAccepted
	succ.Message = "Welcome customer"
	json.NewEncoder(w).Encode(succ)
}

func CustomerChefConverter(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value("props")
	claims, ok := props.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid claims type", http.StatusInternalServerError)
		return
	}
	email := claims["email"].(string)
	Customer_chef_arr = append(Customer_chef_arr, models.FindCustomerId(email))
	var succ structures.Error
	succ.Code = http.StatusAccepted
	succ.Message = "The customer's request has been successfully sent to admin to be converted into chef"
	json.NewEncoder(w).Encode(succ)
}
func MenuHandler(w http.ResponseWriter, r *http.Request) {
	food_data := models.GetMenu()
	json.NewEncoder(w).Encode(food_data)
}


func FoodItemsAdded(w http.ResponseWriter, r *http.Request) {
	var data_add structures.Items_added
	props := r.Context().Value("props")
	if props == nil {
		http.Error(w, "No claims found in context", http.StatusUnauthorized)
		return
	}
	claims, ok := props.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid claims type", http.StatusInternalServerError)
		return
	}
	email := claims["email"].(string)
	err := json.NewDecoder(r.Body).Decode(&data_add)
	if err != nil {
		log.Fatal("There is some error in unmarshaling the added items data", err)
		return
	}
	if len(data_add.Items_added) > 9 || len(data_add.Special_instructions) > 9 || len(data_add.Id_arr) > 9 {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "Please put valid data only , there should be only 9 items"
		json.NewEncoder(w).Encode(err)
		return
	}
	id := models.FindFreeChef()
	if id == -1 {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is no chef free at this moment"
		json.NewEncoder(w).Encode(err)
		return
	}
	counter := 0
	for _, value := range data_add.Items_added {
		if value == 0 {
			counter++
		}
	}
	if counter == len(data_add.Items_added) {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "All items cannot be 0 please submit a valid input"
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Printf("%v %v", email, models.FindCustomerId(email))
	id2 := models.AddOrderTable(models.FindCustomerId(email), "left", id)
	total_payment := models.FindPayment(data_add.Items_added, data_add.Id_arr)
	models.AddPaymentDetails(total_payment, id2, models.FindCustomerId(email))
	for key, value := range data_add.Items_added {
		if value != 0 {
			models.AddOrderedItems(data_add.Id_arr[key], data_add.Items_added[key], data_add.Special_instructions[key], id2)
		}
	}
	var succ structures.Error
	succ.Code = http.StatusAccepted
	succ.Message =  strconv.Itoa(id2)
	json.NewEncoder(w).Encode(succ)
}

func GetOrderedItems(w http.ResponseWriter, r *http.Request) {
	var order_id structures.Order_id
	err := json.NewDecoder(r.Body).Decode(&order_id)
	if err != nil {
		var err2 structures.Error
		err2.Code = http.StatusBadRequest
		err2.Message = "There is some error in getting the order id"
		fmt.Println(err)
		json.NewEncoder(w).Encode(err2)
		return
	}
	fmt.Println(order_id.Id)
	food_slices := models.GetOrders(order_id.Id)
	if len(food_slices) == 0 {
		var err2 structures.Error
		err2.Code = http.StatusBadRequest
		err2.Message = "The order is completed or not yet made"
		json.NewEncoder(w).Encode(err2)
		return
	}
	json.NewEncoder(w).Encode(food_slices)
}
