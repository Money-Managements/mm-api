package handlers

import (
	"money-manager/core/helpers"
	"money-manager/core/models"
	"money-manager/core/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetAccountResponseDTO struct {
	ID          models.ID
	Name        string
	Description string
}

func GetAccountHandler(c echo.Context) error {

	errorResponse := helpers.GetErrorResponse(helpers.GetErrorResponseArg{
		StatusCode: http.StatusBadRequest,
		ErrorType:  helpers.BodyPropInvalid,
	})

	accountData := services.GetAccount(services.GetAccountFilter{
		ID: 1,
	})

	getAccountResponseDTO := GetAccountResponseDTO{
		ID:          accountData.ID,
		Name:        accountData.Name,
		Description: accountData.Description,
	}

	if errorResponse.Code == http.StatusBadRequest {
		panic("something went wrong")
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
