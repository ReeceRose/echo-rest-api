package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"github.com/reecerose/echo-rest-api/database"
	"github.com/reecerose/echo-rest-api/models"
	"github.com/reecerose/echo-rest-api/routes"
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

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
