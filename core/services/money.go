package services

import (
	"money-manager/core/models"
	"money-manager/internal/constant"
)

type AddMoneyDTO struct {
	Amount     constant.Amount
	LocationID constant.UUID
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
	Amount     constant.Amount
	AccountID  constant.UUID
	LocationID constant.UUID
}

func AddMoneyAccount(addMoneyAccountDTO AddMoneyAccountDTO) models.MoneyAccount {
	moneyData := AddMoney(AddMoneyDTO{
		Amount:     addMoneyAccountDTO.Amount,
		LocationID: addMoneyAccountDTO.LocationID,
	})

	moneyAccountData := models.MoneyAccount{
		AccountID: addMoneyAccountDTO.AccountID,
		MoneyID:   constant.UUID(moneyData.ID),
	}

	d.DB.Create(&moneyAccountData)
	return moneyAccountData
}

// money transaction

type AddMoneyTransactionDTO struct {
	Amount        constant.Amount
	TransactionID constant.UUID
	LocationID    constant.UUID
}

func AddMoneyTransaction(addTransactionDTO AddMoneyTransactionDTO) models.MoneyTransaction {
	moneyData := AddMoney(AddMoneyDTO{
		Amount:     addTransactionDTO.Amount,
		LocationID: addTransactionDTO.LocationID,
	})

	moneyTransactionData := models.MoneyTransaction{
		TransactionID: addTransactionDTO.TransactionID,
		MoneyID:       constant.UUID(moneyData.ID),
	}

	d.DB.Create(&moneyTransactionData)
	return moneyTransactionData
}
