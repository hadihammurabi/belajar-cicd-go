package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"data": "halo semua!",
		})
	})
}
