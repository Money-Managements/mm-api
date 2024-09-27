package middlewares

import (
	"money-manager/core/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

func SetHttpError(e *echo.Echo) {

	e.Use(md.RecoverWithConfig(md.RecoverConfig{
		DisablePrintStack: true,
	}))

	e.HTTPErrorHandler = HTTPError
}

func HTTPError(err error, c echo.Context) {
	errorResponse := helpers.GetErrorResponse(helpers.GetErrorResponseArg{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  helpers.InternalServer,
	})

	c.JSON(errorResponse.Code, errorResponse)
}
