package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/PhilipFelipe/golang-alura-course/entity"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := entity.GetProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during 'price' conversion:", err)
		}
		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error during 'quantity' conversion:", err)
		}

		entity.CreateProduct(name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently) // STATUS 301
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	entity.DeleteProduct(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := entity.GetProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during 'price' conversion:", err)
		}
		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error during 'quantity' conversion:", err)
		}
		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during 'id' conversion:", err)
		}
		entity.UpdateProduct(convertedId, name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
