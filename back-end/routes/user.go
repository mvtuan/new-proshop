package routes

import (
	"back-end/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	e.POST("/users", controllers.CreateUser)
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetAUser)
	e.PUT("/users/:id", controllers.EditUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
