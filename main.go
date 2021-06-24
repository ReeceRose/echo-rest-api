package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"github.com/reecerose/echo-rest-api/controllers"
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

	e.GET("/api/v1/book", controllers.GetBooks)
	e.GET("/api/v1/book/:id", controllers.GetBook)
	e.POST("/api/v1/book", controllers.NewBook)
	e.DELETE("/api/v1/book/:id", controllers.DeleteBook)

	e.Logger.Fatal(e.Start(":3000"))
}
