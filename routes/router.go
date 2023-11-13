package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sadek-39/echo-sql/controller"
)

func InitRoutes(e *echo.Echo) {
	//e.POST("/users", controller.SaveUser)
	e.GET("/users", controller.GetUser)
	//e.PUT("/users/:id", controller.UpdateUser)
	//e.DELETE("/users/:id", controller.DeleteUser)
}
