package httphandlers

import (
	"galleries/modules/galleries/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetGalleries(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	result, err := repositories.GetAllGalleries(page, 10, "api/v1/gallery")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetGalleryBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.GetGalleryBySlug("api/v1/gallery/"+slug, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelGalleryBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.HardDelGalleryBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostGallery(c echo.Context) error {
	result, err := repositories.PostGallery(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
