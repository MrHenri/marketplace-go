package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/MrHenri/marketplace-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println(err.Error())
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println(err.Error())
		}

		err = models.CreateNewProduct(name, description, priceConvert, quantityConvert)
		if err != nil {
			panic(err.Error())
		}

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	err := models.DeleteProduct(productId)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product, err := models.GetProduct(productId)
	if err != nil {
		panic(err.Error())
	}

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvert, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err.Error())
		}

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println(err.Error())
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println(err.Error())
		}

		err = models.UpdateProduct(idConvert, name, description, priceConvert, quantityConvert)
		if err != nil {
			panic(err.Error())
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
