package server

import (
	"money-manager/core/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	e := echo.New()

	services.SetServicesDependencies(services.Dependencies{
		DB: db,
	})

	setRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
