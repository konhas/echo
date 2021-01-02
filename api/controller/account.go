package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create User")
}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get User")
}
