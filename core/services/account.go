package services

import (
	"money-manager/core/models"
)

type GetAccountFilter struct {
	ID           uint
	Name         string
	Type         models.AccountType
	ManagementID models.ID
}

func GetAccount(getAccountFilter GetAccountFilter) models.Account {
	accountData := models.Account{}
	d.DB.Where(&getAccountFilter).First(&accountData)
	return accountData
}

type AddAccountDTO struct {
	Name         string
	Description  string
	Type         models.AccountType
	ManagementID models.ID
	LocationID   models.ID
	Amount       models.Amount
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

	if addAccountDTO.LocationID == 0 && addAccountDTO.Type == models.AccountTypeBank {
		locationData = AddLocation(AddLocationDTO{
			Name:         addAccountDTO.Name,
			ManagementID: addAccountDTO.ManagementID,
		})
	}

	if addAccountDTO.LocationID != 0 && addAccountDTO.Type != models.AccountTypeBank {
		locationData = GetLocation(GetLocationFilter{
			ID:           uint(addAccountDTO.LocationID),
			ManagementID: addAccountDTO.ManagementID,
		})
	}

	if addAccountDTO.LocationID == 0 && addAccountDTO.Type != models.AccountTypeBank {
		locationData = GetLocation(GetLocationFilter{
			Name:         models.LocationDefaultName,
			ManagementID: addAccountDTO.ManagementID,
		})
	}

	//  money account
	AddMoneyAccount(AddMoneyAccountDTO{
		Amount:     addAccountDTO.Amount,
		AccountID:  accountData.ID,
		LocationID: locationData.ID,
	})

	return accountData
}

func UpdateAccount() {

}

func DeleteAccount() {

}
