package models

type Operation struct {
	Model
	Name         string `gorm:"not null"`
	Description  string
	ManagementID ID `gorm:"not null"`
}
