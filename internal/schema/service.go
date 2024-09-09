package schema

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServiceSetting struct {
	DB *gorm.DB
}

type ActionService func(c echo.Context)

type Service struct {
	Action ActionService
	Name   string
}
