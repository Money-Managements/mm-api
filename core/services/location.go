package services

import "money-manager/core/models"

func GetLocation() {

}

type AddLocationDTO struct {
	Name         string
	ManagementID uint
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
