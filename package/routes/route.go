package routes

import (
	"github/aryan-go/food_ordering_go/package/controllers"

	"github.com/gorilla/mux"
)

func All_routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
	r.HandleFunc("/signup", controllers.GetUsersData).Methods("POST")
	// r.HandleFunc("/signup/{id}", controllers.Getiddata_signup).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	// r.HandleFunc("/logout" , controllers.Logout_handler).Methods("GET")
	r.HandleFunc("/auth_redirect" , controllers.AuthRedirection).Methods("POST")
	r.HandleFunc("/customer" , controllers.CustomerHandler).Methods("GET")
	r.HandleFunc("/chef" , controllers.ChefHandler).Methods("GET")
	r.HandleFunc("/admin" , controllers.AdminHandler).Methods("GET")
	r.HandleFunc("/cus_chef" , controllers.CustomerChefConverter).Methods("GET")
	r.HandleFunc("/menu_show" , controllers.MenuHandler).Methods("GET")
	r.HandleFunc("/food_items_added" , controllers.FoodItemsAdded).Methods("POST")
	r.HandleFunc("/render_waiting" , controllers.GetOrderedItems).Methods("GET")
	r.HandleFunc("/render_order" , controllers.GetOrderedItems).Methods("GET")
	r.HandleFunc("/complete_order" , controllers.CompleteOrder).Methods("POST")
	r.HandleFunc("/render_payment" , controllers.PaymentHandler).Methods("GET")
	r.HandleFunc("/admin_details" , controllers.AdminDetails).Methods("GET")
	return r
}
