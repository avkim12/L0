package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/avkim12/L0/cache"
	"github.com/avkim12/L0/postgres"
	_ "github.com/lib/pq"
)

type Env struct {
	cache  cache.Cache
	orders postgres.OrderDB
}

func main() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		cache:  *cache.New(),
		orders: postgres.OrderDB{DB: db},
	}

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
