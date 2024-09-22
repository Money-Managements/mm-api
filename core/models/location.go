package models

import (
	"money-manager/internal/constant"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name         string        `gorm:"not null"`
	ManagementID constant.UUID `gorm:"not null"`
}
