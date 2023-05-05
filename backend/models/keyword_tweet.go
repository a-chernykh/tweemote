package models

import "time"

type KeywordTweet struct {
	ID        uint `gorm:"primary_key"`
	KeywordID uint
	TweetID   uint

	CreatedAt time.Time
	UpdatedAt time.Time
}
