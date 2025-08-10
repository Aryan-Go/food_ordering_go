package routes

import (
	"github/aryan-go/food_ordering_go/package/controllers"
	"github/aryan-go/food_ordering_go/package/middlewares"

	"github.com/gorilla/mux"
)

func All_routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
	user := r.PathPrefix("/user").Subrouter()
	user.Use(middlewares.VerifyToken)
	user.HandleFunc("/auth_redirect", controllers.AuthRedirection).Methods("POST")
	// r.HandleFunc("/signup/{id}", controllers.Getiddata_signup).Methods("POST")
	// r.HandleFunc("/logout" , controllers.Logout_handler).Methods("GET")

	customer := r.PathPrefix("/customer").Subrouter()
	customer.Use(middlewares.JWTAuthMiddlewareCustomer)
	customer.HandleFunc("", controllers.CustomerHandler).Methods("GET")
	customer.HandleFunc("/menu_show", controllers.MenuHandler).Methods("GET")
	customer.HandleFunc("/cus_chef", controllers.CustomerChefConverter).Methods("GET")
	customer.HandleFunc("/render_waiting", controllers.GetOrderedItems).Methods("GET")
	customer.HandleFunc("/render_payment", controllers.PaymentHandler).Methods("GET")
	customer.HandleFunc("/food_items_added", controllers.FoodItemsAdded).Methods("POST")

	chef := r.PathPrefix("/chef").Subrouter()
	chef.Use(middlewares.JWTAuthMiddlewareChef)
	chef.HandleFunc("", controllers.ChefHandler).Methods("GET")
	chef.HandleFunc("/render_order", controllers.GetOrderedItems).Methods("GET")
	chef.HandleFunc("/complete_order", controllers.CompleteOrder).Methods("POST")

	admin := r.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/signup", controllers.GetUsersData).Methods("POST")
	admin.Use(middlewares.JWTAuthMiddlewareAdmin)
	admin.HandleFunc("", controllers.AdminHandler).Methods("GET")
	admin.HandleFunc("/admin_details", controllers.AdminDetails).Methods("GET")
	return r
}
