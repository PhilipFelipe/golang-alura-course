package database

import (
	"log"

	"github.com/PhilipFelipe/golang-alura-course/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	connStr := "host=localhost user=testuser password=qwerty dbname=testdb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Panic("Error connecting to the dabatase")
	}
	DB.AutoMigrate(&models.Student{})
}
