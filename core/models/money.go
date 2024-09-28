package models

type Money struct {
	Model
	Amount     Amount   `gorm:"default:0.00"`
	LocationID ID       `gorm:"not null"`
	Location   Location `gorm:"not null"`
}

type MoneyAccount struct {
	Model
	AccountID ID    `gorm:"not null"`
	MoneyID   ID    `gorm:"not null"`
	Money     Money `gorm:"not null"`
}

type MoneyTransaction struct {
	Model
	TransactionID ID    `gorm:"not null"`
	MoneyID       ID    `gorm:"not null"`
	Money         Money `gorm:"not null"`
}
