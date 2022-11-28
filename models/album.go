package models

import "time"

type Album struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	ReleaseDate time.Time
	Duration    string
	Stars       uint8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
