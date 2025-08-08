package models

import (
	"log"
)

// ? Here we will create functions to crud datatabase for the user

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Id       int    `json:"id"`
}
var users_data []User

func Add_users(email string, name string, password string, role string) (int, error) {
	query := "INSERT INTO user (email,username,password,role) VALUES (?,?,?,?)"
	result, err := DB.Exec(query, email, name, password, role)
	if err != nil {
		log.Fatal("Some error in adding data of the user in the database : ", err)
		return 0, err
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal("Some error in adding data of the user in the database : ", err)
		}
		return int(id), nil
	}
}

func Get_all_users()([]User) {
	query := "SELECT * FROM user"
	result, err := DB.Query(query)
	if err != nil {
		log.Fatal("There is some error in database to get all the data", err)
	}
	for result.Next() {
		var user User
		if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
			log.Fatal(err)
		}
		users_data = append(users_data , user)
	}
	return users_data
}

func Get_users_id(id int)(User){
	query := "SELECT * FROM user WHERE user_id = (?)"
	result,err := DB.Query(query , id)
	var user User
	if(err != nil){
		log.Fatal("There is some error in finding this id or it doesn't exists")
		} else{
			for result.Next() {
				if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
					log.Fatal(err)
				}
			}
		}
		return user
	}
	
func Find_email(email string) bool {
	query := "SELECT * FROM user WHERE email = (?)"
	result,err := DB.Query(query , email)
	if(err != nil){
		log.Fatal("There is some error in email in login : " , err)
		return false
		} else{
			if !result.Next() {
				return false
				} else{
					return true
				}
			}
		}
		
func Find_password(email string) (string,string) {
	query := "SELECT * FROM user WHERE email = (?)"
	result,err := DB.Query(query , email)
	var user User
	if(err != nil){
		log.Fatal("There is some error in email in login : " , err)
		return "",""
		} else{
			for result.Next() {
				if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
					log.Fatal(err)
				}
			}
		}
	return user.Password,user.Role
}