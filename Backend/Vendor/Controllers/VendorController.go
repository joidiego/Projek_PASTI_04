package controllers

import (
	database "Vendor/Database"
	"Vendor/Models/dto"
	"Vendor/Models/entity"
	utils "Vendor/Utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func VendorRegister(c *fiber.Ctx) error {
	input := new(dto.RequestVendorRegister)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	password, err := utils.GeneratePassword(input.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to generate password",
		})
	}

	input.Password = password

	Vendor := entity.Vendor{
		Name:     input.Name,
		Phone:    input.Phone,
		Username: input.Username,
		Password: password,
	}

	result := database.DB.Create(&Vendor)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't to create cashier account",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": Vendor,
	})
}

func VendorLogin(c *fiber.Ctx) error {
	input := new(dto.RequestVendorLogin)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	var Vendor entity.Vendor

	result := database.DB.First(&Vendor, "username = ?", input.Username)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Username Not Found",
		})
	}

	checkpw := utils.CheckPassword(input.Password, Vendor.Password)

	if !checkpw {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Inccorrect Password",
		})
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(Vendor.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	}

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error Generating Token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": Vendor,
		"token":   token,
	})
}

func VendorGetProfile(c *fiber.Ctx) error {
	Vendor := c.Locals("cashier").(entity.Vendor)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": Vendor,
	})
}

func VendorLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.Locals("Vendor", nil)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Logout Successfully",
	})
}
