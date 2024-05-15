package controllers

import (
	database "Product/Database"
	models "Product/Models"
	"Product/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	cat "category/Models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var PathImageProduct = "./Image"

func init() {
	if _, err := os.Stat(PathImageProduct); os.IsNotExist(err) {
		os.Mkdir(PathImageProduct, os.ModePerm)
	}
}

func getCategoryByID(categoryID int) (*cat.Category, error) {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8002/category/%d", categoryID))
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var category cat.Category
	if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &category, nil
}

func GetCatId(categoryID int) (*models.Category, error) {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8002/category/%d", categoryID))
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var cat models.Category
	if err := json.NewDecoder(resp.Body).Decode(&cat); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &cat, nil
}

func GetAllProduct(c *fiber.Ctx) error {
	var product []models.Product

	result := database.DB.Find(&product)

	if len(product) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": product,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	input := new(models.RequestProductCreate)

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

	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Image is required",
		})
	}

	CatId, err := strconv.Atoi(c.FormValue("category_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "CatId is required",
		})
	}

	filename := utils.GenerateImageFile(input.Name, image.Filename)

	if err := c.SaveFile(image, filepath.Join(PathImageProduct, filename)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't save file image",
		})
	}

	category, err := GetCatId(CatId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	category.Id = uint(CatId)

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Image:       filename,
		CategoryID:  uint(CatId),
	}

	result := database.DB.Create(&product)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't create product",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": product,
	})
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	err := database.DB.First(&product, "id = ?", id).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Product Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(models.RequestProductUpdate)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	var product models.Product

	err := database.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	if input.Name != "" {
		product.Name = input.Name
	}
	if input.Description != "" {
		product.Description = input.Description
	}
	if input.Price != "" {
		product.Price = input.Price
	}

	categoryId, err := strconv.Atoi(c.FormValue("category_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	if categoryId != 0 {
		category, err := getCategoryByID(categoryId)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "failed",
				"message": err.Error(),
			})
		}
		category.Id = uint(categoryId)
		product.CategoryID = uint(categoryId)
	}

	newImage, err := c.FormFile("image")
	if err == nil {
		if product.Image != "" {
			oldPath := filepath.Join(PathImageProduct, product.Image)
			os.Remove(oldPath)
		}

		newFileName := utils.GenerateImageFile(product.Name, newImage.Filename)
		if err := c.SaveFile(newImage, filepath.Join(PathImageProduct, newFileName)); err != nil {
			return err
		}

		product.Image = newFileName
	}

	result := database.DB.Save(&product)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't update product",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	err := database.DB.Debug().First(&product, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Product Not Found",
		})
	}

	if err := database.DB.Debug().Delete(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't Delete Product",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product deleted successfully!",
	})
}
