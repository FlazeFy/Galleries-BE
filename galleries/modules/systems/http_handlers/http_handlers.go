package httphandlers

import (
	"galleries/modules/systems/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetDictionaryByType(c echo.Context) error {
	dctType := c.Param("type")
	result, err := repositories.GetDictionaryByType("api/v1/dct:"+dctType, dctType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetDictionaryByMyTag(c echo.Context) error {
	result, err := repositories.GetDictionaryByMyTag("api/v1/dct/tag/my")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelDictionaryById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelDictionaryById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostDictionary(c echo.Context) error {
	result, err := repositories.PostDictionary(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
