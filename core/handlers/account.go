package handlers

import (
	"money-manager/core/services"

	"github.com/labstack/echo/v4"
)

func GetAccountHandler(c echo.Context) error {
	services.GetAccount()
	return nil
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
