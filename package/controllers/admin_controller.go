package controllers

import (
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/middlewares"
	"github/aryan-go/food_ordering_go/package/models"
	"net/http"
)

type Com_item2 struct {
	Order_id int `json:"order_id"`
	Tip      int `json:"tip"`
}

type Payment_details struct {
	Final_payment float64
	Tip           int
}

func Render_payment(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	state, email, role := middlewares.Verify_token(jwtToken)
	if !state {
		fmt.Fprintf(w, "Your jwt token has expired please login again")
	} else if role != "customer" {
		fmt.Fprintf(w, "This is a protected route and you are not allowed")
	} else {
		var detials Com_item2
		err := json.NewDecoder(r.Body).Decode(&detials)
		if err != nil {
			fmt.Fprintf(w, "There is some error in getting data from the user")
			fmt.Println(err)
		} else {
			total_payment := models.Find_total_payment(detials.Order_id, models.Find_customer_id(email))
			final_payment := total_payment + ((total_payment * float64(detials.Tip)) / float64(100))
			var pay_details Payment_details
			pay_details.Final_payment = final_payment
			pay_details.Tip = detials.Tip
			json.NewEncoder(w).Encode(pay_details)
		}
	}
}

type Incomplete struct{
	Order_id_order []int
	Order_id_payment []int
}

func Render_admin(w http.ResponseWriter, r *http.Request){
jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	if !state {
		fmt.Fprintf(w, "Your jwt token has expired please login again")
	} else if role != "admin" {
		fmt.Fprintf(w, "This is a protected route and you are not allowed")
	} else {
		var details Incomplete
		details.Order_id_order = models.Incomplete_order_id()
		details.Order_id_payment = models.Unpaid_order_id()
		json.NewEncoder(w).Encode(details)
	}
}