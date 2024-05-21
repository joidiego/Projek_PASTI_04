package controllers

import (
	database "Order/Database"
	entity "Order/Models/Entity"
	response "Order/Models/Response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	cust "Backend/Customer/Models/entity"
	prod "Backend/Product/Models"


	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProduct(ProductID int) (*prod.Product, error) {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8006/product/%d", ProductID))
	if err != nil {
		return nil, fmt.Errorf("Failed to make HTTP Request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var product prod.Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &product, nil
}
func GetLoggedInCustomer(token string, id int) (*cust.Customer, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://127.0.0.1:8006/customer/profile/%d", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.AddCookie(&http.Cookie{Name: "jwt", Value: token})

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var customer cust.Customer
	if err := json.NewDecoder(resp.Body).Decode(&customer); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &customer, nil
}

func CreateOrder(c *fiber.Ctx) error {
	input := new(response.RequestOrderCreate)

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

	prodID, err := strconv.Atoi(c.FormValue("product_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Product Id is required",
		})
	}

	product, err := GetProduct(prodID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "Failed",
			"message": err.Error(),
		})
	}

	custId, err := strconv.Atoi(c.FormValue("customer_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "Failed",
			"message": err.Error(),
		})
	}
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	customer, err := GetLoggedInCustomer(cookie, custId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Failed",
			"message": err.Error(),
		})
	}

	customer.Id = uint(custId)

	product.Id = uint(prodID)

	orderID, err := OrderNumber()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to generate order number",
		})
	}

	// Create order entity
	order := entity.Order{
		OrderNumber: orderID,
		Quantity:    input.Quantity,
		ProductID:   product.Id,
		CustomerID:  uint(custId),
		Status:      "Waiting",
	}

	// Save order to database
	result := database.DB.Create(&order)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to create order",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": order,
	})
}

func OrderNumber() (string, error) {
	today := time.Now().Format("2006-01-02")

	var lastOrder entity.Order
	result := database.DB.Where("DATE(created_at) = ?", today).Order("id").First(&lastOrder)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return "", fmt.Errorf("failed to retrieve last order: %v", result.Error)
	}

	var orderNumber int
	if result.RowsAffected == 0 {
		// No orders for today, start from 1
		orderNumber = 1
	} else {
		// Increment the last order number
		lastOrderNumber, err := strconv.Atoi(lastOrder.OrderNumber)
		if err != nil {
			return "", fmt.Errorf("failed to parse last order number: %v", err)
		}
		orderNumber = lastOrderNumber + 1
	}

	return fmt.Sprintf("%d", orderNumber), nil
}
func GetAllOrder(c *fiber.Ctx) error {
	var order []entity.Order

	result := database.DB.Find(&order)

	if len(order) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Not Found",
		})
	}
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "Failed",
			"message": result.Error.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": order,
	})
}

func GetOrderById(c *fiber.Ctx) error {
	id := c.Params("id")

	var order entity.Order

	err := database.DB.First(&order, "id = ?", id).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "Failed",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "Failed",
		"message": order,
	})
}
