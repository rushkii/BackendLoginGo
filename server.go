package main

import (
	"dijokiin/database"
	"dijokiin/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	app := fiber.New()
	routes.Use(app)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen("127.0.0.1:" + port))
}
