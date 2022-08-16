package main

import (
	"log"
	"net/http"

	"github.com/avkim12/L0/handlers"
	"github.com/avkim12/L0/postgres"
	_ "github.com/lib/pq"
)

func main() {

	err := postgres.Open()
	if err != nil {
		log.Println(err)
	}

	// cache := cache.New(5*time.Minute, 10*time.Minute)

	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/get-order", handlers.GetOrderHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
