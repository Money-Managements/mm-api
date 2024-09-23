package models

import (
	"money-manager/internal/constant"

	"gorm.io/gorm"
)

type Operation struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Description  string
	ManagementID constant.UUID `gorm:"not null"`
}
