package model

type Management struct {
	ID           uint
	Name         string
	Accounts     []Account
	Transactions []Transaction
}
