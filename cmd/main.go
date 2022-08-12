package main

import (
	"log"
	"net/http"

	"github.com/avkim12/L0/handlers"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", handlers.GetOrderById)
	log.Fatal(http.ListenAndServe(":8080", nil))
}