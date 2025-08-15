package routes

import (
	"github/aryan-go/food_ordering_go/package/controllers"
	"github/aryan-go/food_ordering_go/package/middlewares"

	"github.com/gorilla/mux"
)

func All_routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST", "OPTIONS")
	user := r.PathPrefix("/user").Subrouter()
	user.Use(middlewares.CorsMiddleware)
	user.Use(middlewares.VerifyToken)
	user.HandleFunc("/auth_redirect", controllers.AuthRedirection).Methods("POST", "OPTIONS")
	// r.HandleFunc("/signup/{id}", controllers.Getiddata_signup).Methods("POST")
	// r.HandleFunc("/logout" , controllers.Logout_handler).Methods("GET")

	customer := r.PathPrefix("/customer").Subrouter()
	customer.Use(middlewares.CorsMiddleware)
	customer.Use(middlewares.JWTAuthMiddlewareCustomer)
	customer.HandleFunc("", controllers.CustomerHandler).Methods("GET")
	customer.HandleFunc("/menu_show", controllers.MenuHandler).Methods("GET")
	customer.HandleFunc("/cus_chef", controllers.CustomerChefConverter).Methods("POST", "OPTIONS")
	customer.HandleFunc("/render_waiting", controllers.GetOrderedItems).Methods("POST")
	customer.HandleFunc("/render_payment", controllers.PaymentHandler).Methods("GET")
	customer.HandleFunc("/food_items_added", controllers.FoodItemsAdded).Methods("POST", "OPTIONS")

	chef := r.PathPrefix("/chef").Subrouter()
	chef.Use(middlewares.CorsMiddleware)
	chef.Use(middlewares.JWTAuthMiddlewareChef)
	chef.HandleFunc("", controllers.ChefHandler).Methods("GET")
	chef.HandleFunc("/render_order", controllers.GetChefOrderedItems).Methods("GET")
	chef.HandleFunc("/complete_order", controllers.CompleteOrder).Methods("POST", "OPTIONS")

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middlewares.CorsMiddleware)
	admin.Use(middlewares.JWTAuthMiddlewareAdmin)
	admin.HandleFunc("/signup", controllers.GetUsersData).Methods("GET")
	admin.HandleFunc("/admin_chef_conversion", controllers.AdminConvertChef).Methods("POST", "OPTIONS")
	admin.HandleFunc("/admin_payment_complete", controllers.AdminCompletePaymemt).Methods("POST", "OPTIONS")
	// admin.HandleFunc("/admin_chef_conversion", controllers.AdminCompleteItem).Methods("POST")
	admin.HandleFunc("", controllers.AdminHandler).Methods("GET")
	admin.HandleFunc("/admin_details", controllers.AdminDetails).Methods("GET")
	return r
}
