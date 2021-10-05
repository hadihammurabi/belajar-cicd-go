package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func NewRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{
			Message: "success",
			Data:    "halo epribadeh!",
		})
	})
}
