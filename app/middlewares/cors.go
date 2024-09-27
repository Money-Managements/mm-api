package middlewares

import (
	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

func SetCors(e *echo.Echo) {
	e.Use(md.CORSWithConfig(md.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
}
