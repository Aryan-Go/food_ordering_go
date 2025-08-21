package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/structures"

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
	var detials structures.Complete_payment_item
	err := json.NewDecoder(r.Body).Decode(&detials)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is some error in getting data from the user"
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
		return
	}
	total_payment := models.FindTotalPayment(detials.Order_id, models.FindCustomerId(email))
	if total_payment == 0 {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is no payment details. Either this is not the related customer or order id is wrong"
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
		return
	}
	final_payment := total_payment + ((total_payment * float64(detials.Tip)) / float64(100))
	var pay_details structures.Payment_details
	pay_details.Final_payment = final_payment
	pay_details.Tip = detials.Tip
	json.NewEncoder(w).Encode(pay_details)
}

func CompletePayment(w http.ResponseWriter, r *http.Request) {
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
	var detials structures.Complete_payment_item
	err := json.NewDecoder(r.Body).Decode(&detials)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is some error in getting data from the user"
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
		return
	}
	total_payment := models.FindTotalPayment(detials.Order_id, models.FindCustomerId(email))
	final_payment := total_payment + ((total_payment * float64(detials.Tip)) / float64(100))
	var pay_details structures.Payment_details
	pay_details.Final_payment = final_payment
	pay_details.Tip = detials.Tip
	models.UpdatePaymentTable(detials.Order_id, models.FindCustomerId(email))
	json.NewEncoder(w).Encode(pay_details)
}

func AdminDetails(w http.ResponseWriter, r *http.Request) {
	var details structures.Incomplete
	details.Order_id_order = models.IncompleteOrderId()
	details.Payment_id = models.UnpaidPaymentId()
	details.Customer_chef_id = Customer_chef_arr
	json.NewEncoder(w).Encode(details)
}

func AdminConvertChef(w http.ResponseWriter, r *http.Request) {
	var info structures.Customer_id
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "Please send a valid input"
		json.NewEncoder(w).Encode(err)
		return
	}
	models.CustomerToChef(info.Id)
	var succ = structures.Error{
		Code:    http.StatusAccepted,
		Message: "The customer is successfully converted into chef",
	}
	json.NewEncoder(w).Encode(succ)
}

// func AdminDetails(w http.ResponseWriter, r *http.Request) {
// 	var details structures.Incomplete
// 	details.Order_id_order = models.IncompleteOrderId()
// 	details.Payment_id = models.UnpaidPaymentId()
// 	details.Customer_chef_id = Customer_chef_arr
// 	json.NewEncoder(w).Encode(details)
// }

func AdminConvertAdmin(w http.ResponseWriter, r *http.Request) {
	var info structures.Customer_id
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "Please send a valid input"
		json.NewEncoder(w).Encode(err)
		return
	}
	models.CustomerToAdmin(info.Id)
	var succ = structures.Error{
		Code:    http.StatusAccepted,
		Message: "The customer is successfully converted into admin",
	}
	json.NewEncoder(w).Encode(succ)
}

// func AdminCompletePaymemt(w http.ResponseWriter, r *http.Request) {
// 	var info structures.Order_id
// 	err := json.NewDecoder(r.Body).Decode(&info)
// 	if err != nil {
// 		var err structures.Error
// 		err.Code = http.StatusBadRequest
// 		err.Message = "Please send a valid input"
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	var details structures.Incomplete
// 	details.Payment_id = models.UnpaidPaymentId()
// 	for _, value := range details.Payment_id {
// 		if info.Id == value {
// 			models.UpdatePaymentId(value)
// 			total_payment, order_id := models.GetPaymentId(value)
// 			var details structures.Payment_details_admin
// 			details.Final_payment = total_payment
// 			details.Order_id = order_id
// 			json.NewEncoder(w).Encode(details)
// 			return

// 		}
// 	}
// 	var err2 structures.Error
// 	err2.Code = http.StatusBadRequest
// 	err2.Message = "This is the incorrect the payment id or it is paid"
// 	json.NewEncoder(w).Encode(err2)
// }

func MenuEdit(w http.ResponseWriter, r *http.Request) {
	var info structures.Food
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "Please send a valid input"
		json.NewEncoder(w).Encode(err)
		return
	}
	models.EditMenu(info.Name , info.Desc , info.Price , info.Category_id)
	var succ structures.Error
	succ.Code = http.StatusAccepted
	succ.Message = "The data has been successfully added in menu"
	json.NewEncoder(w).Encode(succ)
}