package database

import (
	"dijokiin/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DSN")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connecting to database err: ", err)
	}

	DB = conn
	conn.AutoMigrate(&models.Users{})
}
