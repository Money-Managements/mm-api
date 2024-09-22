package services

import (
	"money-manager/core/models"
	"money-manager/internal/constant"
)

func GetAccount() {
}

type AddAccountDTO struct {
	Name         string
	Description  string
	Type         constant.AccountType
	ManagementID uint
	LocationID   uint
	Amount       float64
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

	if addAccountDTO.LocationID != 0 && addAccountDTO.Type == constant.AccountTypeAssign {
		d.DB.First(&locationData, addAccountDTO.LocationID)
	}

	// money - money account

	moneyData := models.Money{
		Amount:     addAccountDTO.Amount,
		LocationID: locationData.ID,
	}
	d.DB.Create(&moneyData)

	moneyAccountData := models.MoneyAccount{
		AccountID: accountData.ID,
		MoneyID:   moneyData.ID,
	}
	d.DB.Create(&moneyAccountData)

	return accountData
}

func UpdateAccount() {

}

func DeleteAccount() {

}
