package models

import (
	"time"
)

type TwitterApplication struct {
	ID uint `gorm:"primary_key"`

	Name           string
	ConsumerKey    string
	ConsumerSecret string

	TwitterRequestTokens []TwitterRequestToken

	CreatedAt time.Time
	UpdatedAt time.Time
}
