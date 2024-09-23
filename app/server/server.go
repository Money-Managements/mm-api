package server

import (
	"money-manager/app/middleware"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()

	middleware.SetCors(e)
	setRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
