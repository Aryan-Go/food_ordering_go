package controllers

import (
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/middlewares"
	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/structures"
	"net/http"
)





func Render_payment(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	state, email, role := middlewares.Verify_token(jwtToken)
	if !state {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	} else if role != "customer" {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "This is a protected route and you are not allowed"
		json.NewEncoder(w).Encode(err)
	} else {
		var detials structures.Com_item2
		err := json.NewDecoder(r.Body).Decode(&detials)
		if err != nil {
			var err structures.Error
			err.Code = http.StatusBadRequest
			err.Message = "There is some error in getting data from the user"
			json.NewEncoder(w).Encode(err)
			fmt.Println(err)
		} else {
			total_payment := models.Find_total_payment(detials.Order_id, models.Find_customer_id(email))
			final_payment := total_payment + ((total_payment * float64(detials.Tip)) / float64(100))
			var pay_details structures.Payment_details
			pay_details.Final_payment = final_payment
			pay_details.Tip = detials.Tip
			models.Update_payment_table(detials.Order_id, models.Find_customer_id(email))
			json.NewEncoder(w).Encode(pay_details)
		}
	}
}



func Render_admin(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	if !state {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	} else if role != "admin" {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "This is a protected route and you are not allowed"
		json.NewEncoder(w).Encode(err)
	} else {
		var details structures.Incomplete
		details.Order_id_order = models.Incomplete_order_id()
		details.Order_id_payment = models.Unpaid_order_id()
		json.NewEncoder(w).Encode(details)
	}
}
