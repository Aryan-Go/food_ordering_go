package controllers

import (
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/middlewares"
	"net/http"
)

// ! here we will be writing the logic for routing fo the user

type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
	Id         string `json:"id"`
	Role       string `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

type Error struct {
	Code    int    `json:"status_code"`
	Message string `json:"error"`
}

func Home_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are the the main link")
	fmt.Fprintln(w, "Home page")
}

func Render_signup(w http.ResponseWriter, r *http.Request) {
	var newUser User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		var errorAPi = Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorAPi)
	} else {
		if len(newUser.Name) == 0 || len(newUser.Email) == 0 || len(newUser.Password) == 0 || len(newUser.Repassword) == 0 || len(newUser.Id) == 0 || len(newUser.Role) == 0 {
			var errorAPi = Error{
				Code:    http.StatusBadRequest,
				Message: "Your input is invalid or empty",
			}
			json.NewEncoder(w).Encode(errorAPi)
		} else {
			var check bool
			for _, value := range users {
				if value.Email == newUser.Email {
					check = true
				}
			}
			if check {
				var errorAPi = Error{
					Code:    http.StatusBadRequest,
					Message: "Email id already present please try some other or login",
				}
				json.NewEncoder(w).Encode(errorAPi)
			} else {
				if middlewares.Email_verification(newUser.Email) {
					if middlewares.Password_verification(newUser.Password) {
						if newUser.Password == newUser.Repassword {
							users = append(users, newUser)
							fmt.Fprint(w, "Data has been added successfully")
						} else {
							var errorAPi = Error{
								Code:    http.StatusBadRequest,
								Message: "Your password and repassword is not matching please try again",
							}
							json.NewEncoder(w).Encode(errorAPi)
						}
					} else {
						var errorAPi = Error{
							Code:    http.StatusBadRequest,
							Message: "Your password is not strong enough it must have special characters, numbers, upper case",
						}
						json.NewEncoder(w).Encode(errorAPi)
					}
				} else {
					var errorAPi = Error{
						Code:    http.StatusBadRequest,
						Message: "Your email id is not valid",
					}
					json.NewEncoder(w).Encode(errorAPi)
				}
			}

		}
	}
}

func Getdata_signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(users) != 0 {
		err := json.NewEncoder(w).Encode(&users)
		if err != nil {
			var errorAPi = Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
			json.NewEncoder(w).Encode(errorAPi)
		}
	} else {
		fmt.Fprintf(w, "There is no data of users right now")
	}

}

func Render_login(w http.ResponseWriter, r *http.Request) {
	var loginUser Login
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		var errorAPi = Error{
			Code:    http.StatusBadRequest,
			Message: "There is some error with the email or password sent please send valid input for login",
		}
		json.NewEncoder(w).Encode(errorAPi)
	} else {
		for _, user := range users {
			if user.Email == loginUser.Email {
				if user.Password == loginUser.Password {
					jwtToken, err := middlewares.Create_token(user.Email, user.Role)
					if err != nil {
						fmt.Fprintf(w, "There is some error in generating jwt token")
					} else {
						fmt.Fprintf(w, "You are successfullt logged in %v", jwtToken)
					}
				} else {
					var errorAPi = Error{
						Code:    http.StatusForbidden,
						Message: "Your password is wrong for logging in please check once",
					}
					json.NewEncoder(w).Encode(errorAPi)
				}
			} else {
				var errorAPi = Error{
					Code:    http.StatusForbidden,
					Message: "Your email is wrong for logging in please check once",
				}
				json.NewEncoder(w).Encode(errorAPi)
			}
		}
	}
}

func Auth_redirection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	if state {
		if role == "customer" {
			http.Redirect(w, r, "/customer", http.StatusSeeOther)
		} else if role == "chef" {
			http.Redirect(w, r, "/chef", http.StatusSeeOther)
		} else if role == "admin" {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "The role is not valid please check once again")
		}
	} else {
		fmt.Fprintf(w, "Your jwt token has expired please login again")
	}
}

func Customer_render(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome customer")
}
func Chef_render(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Chef")
}
func Admin_render(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Admin")
}
