package routes

import (
	"github/aryan-go/food_ordering_go/package/controllers"

	"github.com/gorilla/mux"
)

func All_routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/signup", controllers.RenderSignup).Methods("GET")
	r.HandleFunc("/signup", controllers.GetdataSignup).Methods("POST")
	// r.HandleFunc("/signup/{id}", controllers.Getiddata_signup).Methods("POST")
	r.HandleFunc("/login", controllers.RenderLogin).Methods("GET")
	// r.HandleFunc("/logout" , controllers.Logout_handler).Methods("GET")
	r.HandleFunc("/auth_redirect" , controllers.AuthRedirection).Methods("POST")
	r.HandleFunc("/customer" , controllers.CustomerRender).Methods("GET")
	r.HandleFunc("/chef" , controllers.ChefRender).Methods("GET")
	r.HandleFunc("/admin" , controllers.AdminRender).Methods("GET")
	r.HandleFunc("/cus_chef" , controllers.CustomerChef).Methods("GET")
	r.HandleFunc("/menu_show" , controllers.MenuRender).Methods("GET")
	r.HandleFunc("/food_items_added" , controllers.FoofItemsAdded).Methods("POST")
	r.HandleFunc("/render_waiting" , controllers.GetOrderedItems).Methods("GET")
	r.HandleFunc("/render_order" , controllers.GetOrderedItems).Methods("GET")
	r.HandleFunc("/complete_order" , controllers.CompleteOrder).Methods("POST")
	r.HandleFunc("/render_payment" , controllers.RenderPayment).Methods("GET")
	r.HandleFunc("/render_admin" , controllers.RenderAdmin).Methods("GET")
	return r
}
