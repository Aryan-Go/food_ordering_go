package routes

import (
	"github/aryan-go/food_ordering_go/package/controllers"

	"github.com/gorilla/mux"
)

func All_routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home_handler).Methods("GET")
	r.HandleFunc("/signup", controllers.Render_signup).Methods("GET")
	r.HandleFunc("/signup", controllers.Getdata_signup).Methods("POST")
	// r.HandleFunc("/signup/{id}", controllers.Getiddata_signup).Methods("POST")
	r.HandleFunc("/login", controllers.Render_login).Methods("GET")
	// r.HandleFunc("/logout" , controllers.Logout_handler).Methods("GET")
	r.HandleFunc("/auth_redirect" , controllers.Auth_redirection).Methods("POST")
	r.HandleFunc("/customer" , controllers.Customer_render).Methods("GET")
	r.HandleFunc("/chef" , controllers.Chef_render).Methods("GET")
	r.HandleFunc("/admin" , controllers.Admin_render).Methods("GET")
	r.HandleFunc("/cus_chef" , controllers.Customer_chef).Methods("GET")
	r.HandleFunc("/menu_show" , controllers.Menu_render).Methods("GET")
	r.HandleFunc("/food_items_added" , controllers.Foof_items_added).Methods("POST")
	r.HandleFunc("/render_waiting" , controllers.Get_ordered_items).Methods("GET")
	r.HandleFunc("/render_order" , controllers.Get_ordered_items).Methods("GET")
	return r
}
