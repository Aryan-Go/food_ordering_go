package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github/aryan-go/food_ordering_go/package/middlewares"
	"github/aryan-go/food_ordering_go/package/models"
	"github/aryan-go/food_ordering_go/package/routes"
)

// ! This is the main file

func main() {
	fmt.Println("This is the main.go file")
	_, err := models.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	server := routes.All_routes()
	serverWithCors := middlewares.CorsMiddleware(server)
	srv := &http.Server{
		Handler:      serverWithCors,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
