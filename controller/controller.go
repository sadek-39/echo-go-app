package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sadek-39/echo-sql/service"
	"net/http"
)

func GetUser(c echo.Context) error {
	users, err := service.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)

}

func SaveUser() {

}

func DeleteUser() {

}

func UpdateUser() {

}
