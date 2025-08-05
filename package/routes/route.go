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
	return r
}
