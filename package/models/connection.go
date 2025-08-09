package models

import (
	"database/sql"
	"fmt"
	backend "github/aryan-go/food_ordering_go"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ! Here I will make the connection means create the database and connect it
var DB *sql.DB

func Init_database() (*sql.DB, error) {
	config, err := backend.LoadConfig("../")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return nil, err
	} else {
		db_user := config.Db_user
		db_host := config.Db_host
		db_port := config.Db_port
		db_password := config.Db_password
		db_database := config.Db_database
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true",
			db_user, db_password, db_host, db_port, db_database)
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			return nil, fmt.Errorf("error opening database: %v", err)
		} else {
			DB.SetMaxOpenConns(50)
			DB.SetMaxIdleConns(10)
			DB.SetConnMaxLifetime(5 * time.Minute)
			err = DB.Ping()
			if err != nil {
				return nil, fmt.Errorf("error connecting to database: %v", err)
			} else {
				fmt.Println("The db is connected and working fine")
				return DB, nil
			}
		}
	}
}

func CloseDatabase() error {
	if DB != nil {
		fmt.Println("Closing database connection...")
		return DB.Close()
	}
	return nil
}
