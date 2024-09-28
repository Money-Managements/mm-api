package models

const (
	LocationDefaultName = "Default"
)

type Location struct {
	Model
	Name         string `gorm:"not null"`
	ManagementID ID     `gorm:"not null"`
}
