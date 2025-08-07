package main

import (
	"fmt"
	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/routes"
	"log"
	"net/http"
	"time"
)

// ! This is the main file

func main() {
	fmt.Println("This is the main.go file")
	_, err := models.Init_database()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	server := routes.All_routes()

	srv := &http.Server{
		Handler:      server,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
