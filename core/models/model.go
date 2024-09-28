package models

import (
	"database/sql"
	"time"
)

type ID uint
type Amount float64
type DeletedAt sql.NullTime

type Model struct {
	ID        ID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
}
