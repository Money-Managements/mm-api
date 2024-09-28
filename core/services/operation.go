package services

import (
	"money-manager/core/models"
)

type GetOperatonFilter struct {
	ID           models.ID
	Name         string
	ManagementID models.ID
}

func GetOperation(getOperatonFilter GetOperatonFilter) models.Operation {
	operationData := models.Operation{}
	d.DB.Where(&getOperatonFilter).First(&operationData)
	return operationData
}

type AddOperationDTO struct {
	Name         string
	Description  string
	ManagementID models.ID
}

func AddOperation(addOperationDTO AddOperationDTO) models.Operation {
	operationData := models.Operation{
		Name:         addOperationDTO.Name,
		Description:  addOperationDTO.Description,
		ManagementID: addOperationDTO.ManagementID,
	}

	d.DB.Create(&operationData)
	return operationData
}
