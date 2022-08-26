package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	connection := "user=postgres dbname=products password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDB()

	allProduct, err := db.Query("Select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProduct.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
