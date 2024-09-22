package models

import (
	"money-manager/internal/constant"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Description     string
	Type            constant.TransactionType `gorm:"not null"`
	ManagementID    constant.UUID            `gorm:"not null"`
	TargetAccountID constant.UUID            `gorm:"not null"`
	TargetAccount   Account                  `gorm:"not null"`
	OriginAccountID constant.UUID            `gorm:"not null"`
	OriginAccount   Account                  `gorm:"not null"`
}

// MoneyTransactionID constant.UUID            `gorm:"not null"`
// MoneyTransaction   MoneyTransaction         `gorm:"not null"`
