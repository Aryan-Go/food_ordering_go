package controllers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/structures"
	"log"

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

var data_add structures.Items_added

func FoodItemsAdded(w http.ResponseWriter, r *http.Request) {
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
	} else {
		if len(data_add.Items_added) > 9 || len(data_add.Special_instructions) > 9 || len(data_add.Id_arr) > 9 {
			var err structures.Error
			err.Code = http.StatusBadRequest
			err.Message = "Please put valid data only , there should be only 9 items"
			json.NewEncoder(w).Encode(err)
		} else {
			id := models.FindFreeChef()
			if id == -1 {
				var err structures.Error
				err.Code = http.StatusBadRequest
				err.Message = "There is no chef free at this moment"
				json.NewEncoder(w).Encode(err)
			} else {
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
				} else {
					fmt.Printf("%v %v", email, models.FindCustomerId(email))
					id := models.AddOrderTable(models.FindCustomerId(email), "left", id)
					total_payment := models.FindPayment(data_add.Items_added, data_add.Id_arr)
					models.AddPaymentDetails(total_payment, id, models.FindCustomerId(email))
					for key, value := range data_add.Items_added {
						if value != 0 {
							models.AddOrderedItems(data_add.Id_arr[key], data_add.Items_added[key], data_add.Special_instructions[key], id)
							var succ structures.Error
							succ.Code = http.StatusAccepted
							succ.Message = fmt.Sprintf("All the data has been added in the order table and ordered_items table successfully with order id %v", id)
							json.NewEncoder(w).Encode(succ)
						}
					}
				}
			}
		}
	}
}

func GetOrderedItems(w http.ResponseWriter, r *http.Request) {
	var order_id structures.Order_id
	err := json.NewDecoder(r.Body).Decode(&order_id)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is some error in getting the order id"
		json.NewEncoder(w).Encode(err)
	} else {
		fmt.Println(order_id.Id)
		food_slices := models.GetOrders(order_id.Id)
		if len(food_slices) == 0 {
			var err structures.Error
			err.Code = http.StatusBadRequest
			err.Message = "The order is completed or not yet made"
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(food_slices)
		}
	}
}
