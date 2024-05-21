package controllers

import (
	database "Admin/Database"
	"Admin/Models/entity"
	"Admin/Models/request"
	utils "Admin/Utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	ord "service/order/Models/Entity"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
			"message": "Incorrect Password",
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

func GetOrder(c *fiber.Ctx) (*ord.Order, error) {
	OrderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8004/orders/%d", OrderId))
	if err != nil {
		return nil, fmt.Errorf("Failed to make HTTP Request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var order ord.Order
	if err := json.NewDecoder(resp.Body).Decode(&order); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	// Log the response for debugging
	fmt.Printf("Fetched Order: %+v\n", order)

	return &order, nil
}

func DataOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	ordId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "Failed",
			"message": fmt.Sprintf("Invalid order ID: %v", err),
		})
	}

	order, err := GetOrder(c)
	if err != nil {
		return err
	}

	order.Id = uint(ordId)

	var ord ord.Order

	// order, err := GetOrder(ordId)
	result := database.DBo.First(&ord, "id = ?", order.Id).Error
	if result != nil {
		return result
	}

	if ord.Status != "" {
		ord.Status = "Approved"
	}

	err = database.DBo.Save(&ord).Error
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Success",
		"order":  ord,
	})
}
