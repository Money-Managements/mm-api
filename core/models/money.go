package models

import (
	"money-manager/internal/constant"

	"gorm.io/gorm"
)

type Money struct {
	gorm.Model
	Amount     constant.Amount `gorm:"default:0.00"`
	LocationID constant.UUID   `gorm:"not null"`
	Location   Location        `gorm:"not null"`
}

type MoneyAccount struct {
	gorm.Model
	AccountID constant.UUID `gorm:"not null"`
	MoneyID   constant.UUID `gorm:"not null"`
	Money     Money         `gorm:"not null"`
}

type MoneyTransaction struct {
	gorm.Model
	TransactionID constant.UUID `gorm:"not null"`
	MoneyID       constant.UUID `gorm:"not null"`
	Money         Money         `gorm:"not null"`
}
