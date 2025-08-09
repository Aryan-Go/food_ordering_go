package controllers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/middlewares"
	"github/aryan-go/food_ordering_go/package/models"
	"log"

	// "log"
	"net/http"
	// "strconv"
	// "mux"
	// "golang.org/x/crypto/bcrypt"
)

func Customer_render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	fmt.Println(role)
	if !state {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	} else if role != "customer" {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "This is a protected route and you are not allowed"
		json.NewEncoder(w).Encode(err)
	} else {
		var succ Error
		succ.Code = http.StatusAccepted
		succ.Message = "Welcome customer"
		json.NewEncoder(w).Encode(succ)
	}
}

func Customer_chef(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	state, email, role := middlewares.Verify_token(jwtToken)
	if !state {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	} else if role != "customer" {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "This is a protected route and you are not allowed"
		json.NewEncoder(w).Encode(err)
	} else {
		models.Customer_to_chef(email)
		var succ Error
		succ.Code = http.StatusAccepted
		succ.Message = "The customer has been successfully turned into chef"
		json.NewEncoder(w).Encode(succ)
	}
}
func Menu_render(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	if !state {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	} else if role != "customer" {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "This is a protected route and you are not allowed"
		json.NewEncoder(w).Encode(err)
	} else {
		food_data := models.Get_menu()
		json.NewEncoder(w).Encode(food_data)
	}
}

type Items_added struct {
	Items_added          []int    `json:"item_add"`
	Special_instructions []string `json:"instructions"`
	Id_arr               []int    `json:"id"`
}

var data_add Items_added

func Foof_items_added(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	state, email, role := middlewares.Verify_token(jwtToken)
	if !state {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	} else if role != "customer" {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "This is a protected route and you are not allowed"
		json.NewEncoder(w).Encode(err)
	} else {
		err := json.NewDecoder(r.Body).Decode(&data_add)
		if err != nil {
			log.Fatal("There is some error in unmarshaling the added items data", err)
		} else {
			if len(data_add.Items_added) > 9 || len(data_add.Special_instructions) > 9 || len(data_add.Id_arr) > 9 {
				var err Error
				err.Code = http.StatusBadRequest
				err.Message = "Please put valid data only , there should be only 9 items"
				json.NewEncoder(w).Encode(err)
			} else {
				id := models.Find_free_chef()
				if id == -1 {
					var err Error
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
						var err Error
						err.Code = http.StatusBadRequest
						err.Message = "All items cannot be 0 please submit a valid input"
						json.NewEncoder(w).Encode(err)
					} else {
						fmt.Printf("%v %v", email, models.Find_customer_id(email))
						id := models.Add_order_table(models.Find_customer_id(email), "left", id)
						total_payment := models.Find_payment(data_add.Items_added, data_add.Id_arr)
						models.Add_payment_details(total_payment, id, models.Find_customer_id(email))
						for key, value := range data_add.Items_added {
							if value != 0 {
								models.Add_ordered_items(data_add.Id_arr[key], data_add.Items_added[key], data_add.Special_instructions[key], id)
								var succ Error
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
}

type Order_id struct {
	Id int `json:"id"`
}

func Get_ordered_items(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	var order_id Order_id
	if !state {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	} else if role != "customer" {
		var err Error
		err.Code = http.StatusBadRequest
		err.Message = "This is a protected route and you are not allowed"
		json.NewEncoder(w).Encode(err)
	} else {
		err := json.NewDecoder(r.Body).Decode(&order_id)
		if err != nil {
			var err Error
			err.Code = http.StatusBadRequest
			err.Message = "There is some error in getting the order id"
			json.NewEncoder(w).Encode(err)
		} else {
			fmt.Println(order_id.Id)
			food_slices := models.Get_orders(order_id.Id)
			if len(food_slices) == 0 {
				var err Error
				err.Code = http.StatusBadRequest
				err.Message = "The order is completed or not yet made"
				json.NewEncoder(w).Encode(err)
			} else {
				json.NewEncoder(w).Encode(food_slices)
			}
		}
	}
}
