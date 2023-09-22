package main

import (
	"github.com/PhilipFelipe/golang-alura-course/database"
	"github.com/PhilipFelipe/golang-alura-course/routes"
)

func main() {
	database.DbConnect()
	routes.HandleRequest()
}
