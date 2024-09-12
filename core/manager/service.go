package manager

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
	Dependencies
}

type Dependencies struct {
	AccountServiceAdd schema.ActionService
}

func SetService(serviceSetting schema.ServiceSetting, dependencies Dependencies) *Service {
	service := &Service{
		DB:   serviceSetting.DB,
		Name: constant.ManagerService,
	}

	service.Services = append(service.Services, schema.Service{
		Name:   constant.ManagerServiceAdd,
		Action: service.Add,
	})

	service.AccountServiceAdd = dependencies.AccountServiceAdd

	return service
}

func (s *Service) Get(c echo.Context) {
}

func (s *Service) Add(c echo.Context) {

}

func (s *Service) Update() {

}

func (s *Service) Delete() {

}
