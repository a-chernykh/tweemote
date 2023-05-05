package models

import "time"

type Tweet struct {
	ID              uint `gorm:"primary_key"`
	TweetID         string
	UserID          string
	TweetText       string
	RetweetedID     *string
	RetweetedUserID *string
	LikesCount      uint
	RetweetCount    uint

	Keywords []Keyword `gorm:"many2many:keyword_tweets;"`

	CreatedAt time.Time
}
