package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"github.com/avkim12/L0/postgres"
	_ "github.com/lib/pq"
)

type Env struct {
	orders interface {
		CreateOrder(postgres.Order) error
		GetOrder(string) (postgres.Order, error)
	}
}

func main() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		orders: postgres.OrderDB{DB: db},
	}

	// cache := cache.New(5*time.Minute, 10*time.Minute)

	http.HandleFunc("/", env.HomePageHandler)
	http.HandleFunc("/get-order", env.GetOrderHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var templates = template.Must(template.ParseFiles("templates/homePage.html", "templates/orderPage.html"))

func (env *Env) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "homePage.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (env *Env) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := templates.ExecuteTemplate(w, "orderPage.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	uid := r.FormValue("orderID")

	order, err := env.orders.GetOrder(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	w.Write(order.Model)
}