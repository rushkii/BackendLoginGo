package handlers

import (
	db "dijokiin/database"
	"dijokiin/models"
	"dijokiin/utils.go"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterUsers(c *fiber.Ctx) error {
	var p models.Users

	if err := c.BodyParser(&p); err != nil {
		log.Println("Register users payload err: ", err)
		return c.Status(400).JSON(models.ToJSON(400, "Invalid data request", nil))
	}

	if p.Username == "" || p.Email == "" || p.Name == "" || p.Password == "" {
		log.Println("Users payload doesn't meet requirement")
		return c.Status(400).JSON(models.ToJSON(400, "Invalid data request", nil))
	}

	var exists bool

	cus := &models.Users{
		ID:        utils.RandInt(9999999, 999999999),
		Name:      p.Name,
		Username:  p.Username,
		Email:     p.Email,
		Password:  utils.EncPwd(p.Password),
		Picture:   "/static/momose-boy-1.png",
		CreatedAt: time.Now().Unix(),
	}

	db.DB.Model(&cus).
		Raw("SELECT EXISTS(SELECT * FROM users WHERE username = ? AND email = ?)", p.Username, p.Email).
		Scan(&exists)

	if !exists {
		db.DB.Create(&cus)
	} else {
		log.Println("Users record field already exist")
		return c.Status(409).JSON(models.ToJSON(409, "Maybe already exists", nil))
	}

	return c.JSON(models.ToJSON(200, "Success", nil))
}
