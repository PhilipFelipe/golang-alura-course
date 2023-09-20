package main

import (
	"net/http"

	"github.com/PhilipFelipe/golang-alura-course/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
