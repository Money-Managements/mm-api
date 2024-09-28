package models

type TransactionType int

const (
	TransactionTypeIncome TransactionType = iota
	TransactionTypeSpend
	TransactionTypeTransfer
)

type Transaction struct {
	Model
	Description     string
	Type            TransactionType `gorm:"not null"`
	ManagementID    ID              `gorm:"not null"`
	OperationID     ID              `gorm:"not null"`
	Operation       Operation       `gorm:"not null"`
	TargetAccountID ID              `gorm:"not null"`
	TargetAccount   Account         `gorm:"not null"`
	OriginAccountID ID              `gorm:"not null"`
	OriginAccount   Account         `gorm:"not null"`
}

// MoneyTransactionID constant.UUID            `gorm:"not null"`
// MoneyTransaction   MoneyTransaction         `gorm:"not null"`
