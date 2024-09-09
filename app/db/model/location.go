package model

import "time"

type Location struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
