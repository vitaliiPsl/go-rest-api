package database

import (
	"log"
	"os"

	"github.com/vitaliiPsl/go-rest-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := os.Getenv("DB_CONNECTION_STRING")

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Connected to the database")

	Db.AutoMigrate(&model.City{}, &model.Landmark{}, &model.Airport{})
	log.Println("Migrated models")
}
