package routes

import (
	"net/http"

	galhandlers "galleries/modules/galleries/http_handlers"
	syshandlers "galleries/modules/systems/http_handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitV1() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Galleries")
	})

	// =============== Public routes ===============
	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)

	// Gallery
	e.GET("api/v1/gallery", galhandlers.GetGalleries)
	e.GET("api/v1/gallery/:slug", galhandlers.GetGalleryBySlug)

	return e
}
