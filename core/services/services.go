package services

import "gorm.io/gorm"

type Dependencies struct {
	DB *gorm.DB
}

var d Dependencies

func SetServicesDependencies(dep Dependencies) {
	d = dep
}
