package services

import (
	"money-manager/core/models"
	"money-manager/internal/constant"
)

type GetAccountFilter struct {
	ID           uint
	Name         string
	Type         constant.AccountType
	ManagementID constant.UUID
}

func GetAccount(getAccountFilter GetAccountFilter) models.Account {
	accountData := models.Account{}
	d.DB.Where(&getAccountFilter).First(&accountData)
	return accountData
}

type AddAccountDTO struct {
	Name         string
	Description  string
	Type         constant.AccountType
	ManagementID constant.UUID
	LocationID   constant.UUID
	Amount       constant.Amount
}

func AddAccount(addAccountDTO AddAccountDTO) models.Account {
	accountData := models.Account{
		Name:         addAccountDTO.Name,
		Description:  addAccountDTO.Description,
		Type:         addAccountDTO.Type,
		ManagementID: addAccountDTO.ManagementID,
	}
	d.DB.Create(&accountData)

	// locations

	locationData := models.Location{}

	if addAccountDTO.LocationID == 0 && addAccountDTO.Type == constant.AccountTypeBank {
		locationData = AddLocation(AddLocationDTO{
			Name:         addAccountDTO.Name,
			ManagementID: addAccountDTO.ManagementID,
		})
	}

	if addAccountDTO.LocationID != 0 && addAccountDTO.Type != constant.AccountTypeBank {
		locationData = GetLocation(GetLocationFilter{
			ID:           uint(addAccountDTO.LocationID),
			ManagementID: addAccountDTO.ManagementID,
		})
	}

	if addAccountDTO.LocationID == 0 && addAccountDTO.Type != constant.AccountTypeBank {
		locationData = GetLocation(GetLocationFilter{
			Name:         constant.LocationDefaultName,
			ManagementID: addAccountDTO.ManagementID,
		})
	}

	//  money account
	AddMoneyAccount(AddMoneyAccountDTO{
		Amount:     addAccountDTO.Amount,
		AccountID:  constant.UUID(accountData.ID),
		LocationID: constant.UUID(locationData.ID),
	})

	return accountData
}

func UpdateAccount() {

}

func DeleteAccount() {

}
