package middleware

import (
	database "Vendor/Database"
	"Vendor/Models/entity"
	utils "Vendor/Utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
) 

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		tokenString := strings.Replace(cookie, "jwt", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.Secret_Key), nil
		})

		claims := token.Claims.(*jwt.StandardClaims)

		claims.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err = newToken.SignedString([]byte(utils.Secret_Key))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error refreshing token",
			})
		}

		cookie = tokenString
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    cookie,
			Expires:  time.Now().Add(time.Hour * 2),
			HTTPOnly: true,
		})

		var Vendor entity.Vendor
		database.DB.Where("id = ?", claims.Issuer).First(&Vendor)
		c.Locals("Vendor", Vendor)
		return c.Next()
	}
}
