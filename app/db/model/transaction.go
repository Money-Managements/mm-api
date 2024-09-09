package model

import "time"

type Transaction struct {
	ID                 uint
	Description        string
	Type               int
	ManagementID       uint
	MoneyTransactionID uint
	MoneyTransaction   MoneyTransaction
	OriginAccountID    uint
	OriginAccount      Account
	TargetAccountID    uint
	TargetAccount      Account
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
