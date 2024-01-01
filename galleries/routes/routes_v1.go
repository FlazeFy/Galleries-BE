package routes

import (
	"net/http"

	galhandlers "galleries/modules/galleries/http_handlers"
	stshandlers "galleries/modules/stats/http_handlers"
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
	e.GET("api/v1/dct/tag/my", syshandlers.GetDictionaryByMyTag)
	e.DELETE("api/v1/dct/destroy/:id", syshandlers.HardDelDictionaryById)
	e.POST("api/v1/dct", syshandlers.PostDictionary)

	// Gallery
	e.GET("api/v1/gallery", galhandlers.GetGalleries)
	e.GET("api/v1/gallery/:slug", galhandlers.GetGalleryBySlug)
	e.DELETE("api/v1/gallery/destroy/:slug", galhandlers.HardDelGalleryBySlug)
	e.POST("api/v1/gallery", galhandlers.PostGallery)

	// Stats
	e.GET("api/v1/stats/dcttype/:ord", stshandlers.GetTotalDictionaryByType)
	e.GET("api/v1/stats/galleryformat/:ord", stshandlers.GetTotalGalleryByFormat)

	return e
}
