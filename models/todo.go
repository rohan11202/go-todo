package models

import (
	"log"
	"time"

	"gorm.io/gorm"

	"p4nda/go-todo/config"
)

var (
	db *gorm.DB
)

type Todo struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"default:''"`
	Completed   bool      `gorm:"default:false"`
	Deadline    time.Time `gorm:"default:null"`
}

func Initiate() {

	config.ConnectDB()
	db = config.GetDB()
	err := db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal("failed to migrate db :", err)
	}
	log.Println("Database Migration Successful")
}
