package models

import (
	"github.com/MrHenri/marketplace-go/db"
)

type Product struct {
	Name, Description string
	Price             float64
	Id, Quantity      int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()

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

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) error {
	db := db.ConnectDB()

	insertQuery, err := db.Prepare("insert into products (name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		return err
	}

	insertQuery.Exec(name, description, price, quantity)
	defer db.Close()

	return nil
}

func DeleteProduct(productId string) error {
	db := db.ConnectDB()

	deleteQuery, err := db.Prepare("delete from products Where id = $1")

	if err != nil {
		return err
	}

	deleteQuery.Exec(productId)

	defer db.Close()
	return nil
}

func GetProduct(productId string) (Product, error) {
	db := db.ConnectDB()

	getQuery := db.QueryRow("select * from products where id = $1", productId)

	var id, quantity int
	var name, description string
	var price float64

	if err := getQuery.Scan(&id, &name, &description, &price, &quantity); err != nil {
		return Product{}, err
	}

	product := Product{
		Id:          id,
		Quantity:    quantity,
		Name:        name,
		Description: description,
		Price:       price,
	}

	defer db.Close()
	return product, nil
}

func UpdateProduct(id int, name, description string, price float64, quantity int) error {
	db := db.ConnectDB()

	updateQuery, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		return err
	}

	updateQuery.Exec(&name, &description, &price, &quantity, &id)

	defer db.Close()
	return nil
}
