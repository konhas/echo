package main

import (
	"github/konhas/echo-server/api/route"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	route.Init(e)

	e.Start(":8081")
}
