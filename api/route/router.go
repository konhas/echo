package route

import (
	"github/konhas/echo-server/api/controller"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	account := e.Group("/api/account")
	{
		account.GET("/users", controller.GetUser)
		account.POST("/users", controller.CreateUser)
	}
}
