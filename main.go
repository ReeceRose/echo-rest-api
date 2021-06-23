package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"github.com/reecerose/echo-rest-api/database"
	"github.com/reecerose/echo-rest-api/models"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Datbase connection successfully opened")

	// Auto migrate DB
	database.DBConn.AutoMigrate(&models.Book{})
	fmt.Println("Database migrated")
}

func main() {
	e := echo.New()

	initDatabase()
	defer database.DBConn.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
