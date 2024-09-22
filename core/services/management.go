package services

import (
	"money-manager/core/models"
	"money-manager/internal/constant"
)

func GetManagement() {

}

func AddManagement(dataManagement models.Management) models.Management {
	d.DB.Create(&dataManagement)

	locationData := AddLocation(AddLocationDTO{
		Name:         constant.LocationDefaultName,
		ManagementID: constant.UUID(dataManagement.ID),
	})

	AddAccount(AddAccountDTO{
		Name:         constant.AccountDefaultNameIncome,
		Type:         constant.AccountTypeHiddenIncome,
		ManagementID: constant.UUID(dataManagement.ID),
		Amount:       0,
		LocationID:   constant.UUID(locationData.ID),
	})

	AddAccount(AddAccountDTO{
		Name:         constant.AccountDefaultNameSpend,
		Type:         constant.AccountTypeHiddenSpend,
		ManagementID: constant.UUID(dataManagement.ID),
		Amount:       0,
		LocationID:   constant.UUID(locationData.ID),
	})

	return dataManagement
}

func UpdateManagement() {

}

func DeleteManagement() {

}
