package entity

import (
	"fmt"

	"github.com/PhilipFelipe/golang-alura-course/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.DbConnect()
	defer db.Close()

	allProducts, err := db.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
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
	return products
}

func GetProduct(productId string) Product {
	db := db.DbConnect()
	defer db.Close()

	product, err := db.Query("SELECT * FROM products WHERE id=$1", productId)
	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	var id, quantity int
	var name, description string
	var price float64

	for product.Next() {
		err = product.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
	}
	return p
}

func CreateProduct(name string, description string, price float64, quantity int) {
	db := db.DbConnect()
	defer db.Close()

	fmt.Println(name, description, quantity, price)
	insertNewProduct, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	_, err = insertNewProduct.Exec(name, description, price, quantity)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteProduct(id string) {
	db := db.DbConnect()
	defer db.Close()

	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	_, err = deleteProduct.Exec(id)
	if err != nil {
		panic(err.Error())
	}
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.DbConnect()
	defer db.Close()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	_, err = updateProduct.Exec(name, description, price, quantity, id)
	if err != nil {
		panic(err.Error())
	}
}
