package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name         string `gorm:"not null"`
	ManagementID uint   `gorm:"not null"`
}
