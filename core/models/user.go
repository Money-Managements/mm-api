package models

type User struct {
	Model
	Email    string `gorm:"not null;unique"`
	Password string
}
