package handlers

import (
	"money-manager/core/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetAccountResponseDTO struct {
	ID          uint
	Name        string
	Description string
}

func GetAccountHandler(c echo.Context) error {
	accountData := services.GetAccount(services.GetAccountFilter{
		ID: 1,
	})

	getAccountResponseDTO := GetAccountResponseDTO{
		ID:          accountData.ID,
		Name:        accountData.Name,
		Description: accountData.Description,
	}

	return c.JSON(http.StatusOK, getAccountResponseDTO)
}

func AddAccountHandler(c echo.Context) error {
	return nil
}

func UpdateAccountHandler() error {
	return nil
}

func DeleteAccountHandler() error {
	return nil
}
