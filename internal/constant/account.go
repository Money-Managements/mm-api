package constant

type AccountType int

const (
	AccountTypeBank AccountType = iota
	AccountTypeAssign
	AccountTypeHiddenIncome
)
