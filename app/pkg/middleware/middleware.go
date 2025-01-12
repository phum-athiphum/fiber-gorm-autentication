package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized. Missing Authoriztion header.",
		})
	}

	if !isValidToken(authHeader) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Token",
		})
	}

	return c.Next()
}

func isValidToken(token string) bool {
	expectedToken := "Bearer Test token"
	return token == expectedToken
}
