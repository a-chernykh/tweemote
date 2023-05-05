package models

import (
	"time"
)

type TwitterRequestToken struct {
	ID uint `gorm:"primary_key"`

	TwitterApplicationID uint
	UserID               uint

	OauthToken  string
	OauthSecret string

	CreatedAt time.Time
	UpdatedAt time.Time
}
