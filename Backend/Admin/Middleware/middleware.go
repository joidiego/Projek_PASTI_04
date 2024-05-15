package middlewares

import (
	database "Admin/Database/seeders"
	"Admin/Models/entity"
	utils "Admin/Utils"
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

		claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
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
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		})

		var admin entity.Admin
		database.DB.Where("id = ?", claims.Issuer).First(&admin)

		c.Locals("admin", admin)

		return c.Next()
	}
}
