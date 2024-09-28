package services

import (
	"money-manager/core/models"
)

type AddTransactionDTO struct {
	Description     string
	Type            models.TransactionType
	ManagementID    models.ID
	TargetAccountID models.ID
	OriginAccountID models.ID
	LocationID      models.ID
	OperationID     models.ID
	Amount          models.Amount
}

func AddTransaction(addTransactionDTO AddTransactionDTO) models.Transaction {
	transactionData := models.Transaction{
		Description:  addTransactionDTO.Description,
		Type:         addTransactionDTO.Type,
		ManagementID: addTransactionDTO.ManagementID,
		OperationID:  addTransactionDTO.OperationID,
	}

	getAccountFilter := GetAccountFilter{
		ManagementID: addTransactionDTO.ManagementID,
	}

	if addTransactionDTO.Type == models.TransactionTypeIncome {
		getAccountFilter.Type = models.AccountTypeHiddenIncome
		accountIncomeData := GetAccount(getAccountFilter)

		transactionData.OriginAccountID = accountIncomeData.ID
		transactionData.TargetAccountID = addTransactionDTO.TargetAccountID
	}

	if addTransactionDTO.Type == models.TransactionTypeSpend {
		getAccountFilter.Type = models.AccountTypeHiddenSpend
		accountSpendData := GetAccount(getAccountFilter)

		transactionData.TargetAccountID = accountSpendData.ID
		transactionData.OriginAccountID = addTransactionDTO.OriginAccountID
	}

	if addTransactionDTO.Type == models.TransactionTypeTransfer {
		transactionData.TargetAccountID = addTransactionDTO.TargetAccountID
		transactionData.OriginAccountID = addTransactionDTO.OriginAccountID
	}

	d.DB.Create(&transactionData)

	AddMoneyTransaction(AddMoneyTransactionDTO{
		Amount:        addTransactionDTO.Amount,
		TransactionID: transactionData.ID,
		LocationID:    addTransactionDTO.LocationID,
	})

	return transactionData
}

//

type AddTransactionBaseDTO struct {
	Description  string
	ManagementID models.ID
	LocationID   models.ID
	OperationID  models.ID
	Amount       models.Amount
}

//

type AddTransactionIncomeDTO struct {
	AddTransactionBaseDTO
	AccountID models.ID
}

func AddTransactionIncome(addTransactionIncomeDTO AddTransactionIncomeDTO) models.Transaction {
	return AddTransaction(AddTransactionDTO{
		Description:     addTransactionIncomeDTO.Description,
		ManagementID:    addTransactionIncomeDTO.ManagementID,
		Amount:          addTransactionIncomeDTO.Amount,
		LocationID:      addTransactionIncomeDTO.LocationID,
		TargetAccountID: addTransactionIncomeDTO.AccountID,
		OperationID:     addTransactionIncomeDTO.OperationID,
		Type:            models.TransactionTypeIncome,
	})
}

type AddTransactionSpendDTO struct {
	AddTransactionBaseDTO
	AccountID models.ID
}

func AddTransactionSpend(addTransactionSpendDTO AddTransactionSpendDTO) models.Transaction {
	return AddTransaction(AddTransactionDTO{
		Description:     addTransactionSpendDTO.Description,
		ManagementID:    addTransactionSpendDTO.ManagementID,
		Amount:          addTransactionSpendDTO.Amount,
		LocationID:      addTransactionSpendDTO.LocationID,
		OriginAccountID: addTransactionSpendDTO.AccountID,
		OperationID:     addTransactionSpendDTO.OperationID,
		Type:            models.TransactionTypeSpend,
	})
}

type AddTransactionTransferDTO struct {
	AddTransactionBaseDTO
	TargetAccountID models.ID
	OriginAccountID models.ID
}

func AddTransactionTransfer(addTransactionTransferDTO AddTransactionTransferDTO) models.Transaction {
	return AddTransaction(AddTransactionDTO{
		Description:     addTransactionTransferDTO.Description,
		ManagementID:    addTransactionTransferDTO.ManagementID,
		Amount:          addTransactionTransferDTO.Amount,
		LocationID:      addTransactionTransferDTO.LocationID,
		TargetAccountID: addTransactionTransferDTO.TargetAccountID,
		OriginAccountID: addTransactionTransferDTO.OriginAccountID,
		OperationID:     addTransactionTransferDTO.OperationID,
		Type:            models.TransactionTypeTransfer,
	})

}
