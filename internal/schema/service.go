package schema

import (
	"gorm.io/gorm"
)

type ServiceSetting struct {
	DB *gorm.DB
}
