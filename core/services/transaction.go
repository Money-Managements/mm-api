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
		transactionData.TargetAccountID = addTransactionDTO.TargetAccountID
		transactionData.OriginAccountID = addTransactionDTO.OriginAccountID
	}

	d.DB.Create(&transactionData)

	AddMoneyTransaction(AddMoneyTransactionDTO{
		Amount:        addTransactionDTO.Amount,
		TransactionID: constant.UUID(transactionData.ID),
		LocationID:    addTransactionDTO.LocationID,
	})

	return transactionData
}

//

type AddTransactionBaseDTO struct {
	Description  string
	ManagementID constant.UUID
	LocationID   constant.UUID
	Amount       constant.Amount
}

//

type AddTransactionIncomeDTO struct {
	AddTransactionBaseDTO
	AccountID constant.UUID
}

func AddTransactionIncome(addTransactionIncomeDTO AddTransactionIncomeDTO) models.Transaction {
	return AddTransaction(AddTransactionDTO{
		Description:     addTransactionIncomeDTO.Description,
		ManagementID:    addTransactionIncomeDTO.ManagementID,
		Amount:          addTransactionIncomeDTO.Amount,
		LocationID:      addTransactionIncomeDTO.LocationID,
		TargetAccountID: addTransactionIncomeDTO.AccountID,
		Type:            constant.TransactionTypeIncome,
	})
}

type AddTransactionSpendDTO struct {
	AddTransactionBaseDTO
	AccountID constant.UUID
}

func AddTransactionSpend(addTransactionSpendDTO AddTransactionSpendDTO) models.Transaction {
	return AddTransaction(AddTransactionDTO{
		Description:     addTransactionSpendDTO.Description,
		ManagementID:    addTransactionSpendDTO.ManagementID,
		Amount:          addTransactionSpendDTO.Amount,
		LocationID:      addTransactionSpendDTO.LocationID,
		OriginAccountID: addTransactionSpendDTO.AccountID,
		Type:            constant.TransactionTypeSpend,
	})
}

type AddTransactionTransferDTO struct {
	AddTransactionBaseDTO
	TargetAccountID constant.UUID
	OriginAccountID constant.UUID
}

func AddTransactionTransfer(addTransactionTransferDTO AddTransactionTransferDTO) models.Transaction {
	return AddTransaction(AddTransactionDTO{
		Description:     addTransactionTransferDTO.Description,
		ManagementID:    addTransactionTransferDTO.ManagementID,
		Amount:          addTransactionTransferDTO.Amount,
		LocationID:      addTransactionTransferDTO.LocationID,
		TargetAccountID: addTransactionTransferDTO.TargetAccountID,
		OriginAccountID: addTransactionTransferDTO.OriginAccountID,
		Type:            constant.TransactionTypeTransfer,
	})

}
