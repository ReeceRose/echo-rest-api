package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/reecerose/echo-rest-api/controllers"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/api/v1/book", controllers.GetBooks)
	e.GET("/api/v1/book/:id", controllers.GetBook)
	e.POST("/api/v1/book", controllers.NewBook)
	e.DELETE("/api/v1/book/:id", controllers.DeleteBook)
}
