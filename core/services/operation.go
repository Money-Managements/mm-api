package services

import (
	"money-manager/core/models"
	"money-manager/internal/constant"
)

type GetOperatonFilter struct {
	ID           constant.UUID
	Name         string
	ManagementID constant.UUID
}

func GetOperation(getOperatonFilter GetOperatonFilter) models.Operation {
	operationData := models.Operation{}
	d.DB.Where(&getOperatonFilter).First(&operationData)
	return operationData
}

type AddOperationDTO struct {
	Name         string
	Description  string
	ManagementID constant.UUID
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
