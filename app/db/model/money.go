package model

import "time"

type Money struct {
	ID         uint
	Amount     float64
	LocationID uint
	Location   Location
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type MoneyAccount struct {
	ID        uint
	AccountID uint
	MoneyID   uint
	Money     Money
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MoneyTransaction struct {
	ID        uint
	MoneyID   uint
	Money     Money
	CreatedAt time.Time
	UpdatedAt time.Time
}
