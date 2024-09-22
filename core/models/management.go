package models

import "gorm.io/gorm"

type Management struct {
	gorm.Model
	Name     string
	Accounts []Account
	// Transactions []Transaction
}
