package models

import (
	"money-manager/internal/constant"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Description  string
	Type         constant.AccountType `gorm:"not null"`
	ManagementID constant.UUID        `gorm:"not null"`
	MoneyAccount []MoneyAccount
}
