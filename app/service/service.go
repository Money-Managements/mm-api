package service

import (
	"money-manager/core/services"

	"gorm.io/gorm"
)

func SetServices(db *gorm.DB) {
	services.SetServicesDependencies(services.Dependencies{
		DB: db,
	})
}
