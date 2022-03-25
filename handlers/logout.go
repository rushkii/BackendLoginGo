package handlers

import (
	"dijokiin/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogoutUsers(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(models.ToJSON(200, "Success", nil))
}
