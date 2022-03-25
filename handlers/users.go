package handlers

import (
	db "dijokiin/database"
	"dijokiin/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetMe(c *fiber.Ctx) error {
	cookie := c.Cookies("ses")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return c.Status(401).JSON(models.ToJSON(401, "Not logged in", nil))
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var res *models.Users
	if err := db.DB.Where("id = ?", claims.Issuer).First(&res).Error; err != nil {
		log.Println("Unexpected error while get users: ", err)
		return c.Status(404).JSON(models.ToJSON(404, "User not exists", nil))
	}

	return c.JSON(models.ToJSON(200, "success", res))
}
