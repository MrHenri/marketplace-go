package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Camiseta", Description: "Azul", Price: 39.99, Quantity: 1},
		{Name: "Notebook", Description: "Cedil", Price: 1339.79, Quantity: 2},
		{Name: "Fone", Description: "Bom", Price: 29.79, Quantity: 1},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
