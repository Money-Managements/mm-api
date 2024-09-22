package constant

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
