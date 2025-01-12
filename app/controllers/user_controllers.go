package controllers

import (
	"gorm-authentication/app/model"
	"gorm-authentication/app/queries"
	"gorm-authentication/app/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var newUser model.User

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	validate := validator.New()

	if err := validate.Struct(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	newUser.Password = hashedPassword

	if err := queries.CreateUser(&newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Created",
	})
}

func Login(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Created",
	})
}
