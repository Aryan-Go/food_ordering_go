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
	r.HandleFunc("/login", controllers.Render_login).Methods("GET")
	// r.HandleFunc("/logout" , controllers.Logout_handler).Methods("GET")
	r.HandleFunc("/auth_redirect" , controllers.Auth_redirection).Methods("POST")
	r.HandleFunc("/customer" , controllers.Customer_render).Methods("GET")
	r.HandleFunc("/chef" , controllers.Chef_render).Methods("GET")
	r.HandleFunc("/admin" , controllers.Admin_render).Methods("GET")
	return r
}
