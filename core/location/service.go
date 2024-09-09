package location

import (
	"money-manager/internal/constant"
	"money-manager/internal/schema"

	"gorm.io/gorm"
)

type Service struct {
	DB   *gorm.DB
	Name string
}

func SetLocationService(serviceSetting schema.ServiceSetting) *Service {
	return &Service{
		DB:   serviceSetting.DB,
		Name: constant.LocationService,
	}
}

func (ls *Service) Get() {

}

func (ls *Service) Add() {

}

func (ls *Service) Update() {

}

func (ls *Service) Delete() {

}
