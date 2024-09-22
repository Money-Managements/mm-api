package services

import (
	"money-manager/core/models"
	"money-manager/internal/constant"
)

type AddTransactionDTO struct {
	Description     string
	Type            constant.TransactionType
	ManagementID    constant.UUID
	TargetAccountID constant.UUID
	OriginAccountID constant.UUID
	LocationID      constant.UUID
	Amount          constant.Amount
}

// TODO: improve this. One function for income and spend. And another different for transfers
// func AddTransactionIncome AddTransactionSpend func AddTransactonTransfer
func AddTransaction(addTransactionDTO AddTransactionDTO) models.Transaction {
	transactionData := models.Transaction{
		Description:  addTransactionDTO.Description,
		Type:         addTransactionDTO.Type,
		ManagementID: addTransactionDTO.ManagementID,
	}

	getAccountFilter := GetAccountFilter{
		ManagementID: addTransactionDTO.ManagementID,
	}

	if addTransactionDTO.Type == constant.TransactionTypeIncome {
		getAccountFilter.Type = constant.AccountTypeHiddenIncome
		accountIncomeData := GetAccount(getAccountFilter)

		transactionData.OriginAccountID = constant.UUID(accountIncomeData.ID)
		transactionData.TargetAccountID = addTransactionDTO.TargetAccountID
	}

	if addTransactionDTO.Type == constant.TransactionTypeSpend {
		getAccountFilter.Type = constant.AccountTypeHiddenSpend
		accountSpendData := GetAccount(getAccountFilter)

		transactionData.TargetAccountID = constant.UUID(accountSpendData.ID)
		transactionData.OriginAccountID = addTransactionDTO.OriginAccountID
	}

	if addTransactionDTO.Type == constant.TransactionTypeTransfer {

	}

	d.DB.Create(&transactionData)

	AddMoneyTransaction(AddMoneyTransactionDTO{
		Amount:        addTransactionDTO.Amount,
		TransactionID: constant.UUID(transactionData.ID),
		LocationID:    addTransactionDTO.LocationID,
	})

	return transactionData
}
