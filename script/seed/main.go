package main

import (
	"money-manager/app/db"
	"money-manager/core/models"
	"money-manager/core/services"
)

func main() {
	dataBase := db.ConnectDB()
	db.DropTables(dataBase)
	db.AutoMigrations(dataBase)

	services.SetServicesDependencies(services.Dependencies{
		DB: dataBase,
	})

	management := seedManagement()
	seedOperation(management)
	accounts := seedAccount(management)
	seedTransaction(accounts)
}

func seedManagement() models.Management {
	shijiroManagement := services.AddManagement(models.Management{
		Name: "Shijiro",
	})

	return shijiroManagement
}

func seedOperation(management models.Management) {
	services.AddOperation(services.AddOperationDTO{
		Name:         "Water Service",
		ManagementID: management.ID,
	})

	services.AddOperation(services.AddOperationDTO{
		Name:         "Gym",
		ManagementID: management.ID,
	})

	services.AddOperation(services.AddOperationDTO{
		Name:         "Home Support",
		ManagementID: management.ID,
	})

	services.AddOperation(services.AddOperationDTO{
		Name:         "Internet",
		ManagementID: management.ID,
	})
}

func seedAccount(management models.Management) map[string]models.Account {
	createdAccounts := make(map[string]models.Account)

	accountLuloBank := services.AddAccount(services.AddAccountDTO{
		Name:         "Lulo Bank",
		Type:         models.AccountTypeBank,
		ManagementID: management.ID,
	})
	createdAccounts["Lulo Bank"] = accountLuloBank

	accountNuBank := services.AddAccount(services.AddAccountDTO{
		Name:         "Nu Bank",
		Type:         models.AccountTypeBank,
		ManagementID: management.ID,
	})
	createdAccounts["Nu Bank"] = accountNuBank

	accountNequi := services.AddAccount(services.AddAccountDTO{
		Name:         "Nequi",
		Type:         models.AccountTypeBank,
		ManagementID: management.ID,
	})
	createdAccounts["Nequi"] = accountNequi

	accountSupport := services.AddAccount(services.AddAccountDTO{
		Name:         "Support",
		Type:         models.AccountTypeAssign,
		ManagementID: management.ID,
	})
	createdAccounts["Support"] = accountSupport

	accountMotorcycle := services.AddAccount(services.AddAccountDTO{
		Name:         "Motorcycle",
		Type:         models.AccountTypeAssign,
		ManagementID: management.ID,
	})
	createdAccounts["Motorcycle"] = accountMotorcycle

	accountServices := services.AddAccount(services.AddAccountDTO{
		Name:         "Services",
		Type:         models.AccountTypeAssign,
		ManagementID: management.ID,
	})
	createdAccounts["Services"] = accountServices

	return createdAccounts
}

func seedTransaction(accounts map[string]models.Account) {
	accountNequi := accounts["Nequi"]
	accountLuloBank := accounts["Lulo Bank"]

	locationData := services.GetLocation(services.GetLocationFilter{
		Name:         accountNequi.Name,
		ManagementID: accountNequi.ManagementID,
	})

	operationData := services.GetOperation(services.GetOperatonFilter{
		Name:         "Gym",
		ManagementID: accountNequi.ManagementID,
	})

	services.AddTransactionIncome(services.AddTransactionIncomeDTO{
		AccountID: accountNequi.ID,
		AddTransactionBaseDTO: services.AddTransactionBaseDTO{
			Description:  "Income from job",
			ManagementID: accountNequi.ManagementID,
			Amount:       300000,
			LocationID:   locationData.ID,
			OperationID:  operationData.ID,
		},
	})

	services.AddTransactionSpend(services.AddTransactionSpendDTO{
		AccountID: accountNequi.ID,
		AddTransactionBaseDTO: services.AddTransactionBaseDTO{
			Description:  "Dinner with friends",
			ManagementID: accountNequi.ManagementID,
			Amount:       6000,
			LocationID:   locationData.ID,
			OperationID:  operationData.ID,
		},
	})

	services.AddTransactionTransfer(services.AddTransactionTransferDTO{
		TargetAccountID: accountLuloBank.ID,
		OriginAccountID: accountNequi.ID,
		AddTransactionBaseDTO: services.AddTransactionBaseDTO{
			Description:  "Transfering",
			ManagementID: accountNequi.ManagementID,
			Amount:       600,
			LocationID:   locationData.ID,
			OperationID:  operationData.ID,
		},
	})
}
