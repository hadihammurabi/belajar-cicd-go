package main

import (
	"simple-go/config"
	"simple-go/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()

	e := echo.New()

	route.NewRoute(e)
	panic(e.Start(":8080"))
}
