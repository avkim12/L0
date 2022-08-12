package handlers

import (
	"net/http"
	"text/template"

	"github.com/avkim12/L0/model"
	"github.com/avkim12/L0/postgres"
)

// var tmpl = template.Must(template.ParseFiles("../templates/index.html"))

func GetOrderById(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("../templates/index.html")
	tmpl.Execute(w, nil)

	db := postgres.OpenConnection()

	row := db.QueryRow("SELECT model FROM models WHERE id = %s", r.Body)

	var order model.Order
	row.Scan(&order.Id, &order.Model)
	// orderInfo, _ := json.Marshal(order)

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(orderInfo)

	defer db.Close()
}