package services

import (
	"money-manager/core/models"
)

func GetManagement() {

}

func AddManagement(dataManagement models.Management) models.Management {
	d.DB.Create(&dataManagement)

	locationData := AddLocation(AddLocationDTO{
		Name:         models.LocationDefaultName,
		ManagementID: dataManagement.ID,
	})

	AddAccount(AddAccountDTO{
		Name:         models.AccountDefaultNameIncome,
		Type:         models.AccountTypeHiddenIncome,
		ManagementID: dataManagement.ID,
		Amount:       0,
		LocationID:   locationData.ID,
	})

	AddAccount(AddAccountDTO{
		Name:         models.AccountDefaultNameSpend,
		Type:         models.AccountTypeHiddenSpend,
		ManagementID: dataManagement.ID,
		Amount:       0,
		LocationID:   locationData.ID,
	})

	return dataManagement
}

func UpdateManagement() {

}

func DeleteManagement() {

}
