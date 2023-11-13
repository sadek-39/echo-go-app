package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sadek-39/echo-sql/config"
	"github.com/sadek-39/echo-sql/routes"
)

func main() {
	e := echo.New()

	config.NewDB()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	routes.InitRoutes(e)

	e.Start(":8080")
}
