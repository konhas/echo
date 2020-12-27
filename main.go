package main

import (
	"github.com/konhas/echo-server/route"
	"github.com/labstack/echo/v4/engine/fasthttp"
)

func main() {
	router := route.Init()
	router.Run(fasthttp.New(":8081"))
}
