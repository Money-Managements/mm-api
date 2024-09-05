package location

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetLocations(c echo.Context) error {
	return c.String(http.StatusOK, "location")
}
