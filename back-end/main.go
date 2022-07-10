package main

import (
	"net/http"

	"back-end/configs"
	"back-end/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"data": "welcome"})
	})
	configs.ConnectDB()
	routes.UserRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
