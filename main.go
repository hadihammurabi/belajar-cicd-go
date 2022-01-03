package main

import (
	// 	"simple-go/config"
	// 	"simple-go/route"

	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// 	config.InitDB()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"Message": "success",
			"Data":    "halo epriwan!",
		})
	})

	// 	route.NewRoute(e)
	panic(e.Start(":8080"))
}
