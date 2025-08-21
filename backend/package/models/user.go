package models

import (
	"log"

	"github/aryan-go/food_ordering_go/package/structures"
)

// ? Here we will create functions to crud datatabase for the user



func AddUsers(email string, name string, password string, role string) (int, error) {
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

func GetAllUsers() []structures.Get_user {
	var users_data []structures.Get_user
	query := "SELECT * FROM user"
	result, err := DB.Query(query)
	if err != nil {
		log.Fatal("There is some error in database to get all the data", err)
	}
	defer result.Close()
	for result.Next() {
		var user structures.Get_user
		if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
			log.Fatal(err)
		}
		users_data = append(users_data, user)
	}
	return users_data
}

func GetUsersId(id int) structures.Get_user {
	query := "SELECT * FROM user WHERE user_id = (?)"
	result, err := DB.Query(query, id)
	var user structures.Get_user
	defer result.Close()
	if err != nil {
		log.Fatal("There is some error in finding this id or it doesn't exists")
	} else {
		for result.Next() {
			if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
				log.Fatal(err)
			}
		}
	}
	return user
}

func FindEmail(email string) bool {
	query := "SELECT * FROM user WHERE email = (?)"
	result, err := DB.Query(query, email)
	if err != nil {
		log.Fatal("There is some error in email in login : ", err)
		defer result.Close()
		return false
	} else {
		if result.Next() {
			defer result.Close()
			return true
		} else {
			defer result.Close()
			return false
		}
	}
}

func FindPassword(email string) (string, string) {
	query := "SELECT * FROM user WHERE email = (?)"
	result, err := DB.Query(query, email)
	var user structures.Get_user
	if err != nil {
		log.Fatal("There is some error in email in login : ", err)
		defer result.Close()
		return "", ""
	} else {
		for result.Next() {
			if err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
				log.Fatal(err)
			}
		}
	}
	defer result.Close()
	return user.Password, user.Role
}
