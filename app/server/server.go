package server

import (
	"money-manager/app/middlewares"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()

	middlewares.SetCors(e)
	middlewares.SetHttpError(e)
	setRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
