package main

import (
	"fmt"

	"github.com/PhilipFelipe/golang-alura-course/database"
	"github.com/PhilipFelipe/golang-alura-course/models"
	"github.com/PhilipFelipe/golang-alura-course/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History of the person"},
		{Id: 2, Name: "Name 2", History: "History of the person"},
		{Id: 3, Name: "Name 3", History: "History of the person"},
	}
	database.DbConnect()
	fmt.Println("Starting server...")
	routes.HandleRequest()
}
