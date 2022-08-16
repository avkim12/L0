package handlers

import (
	"net/http"
	"html/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "homePage.html", nil)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	uid := r.FormValue("orderID")
	data := struct{
		UID string
	}{
		UID: uid,
	}
	tmpl.ExecuteTemplate(w, "orderPage.html", data)
	
	// order, err := postgres.GetOrder(uid)
	// if err != nil {
	// 	log.Printf("No such order, err=%v \n", err)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(order.Model)
}
