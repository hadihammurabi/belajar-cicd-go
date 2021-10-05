package main

import (
	"simple-go/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	route.NewRoute(e)
	e.Start(":8080")
}
