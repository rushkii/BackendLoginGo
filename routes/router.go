package routes

import (
	h "dijokiin/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Use(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${ip}:${port} -> ${status} | ${method} | ${path}\n",
		TimeFormat: "Jan, 02 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Static("/static", "./assets")

	app.Post("/register", h.RegisterUsers)
	app.Post("/login", h.LoginUsers)
	app.Get("/me", h.GetMe)
}
