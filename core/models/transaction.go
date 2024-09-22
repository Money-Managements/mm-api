package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Description        string
	Type               int
	ManagementID       uint
	MoneyTransactionID uint
	MoneyTransaction   MoneyTransaction
	OriginAccountID    uint
	OriginAccount      Account
	TargetAccountID    uint
	TargetAccount      Account
}
