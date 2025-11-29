package controllers

import (
	"errors"
	"strconv"

	"github.com/akashigamedev/book-mgmt/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

var NewBook models.Book

func GetBook(c *fiber.Ctx) error {
	newBooks := models.GetAllBooks()
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"responseTxt": "success",
		"data":        newBooks,
	})
}

func GetBookById(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	ID, err := strconv.ParseUint(bookId, 10, 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"responseTxt": "invalid_id",
		})
	}

	book, err := models.GetBookById(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"responseTxt": "not_found",
				"data":        nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"responseTxt": "internal_server_error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"responseTxt": "success",
		"data":        book,
	})
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"responseTxt": "invalid_request_body",
		})
	}

	if err := book.CreateBook(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"responseTxt": "failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"responseTxt": "success",
		"data":        book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	ID, err := strconv.ParseUint(bookId, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"responseTxt": "invalid_id",
		})
	}
	if err := models.DeleteBook(ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"responseTxt": "not_found",
				"data":        nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"responseTxt": "internal_server_error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"responseTxt": "success",
	})
}

func UpdateBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	ID, err := strconv.ParseUint(bookId, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"responseTxt": "invalid_id",
		})
	}

	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"responseTxt": "invalid_request_body",
		})
	}

	newBook, err := models.GetBookById(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"responseTxt": "not_found",
				"data":        nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"responseTxt": "internal_server_error",
		})
	}

	if book.Name != "" {
		newBook.Name = book.Name
	}
	if book.Author != "" {
		newBook.Author = book.Author
	}
	if book.Publication != "" {
		newBook.Publication = book.Publication
	}

	if err := newBook.UpdateBook(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"responseTxt": "internal_server_error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"responseTxt": "success",
		"data":        newBook,
	})
}
