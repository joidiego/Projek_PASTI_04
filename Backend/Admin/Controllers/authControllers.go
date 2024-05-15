package controllers

import (
	database "Admin/Database/seeders"
	"Admin/Models/entity"
	"Admin/Models/request"
	utils "Admin/Utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	
)

func LoginAdmin(c *fiber.Ctx) error {
	input := new(request.AdminRequestLogin)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	var admin entity.Admin

	result := database.DB.First(&admin, "username = ?", input.Username)

	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Username Not Found",
		})
	}

	checkPassword := utils.CheckPassword(input.Password, admin.Password)

	if !checkPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed",
			"message": "Inccorrect Password",
		})
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(admin.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	}

	tokens, err := utils.GenerateToken(&claims)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error Generating Token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokens,
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": admin,
		"token":   tokens,
	})
}

func GetProfile(c *fiber.Ctx) error {
	admin := c.Locals("admin").(entity.Admin)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": admin,
	})
}

func LogouAdmin(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.Locals("admin", nil)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Logout Successfully",
	})
}

func CreateCashier(c *fiber.Ctx) error {
	return nil
}
