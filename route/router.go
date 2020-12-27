package route

import (
	"net/http"

	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(mw.Logger())

	v1 := e.Group("/api/v1")
	{
		v1.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
	}
	return e
}
