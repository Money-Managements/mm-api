package services

import (
	"money-manager/core/models"
)

type GetLocationFilter struct {
	ID           uint
	Name         string
	ManagementID models.ID
}

func GetLocation(getLocationFilter GetLocationFilter) models.Location {
	locationData := models.Location{}
	d.DB.Where(&getLocationFilter).First(&locationData)
	return locationData
}

type AddLocationDTO struct {
	Name         string
	ManagementID models.ID
}

func AddLocation(addLocationDTO AddLocationDTO) models.Location {
	locationData := models.Location{
		Name:         addLocationDTO.Name,
		ManagementID: addLocationDTO.ManagementID,
	}

	d.DB.Create(&locationData)
	return locationData
}

func UpdateLocation() {

}

func DeleteLocation() {

}
