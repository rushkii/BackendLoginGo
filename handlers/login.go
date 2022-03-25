package handlers

import (
	db "dijokiin/database"
	"dijokiin/models"
	"dijokiin/utils.go"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const SECRET = "f1921d76639bc8f55325eb77dceb098d9914ef1382fd355b19652086a320b548"

func LoginUsers(c *fiber.Ctx) error {
	if cookie := c.Cookies("ses"); cookie != "" {
		return c.Status(200).JSON(models.ToJSON(200, "Already logged in", nil))
	}

	p := struct {
		Username string
		Email    string
		Password string
	}{}
	aday := time.Now().Add(time.Hour * 24)

	if err := c.BodyParser(&p); err != nil {
		log.Println("Register users payload err: ", err)
		return c.Status(400).JSON(models.ToJSON(400, "Invalid data request", nil))
	}

	var res *models.Users
	if err := db.DB.Where(
		"email = ? OR username = ? AND password = ?",
		p.Email,
		p.Username,
		utils.EncPwd(p.Password),
	).
		First(&res).Error; err != nil {
		return c.Status(404).JSON(models.ToJSON(404, "User not exists", nil))
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(res.ID)),
		ExpiresAt: aday.Unix(),
	})

	token, err := claims.SignedString([]byte(SECRET))
	if err != nil {
		return c.Status(500).JSON(models.ToJSON(500, "Couldn't login", nil))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "ses",
		Value:    token,
		Expires:  aday,
		HTTPOnly: true,
	})

	return c.JSON(models.ToJSON(200, "Success", nil))
}
