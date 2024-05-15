package controllers

import (
	database "Customer/Database"
	"Customer/Models/entity"
	"Customer/Models/response"
	utils "Customer/Utils"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const PathImageCustomer = "./Public"

func init() {
	if _, err := os.Stat(PathImageCustomer); os.IsNotExist(err) {
		os.Mkdir(PathImageCustomer, os.ModePerm)
	}
}

func RegistrationCustomer(c *fiber.Ctx) error {
	input := new(response.RequestCustomerRegistration)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "Failed",
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

	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Image is required",
		})
	}

	filename := utils.GenerateImageFile(input.Username, image.Filename)

	if err := c.SaveFile(image, filepath.Join(PathImageCustomer, filename)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't save file image",
		})
	}

	customer := entity.Customer{
		FullName:    input.FullName,
		Username:    input.Username,
		Email:       input.Email,
		Password:    password,
		Phone:       input.Phone,
		Address:     input.Address,
		Gender:      input.Gender,
		DateOfBirth: input.DateOfBirth,
		Image:       filename,
	}

	result := database.DB.Debug().Create(&customer)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't to create cashier account",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer,
	})
}

func LoginCustomer(c *fiber.Ctx) error {
	input := new(response.RequestCustomerLogin)

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

	var customer entity.Customer

	result := database.DB.Debug().First(&customer, "username = ?", input.Username)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Username Not Found",
		})
	}

	checkPassword := utils.CheckPassword(input.Password, customer.Password)

	if !checkPassword {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Inccorrect Password",
		})
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(customer.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
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
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"token":   token,
		"message": customer,
	})

}

func GetProfile(c *fiber.Ctx) error {
	customer := c.Locals("customer").(entity.Customer)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer,
	})
}

func CustomerLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.Locals("customer", nil)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Logout Successfully",
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	input := new(response.RequestCustomerUpdate)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	customer := c.Locals("customer").(entity.Customer)
	err := database.DB.First(&customer, "id = ?", customer.Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	if input.FullName != "" {
		customer.FullName = input.FullName
	}

	if input.Username != "" {
		customer.Username = input.Username
	}

	if input.Email != "" {
		customer.Email = input.Email
	}

	if input.Phone != "" {
		customer.Phone = input.Phone
	}

	if input.Address != "" {
		customer.Address = input.Address
	}

	if input.Gender != "" {
		customer.Gender = input.Gender
	}

	if input.DateOfBirth != "" {
		customer.DateOfBirth = input.DateOfBirth
	}

	newImage, err := c.FormFile("image")
	if err == nil {
		if customer.Image != "" {
			oldPath := filepath.Join(PathImageCustomer, customer.Image)
			os.Remove(oldPath)
		}

		newFilename := utils.GenerateImageFile(customer.Username, newImage.Filename)
		if err := c.SaveFile(newImage, filepath.Join(PathImageCustomer, newFilename)); err != nil {
			return err
		}

		customer.Image = newFilename
	}

	result := database.DB.Debug().Save(&customer)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't update customer",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": customer,
	})
}

func ForgotPassword(c *fiber.Ctx) error {
	input := new(response.RequestCustomerForgotPassword)

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

	var customer entity.Customer

	if err := database.DB.Debug().First(&customer, "username = ?", input.Username).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Username Not Found",
		})
	}

	if err := database.DB.Debug().Where("phone = ?", input.Phone).First(&customer).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Phone Not Found",
		})
	}

	newPw, err := utils.GeneratePassword(input.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to generate password",
		})
	}

	if input.Password != "" {
		customer.Password = newPw
		if err := database.DB.Save(&customer).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status":  "failed",
				"message": "Failed to update password",
			})
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Password updated successfully",
		"data":    customer.Password,
	})
}

func EditPassword(c *fiber.Ctx) error {
	input := new(response.RequestCustomerEditPassword)

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

	customer := c.Locals("customer").(entity.Customer)
	err := database.DB.First(&customer, "id = ?", customer.Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	newPw, err := utils.GeneratePassword(input.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to generate password",
		})
	}

	if input.Password != "" {
		customer.Password = newPw
	}

	result := database.DB.Debug().Save(&customer)
	if result.Error != nil {
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"status":  "failed",
				"message": "Can't edit password",
			})
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Edit Password Successfully",
		"data":    customer.Password,
	})
}
