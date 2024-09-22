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
	accounts := seedAccount(managements)
	seedTransaction(accounts)
}

func seedManager() map[string]models.Management {
	shijiroManagement := services.AddManagement(models.Management{
		Name: "Shijiro",
	})

	return map[string]models.Management{
		shijiroManagement.Name: shijiroManagement,
	}
}

func seedAccount(managements map[string]models.Management) map[string]models.Account {
	createdAccounts := make(map[string]models.Account)
	shijiroManagement := managements["Shijiro"]

	accountLuloBank := services.AddAccount(services.AddAccountDTO{
		Name:         "Lulo Bank",
		Type:         constant.AccountTypeBank,
		ManagementID: constant.UUID(shijiroManagement.ID),
	})
	createdAccounts["Lulo Bank"] = accountLuloBank

	accountNuBank := services.AddAccount(services.AddAccountDTO{
		Name:         "Nu Bank",
		Type:         constant.AccountTypeBank,
		ManagementID: constant.UUID(shijiroManagement.ID),
	})
	createdAccounts["Nu Bank"] = accountNuBank

	accountNequi := services.AddAccount(services.AddAccountDTO{
		Name:         "Nequi",
		Type:         constant.AccountTypeBank,
		ManagementID: constant.UUID(shijiroManagement.ID),
	})
	createdAccounts["Nequi"] = accountNequi

	accountSupport := services.AddAccount(services.AddAccountDTO{
		Name:         "Support",
		Type:         constant.AccountTypeAssign,
		ManagementID: constant.UUID(shijiroManagement.ID),
	})
	createdAccounts["Support"] = accountSupport

	accountMotorcycle := services.AddAccount(services.AddAccountDTO{
		Name:         "Motorcycle",
		Type:         constant.AccountTypeAssign,
		ManagementID: constant.UUID(shijiroManagement.ID),
	})
	createdAccounts["Motorcycle"] = accountMotorcycle

	accountServices := services.AddAccount(services.AddAccountDTO{
		Name:         "Services",
		Type:         constant.AccountTypeAssign,
		ManagementID: constant.UUID(shijiroManagement.ID),
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

	services.AddTransactionIncome(services.AddTransactionIncomeDTO{
		AccountID: constant.UUID(accountNequi.ID),
		AddTransactionBaseDTO: services.AddTransactionBaseDTO{
			Description:  "Income from job",
			ManagementID: accountNequi.ManagementID,
			Amount:       300000,
			LocationID:   constant.UUID(locationData.ID),
		},
	})

	services.AddTransactionSpend(services.AddTransactionSpendDTO{
		AccountID: constant.UUID(accountNequi.ID),
		AddTransactionBaseDTO: services.AddTransactionBaseDTO{
			Description:  "Dinner with friends",
			ManagementID: accountNequi.ManagementID,
			Amount:       6000,
			LocationID:   constant.UUID(locationData.ID),
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
		},
	})
}
