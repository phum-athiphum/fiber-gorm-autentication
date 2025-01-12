package controllers

import (
	"strconv"

	"gorm-authentication/app/model"
	"gorm-authentication/app/queries"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	book, err := queries.GetBook(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.JSON(book)
}

func GetBooks(c *fiber.Ctx) error {
	book, err := queries.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var newBook model.Book

	// Parse the incoming request body
	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	validate := validator.New()
	// Validate the new book fields
	if err := validate.Struct(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(), // Return the validation error message
		})
	}

	// If the validation passes, proceed to create the book
	if err := queries.CreateBook(&newBook); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating book",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newBook)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	// Fetch the book by ID
	book, err := queries.GetBook(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	// Parse the incoming request body
	var updatedBook model.Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	validate := validator.New()

	// Validate the updated book fields
	if err := validate.Struct(&updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update the book in the database
	book.Name = updatedBook.Name
	book.Author = updatedBook.Author
	book.Description = updatedBook.Description
	book.Price = updatedBook.Price

	// If you have a query to update the book in the database, uncomment this:
	// if err := queries.UpdateBook(book); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "Error updating book",
	// 	})
	// }

	return c.JSON(book)
}
