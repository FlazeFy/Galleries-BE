package httphandlers

import (
	"galleries/modules/stats/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetTotalDictionaryByType(c echo.Context) error {
	ord := c.Param("ord")
	view := "dictionaries_type"
	table := "dictionaries"

	result, err := repositories.GetTotalStats("api/v1/stats/dcttype/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalGalleryByFormat(c echo.Context) error {
	ord := c.Param("ord")
	view := "galleries_format"
	table := "galleries"

	result, err := repositories.GetTotalStats("api/v1/stats/galleryformat/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
