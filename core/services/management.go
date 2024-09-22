package services

import "money-manager/core/models"

func GetManagement() {

}

func AddManagement(dataManagement models.Management) models.Management {
	d.DB.Create(&dataManagement)
	return dataManagement
}

func UpdateManagement() {

}

func DeleteManagement() {

}
