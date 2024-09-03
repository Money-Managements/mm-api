package account

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccounts(c echo.Context) error {
	return c.String(http.StatusOK, "account")
}
