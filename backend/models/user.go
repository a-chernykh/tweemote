package models

import "time"

type User struct {
	ID           uint `gorm:"primary_key"`
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	TwitterAccounts      []TwitterAccount
	TwitterRequestTokens []TwitterRequestToken
}
