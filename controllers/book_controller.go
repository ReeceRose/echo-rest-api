package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reecerose/echo-rest-api/database"
	"github.com/reecerose/echo-rest-api/models"
)

func GetBooks(c echo.Context) error {
	db := database.DBConn

	var books []models.Book
	db.Find(&books)
	return c.JSON(http.StatusOK, books)
}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	db := database.DBConn

	var book models.Book
	db.Find(&book, id)

	return c.JSON(http.StatusOK, book)
}

func NewBook(c echo.Context) error {
	db := database.DBConn

	book := new(models.Book)

	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, models.Map{"message": err.Error()})
	}

	db.Create(&book)
	return c.JSON(http.StatusCreated, book)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	db := database.DBConn

	var book models.Book
	db.First(&book, id)

	if book.Title == "" {
		return c.JSON(http.StatusBadRequest, models.Map{"message": "No book found with given ID"})
	}

	db.Delete(&book)
	return c.JSON(http.StatusOK, models.Map{"message": "Book successfully deleted"})
}
