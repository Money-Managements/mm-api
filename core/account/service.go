package account

import (
	"money-manager/internal/constant"
	"money-manager/internal/schema"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service struct {
	DB       *gorm.DB
	Name     string
	Services []schema.Service
}

func SetService(serviceSetting schema.ServiceSetting) *Service {
	service := &Service{
		DB:   serviceSetting.DB,
		Name: constant.AccountService,
	}

	service.Services = append(service.Services, schema.Service{
		Name:   constant.AcccountServiceGet,
		Action: service.Get,
	})

	return service
}

func (s *Service) Get(c echo.Context) {
	println("Working")
}

func (s *Service) Add(c echo.Context) {

}

func (s *Service) Update() {

}

func (s *Service) Delete() {

}
