package model

import "time"

type Account struct {
	ID           uint
	Name         string
	Description  string
	Type         int
	ManagementID uint
	MoneyAccount []MoneyAccount
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
