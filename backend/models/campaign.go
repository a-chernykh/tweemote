package models

import "time"

type Campaign struct {
	ID               uint `gorm:"primary_key"`
	Name             string
	TwitterAccountID uint

	TwitterAccount TwitterAccount
	Keywords       []Keyword
	Stats          []Stat

	CreatedAt time.Time
	UpdatedAt time.Time
}
