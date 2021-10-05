package main

import (
	"belajar-docker/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	route.NewRoute(e)
	e.Start(":8080")
}
