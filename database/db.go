package database

import (
	"log"

	"github.com/danitrod/go-gin-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDB() {
	connectionString := "host=localhost port=5432 user=root dbname=root password=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.Student{})
}
