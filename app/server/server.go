package server

import (
	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	setRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
