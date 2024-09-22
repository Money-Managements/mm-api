package models

import (
	"gorm.io/gorm"
)

type Money struct {
	gorm.Model
	Amount     float64  `gorm:"default:0.00"`
	LocationID uint     `gorm:"not null"`
	Location   Location `gorm:"not null"`
}

type MoneyAccount struct {
	gorm.Model
	AccountID uint  `gorm:"not null"`
	MoneyID   uint  `gorm:"not null"`
	Money     Money `gorm:"not null"`
}

type MoneyTransaction struct {
	gorm.Model
	MoneyID uint
	Money   Money
}
