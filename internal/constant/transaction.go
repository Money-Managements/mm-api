package constant

type TransactionType int

const (
	TransactionTypeIncome TransactionType = iota
	TransactionTypeSpend
	TransactionTypeTransfer
)
