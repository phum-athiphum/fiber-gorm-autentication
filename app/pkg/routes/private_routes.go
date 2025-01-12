package routes

import (
	"gorm-authentication/app/controllers"
	"gorm-authentication/app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(c *fiber.App) {
	// Create routes groups
	route := c.Group(("/api/v1"))

	route.Patch("/book", middleware.AuthMiddleware, controllers.UpdateBook)
}
