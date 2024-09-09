package server

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	e := echo.New()
	setRoutes(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}
