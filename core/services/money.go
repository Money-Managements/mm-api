package services

import (
	"money-manager/core/models"
)

type AddMoneyDTO struct {
	Amount     models.Amount
	LocationID models.ID
}

func AddMoney(addMoneyDTO AddMoneyDTO) models.Money {
	moneyData := models.Money{
		Amount:     addMoneyDTO.Amount,
		LocationID: addMoneyDTO.LocationID,
	}

	d.DB.Create(&moneyData)
	return moneyData
}

// money account

type AddMoneyAccountDTO struct {
	Amount     models.Amount
	AccountID  models.ID
	LocationID models.ID
}

func AddMoneyAccount(addMoneyAccountDTO AddMoneyAccountDTO) models.MoneyAccount {
	moneyData := AddMoney(AddMoneyDTO{
		Amount:     addMoneyAccountDTO.Amount,
		LocationID: addMoneyAccountDTO.LocationID,
	})

	moneyAccountData := models.MoneyAccount{
		AccountID: addMoneyAccountDTO.AccountID,
		MoneyID:   moneyData.ID,
	}

	d.DB.Create(&moneyAccountData)
	return moneyAccountData
}

// money transaction

type AddMoneyTransactionDTO struct {
	Amount        models.Amount
	TransactionID models.ID
	LocationID    models.ID
}

func AddMoneyTransaction(addTransactionDTO AddMoneyTransactionDTO) models.MoneyTransaction {
	moneyData := AddMoney(AddMoneyDTO{
		Amount:     addTransactionDTO.Amount,
		LocationID: addTransactionDTO.LocationID,
	})

	moneyTransactionData := models.MoneyTransaction{
		TransactionID: addTransactionDTO.TransactionID,
		MoneyID:       moneyData.ID,
	}

	d.DB.Create(&moneyTransactionData)
	return moneyTransactionData
}
