package models

type AccountType int

const (
	AccountTypeBank AccountType = iota
	AccountTypeAssign
	AccountTypeHiddenIncome
	AccountTypeHiddenSpend
)

const (
	AccountDefaultNameIncome = "income"
	AccountDefaultNameSpend  = "spend"
)

type Account struct {
	Model
	Name         string `gorm:"not null"`
	Description  string
	Type         AccountType `gorm:"not null"`
	ManagementID ID          `gorm:"not null"`
	MoneyAccount []MoneyAccount
}
