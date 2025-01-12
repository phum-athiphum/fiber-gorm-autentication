package routes

import (
	"gorm-authentication/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(c *fiber.App) {
	// Create routes groups
	route := c.Group(("/api/v1"))

	route.Get("/book", controllers.GetBook)
	route.Get("/books", controllers.GetBooks)
}
