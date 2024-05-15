package controllers

import (
	database "Category/Database"
	models "Category/Models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(c *fiber.Ctx) error {
	var category []models.Category

	result := database.DB.Debug().Find(&category)

	if len(category) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Data Not Found",
		})
	}
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": category,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	input := new(models.RequestCategoryCreate)

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

	categories := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	err := database.DB.Create(&categories).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't create category",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": categories,
	})
}

func GetCategoryById(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Category

	err := database.DB.First(&category, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": category,
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(models.RequestCategoryUpdate)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	var category models.Category

	err := database.DB.First(&category, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	if input.Name != "" {
		category.Name = input.Name
	}
	if input.Description != "" {
		category.Description = input.Description
	}

	err = database.DB.Save(&category).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't to update category",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": category,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Category

	err := database.DB.Debug().First(&category, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	if err := database.DB.Debug().Delete(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error deleting category",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully deleted category!",
	})
}