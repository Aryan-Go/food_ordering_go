package controllers

import (
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/structures"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
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
	var detials structures.Com_item2
	err := json.NewDecoder(r.Body).Decode(&detials)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is some error in getting data from the user"
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
	} else {
		total_payment := models.FindTotalPayment(detials.Order_id, models.FindCustomerId(email))
		final_payment := total_payment + ((total_payment * float64(detials.Tip)) / float64(100))
		var pay_details structures.Payment_details
		pay_details.Final_payment = final_payment
		pay_details.Tip = detials.Tip
		models.UpdatePaymentTable(detials.Order_id, models.FindCustomerId(email))
		json.NewEncoder(w).Encode(pay_details)
	}
}

func AdminDetails(w http.ResponseWriter, r *http.Request) {
	var details structures.Incomplete
	details.Order_id_order = models.IncompleteOrderId()
	details.Order_id_payment = models.UnpaidOrderId()
	json.NewEncoder(w).Encode(details)
}
