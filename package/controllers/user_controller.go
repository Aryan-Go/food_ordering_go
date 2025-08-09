package controllers

import (
	"encoding/json"
	"fmt"
	"github/aryan-go/food_ordering_go/package/middlewares"
	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/structures"
	"log"
	"net/http"
	"strconv"

	// "mux"
	"golang.org/x/crypto/bcrypt"
)

// ! here we will be writing the logic for routing fo the user





var users []structures.User



func Home_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are the the main link")
	fmt.Fprintln(w, "Home page")
}

func Render_signup(w http.ResponseWriter, r *http.Request) {
	var newUser structures.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		var errorAPi = structures.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorAPi)
	} else {
		if len(newUser.Name) == 0 || len(newUser.Email) == 0 || len(newUser.Password) == 0 || len(newUser.Repassword) == 0 || len(newUser.Role) == 0 {
			var errorAPi = structures.Error{
				Code:    http.StatusBadRequest,
				Message: "Your input is invalid or empty",
			}
			json.NewEncoder(w).Encode(errorAPi)
		} else {
			check := models.Find_email(newUser.Email)
			if !check {
				var errorAPi = structures.Error{
					Code:    http.StatusBadRequest,
					Message: "Email id already present please try to login",
				}
				json.NewEncoder(w).Encode(errorAPi)
			} else {
				if middlewares.Email_verification(newUser.Email) {
					if middlewares.Password_verification(newUser.Password) {
						if newUser.Password == newUser.Repassword {
							password := []byte(newUser.Password)
							hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
							password2 := []byte(newUser.Repassword)
							hashedPassword2, _ := bcrypt.GenerateFromPassword(password2, bcrypt.DefaultCost)
							if err != nil {
								log.Fatal("There is some error in encryption : ", err)
							}
							newUser.Password = string(hashedPassword)
							newUser.Repassword = string(hashedPassword2)
							users = append(users, newUser)
							models.Add_users(newUser.Email, newUser.Name, newUser.Password, newUser.Role)
							fmt.Fprint(w, "Data has been added successfully")
						} else {
							var errorAPi = structures.Error{
								Code:    http.StatusBadRequest,
								Message: "Your password and repassword is not matching please try again",
							}
							json.NewEncoder(w).Encode(errorAPi)
						}
					} else {
						var errorAPi = structures.Error{
							Code:    http.StatusBadRequest,
							Message: "Your password is not strong enough it must have special characters, numbers, upper case",
						}
						json.NewEncoder(w).Encode(errorAPi)
					}
				} else {
					var errorAPi = structures.Error{
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
	users := models.Get_all_users()
	if len(users) != 0 {
		err := json.NewEncoder(w).Encode(&users)
		if err != nil {
			var errorAPi = structures.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
			json.NewEncoder(w).Encode(errorAPi)
		}
	} else {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "There is no data of users right now"
		json.NewEncoder(w).Encode(err)
	}

}

func Render_login(w http.ResponseWriter, r *http.Request) {
	var loginUser structures.Login
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		var errorAPi = structures.Error{
			Code:    http.StatusBadRequest,
			Message: "There is some error with the email or password sent please send valid input for login",
		}
		json.NewEncoder(w).Encode(errorAPi)
	} else {
		var counter int = 0
		if !models.Find_email(loginUser.Email) {
			password, role := models.Find_password(loginUser.Email)
			fmt.Println(password, role, loginUser.Password)
			err = bcrypt.CompareHashAndPassword([]byte(password), []byte(loginUser.Password))
			if err == nil {
				jwtToken, err := middlewares.Create_token(loginUser.Email, role)
				if err != nil {
					var err structures.Error
					err.Code = http.StatusBadRequest
					err.Message = "There is some error in generating jwt token"
					json.NewEncoder(w).Encode(err)
				} else {
					var succ structures.Error
					succ.Code = http.StatusAccepted
					succ.Message = jwtToken
					json.NewEncoder(w).Encode(succ)
				}
			} else {
				var errorAPi = structures.Error{
					Code:    http.StatusForbidden,
					Message: "Your password is wrong for logging in please check once",
				}
				json.NewEncoder(w).Encode(errorAPi)
			}
		} else {
			if counter == len(users) {
				var errorAPi = structures.Error{
					Code:    http.StatusForbidden,
					Message: "Your email is wrong for logging in please check once",
				}
				json.NewEncoder(w).Encode(errorAPi)
			}
		}
	}
}

func Getiddata_signup(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
	num, errr := strconv.Atoi(id)
	fmt.Println(num)
	if errr != nil {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = errr.Error()
		json.NewEncoder(w).Encode(err)
	}
	user := models.Get_users_id(num)
	json.NewEncoder(w).Encode(&user)
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
			var err structures.Error
			err.Code = http.StatusBadRequest
			err.Message = "This is a protected route and you are not allowed"
			json.NewEncoder(w).Encode(err)
		}
	} else {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	}
}

func Admin_render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jwtToken := r.Header.Get("Authorization")
	state, _, role := middlewares.Verify_token(jwtToken)
	if !state {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "Your jwt token has expired please login again"
		json.NewEncoder(w).Encode(err)
	} else if role != "admin" {
		var err structures.Error
		err.Code = http.StatusBadRequest
		err.Message = "This is a protected route and you are not allowed"
		json.NewEncoder(w).Encode(err)
	} else {
		var succ structures.Error
		succ.Code = http.StatusBadRequest
		succ.Message = "Welcome admin"
		json.NewEncoder(w).Encode(succ)
	}
}
