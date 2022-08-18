package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/avkim12/L0/cache"
	"github.com/avkim12/L0/postgres"
	_ "github.com/lib/pq"
)

type Env struct {
	db    *postgres.OrderDB
	cache *cache.Cache
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
	order, found := env.cache.Get(uid)
	if !found {
		w.Write([]byte("The order ID specified does not exist"))
	} else {
		w.Write(order.Model)
	}
}

func (env *Env) Backup() {
	orders, err := env.db.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range orders {
		env.cache.Set(value.UID, value.Model)
	}
}


func main() {

	env := &Env{
		db:    postgres.New(),
		cache: cache.New(),
	}
	env.Backup()
	
	http.HandleFunc("/", env.HomePageHandler)
	http.HandleFunc("/get-order", env.GetOrderHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
