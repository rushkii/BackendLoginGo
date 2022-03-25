package handlers

import (
	"dijokiin/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogoutUsers(c *fiber.Ctx) error {
	if c.Cookies("ses") == "" {
		return c.JSON(models.ToJSON(200, "Already logout", nil))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "ses",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(models.ToJSON(200, "Success", nil))
}
