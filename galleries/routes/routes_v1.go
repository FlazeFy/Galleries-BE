package routes

import (
	"net/http"

	syshandlers "galleries/modules/systems/http_handlers"

	"github.com/labstack/echo"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Galleries")
	})

	// =============== Public routes ===============
	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)

	return e
}
