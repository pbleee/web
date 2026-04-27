package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c fiber.Ctx) error {
	header := c.Get("Authorization")

	if header == "" || len(header) < 7 {
		return c.Status(401).JSON(fiber.Map{
			"message": "belum login",
		})
	}
	tokenString := header[7:]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "token tidak valid",
		})
	}

	return c.Next()
}
