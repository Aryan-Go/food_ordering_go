package main

import (
	"fmt"
	"github/aryan-go/food_ordering_go/package/routes"
	"log"
	"net/http"
	"time"
)

// ! This is the main file

func main() {
	fmt.Println("This is the main.go file")
	server := routes.All_routes()
	srv := &http.Server{
		Handler: server,
		Addr:    "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
