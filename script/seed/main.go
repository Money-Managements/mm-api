package main

import (
	"money-manager/app/db"
	"money-manager/core/models"
	"money-manager/core/services"
	"money-manager/internal/constant"
)

func main() {
	dataBase := db.ConnectDB()
	db.DropTables(dataBase)
	db.AutoMigrations(dataBase)

	services.SetServicesDependencies(services.Dependencies{
		DB: dataBase,
	})

	managements := seedManager()
	seedAccount(managements)
}

func seedManager() map[string]models.Management {
	shijiroManagement := services.AddManagement(models.Management{
		Name: "Shijiro",
	})

	totoroManagement := services.AddManagement(models.Management{
		Name: "Tororo",
	})

	return map[string]models.Management{
		shijiroManagement.Name: shijiroManagement,
		totoroManagement.Name:  totoroManagement,
	}
}

func seedAccount(managements map[string]models.Management) {
	shijiroManagement := managements["Shijiro"]

	services.AddAccount(services.AddAccountDTO{
		Name:         "Lulo Bank",
		Type:         constant.AccountTypeBank,
		ManagementID: shijiroManagement.ID,
	})

	services.AddAccount(services.AddAccountDTO{
		Name:         "Nu Bank",
		Type:         constant.AccountTypeBank,
		ManagementID: shijiroManagement.ID,
	})

	services.AddAccount(services.AddAccountDTO{
		Name:         "Nequi",
		Type:         constant.AccountTypeBank,
		ManagementID: shijiroManagement.ID,
	})
}
