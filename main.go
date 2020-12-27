package main

import (
	"github/konhas/echo-server/route"
)

func main() {
	router := route.Init()
	router.Start(":8081")
}
