package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {

	dsn := "host=localhost user=p4nda password=p4nda_pswd dbname=go-todo port=5432 sslmode=disable"
	//go-test-db
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}
	log.Println("Database connection established")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
