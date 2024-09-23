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
		ManagementID: constant.UUID(management.ID),
	})

	services.AddOperation(services.AddOperationDTO{
		Name:         "Gym",
		ManagementID: constant.UUID(management.ID),
	})

	services.AddOperation(services.AddOperationDTO{
		Name:         "Home Support",
		ManagementID: constant.UUID(management.ID),
	})

	services.AddOperation(services.AddOperationDTO{
		Name:         "Internet",
		ManagementID: constant.UUID(management.ID),
	})
}

func seedAccount(management models.Management) map[string]models.Account {
	createdAccounts := make(map[string]models.Account)

	accountLuloBank := services.AddAccount(services.AddAccountDTO{
		Name:         "Lulo Bank",
		Type:         constant.AccountTypeBank,
		ManagementID: constant.UUID(management.ID),
	})
	createdAccounts["Lulo Bank"] = accountLuloBank

	accountNuBank := services.AddAccount(services.AddAccountDTO{
		Name:         "Nu Bank",
		Type:         constant.AccountTypeBank,
		ManagementID: constant.UUID(management.ID),
	})
	createdAccounts["Nu Bank"] = accountNuBank

	accountNequi := services.AddAccount(services.AddAccountDTO{
		Name:         "Nequi",
		Type:         constant.AccountTypeBank,
		ManagementID: constant.UUID(management.ID),
	})
	createdAccounts["Nequi"] = accountNequi

	accountSupport := services.AddAccount(services.AddAccountDTO{
		Name:         "Support",
		Type:         constant.AccountTypeAssign,
		ManagementID: constant.UUID(management.ID),
	})
	createdAccounts["Support"] = accountSupport

	accountMotorcycle := services.AddAccount(services.AddAccountDTO{
		Name:         "Motorcycle",
		Type:         constant.AccountTypeAssign,
		ManagementID: constant.UUID(management.ID),
	})
	createdAccounts["Motorcycle"] = accountMotorcycle

	accountServices := services.AddAccount(services.AddAccountDTO{
		Name:         "Services",
		Type:         constant.AccountTypeAssign,
		ManagementID: constant.UUID(management.ID),
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
		AccountID: constant.UUID(accountNequi.ID),
		AddTransactionBaseDTO: services.AddTransactionBaseDTO{
			Description:  "Income from job",
			ManagementID: accountNequi.ManagementID,
			Amount:       300000,
			LocationID:   constant.UUID(locationData.ID),
			OperationID:  constant.UUID(operationData.ID),
		},
	})

	services.AddTransactionSpend(services.AddTransactionSpendDTO{
		AccountID: constant.UUID(accountNequi.ID),
		AddTransactionBaseDTO: services.AddTransactionBaseDTO{
			Description:  "Dinner with friends",
			ManagementID: accountNequi.ManagementID,
			Amount:       6000,
			LocationID:   constant.UUID(locationData.ID),
			OperationID:  constant.UUID(operationData.ID),
		},
	})

	services.AddTransactionTransfer(services.AddTransactionTransferDTO{
		TargetAccountID: constant.UUID(accountLuloBank.ID),
		OriginAccountID: constant.UUID(accountNequi.ID),
		AddTransactionBaseDTO: services.AddTransactionBaseDTO{
			Description:  "Transfering",
			ManagementID: accountNequi.ManagementID,
			Amount:       600,
			LocationID:   constant.UUID(locationData.ID),
			OperationID:  constant.UUID(operationData.ID),
		},
	})
}
