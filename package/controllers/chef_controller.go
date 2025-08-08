package controllers

import (
	// "encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/middlewares"
	// "github/aryan-go/food_ordering_go/package/models"
	// "log"
	"net/http"
	// "strconv"
	// "mux"
	// "golang.org/x/crypto/bcrypt"
)

func Chef_render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	if !state {
		fmt.Fprintf(w, "Your jwt token has expired please login again")
	} else if role != "chef" {
		fmt.Fprintf(w, "This is a protected route and you are not allowed")
	}else{
		fmt.Fprintf(w, "Welcome chef")
	}
}