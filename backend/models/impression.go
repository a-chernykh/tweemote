package models

import "time"

const (
	IMPRESS_LIKE_TWEET  = "like_tweet"
	IMPRESS_FOLLOW_USER = "follow_user"
)

type Impression struct {
	ID uint `gorm:"primary_key"`

	CampaignID uint

	ActorTwitterUserID   string
	SubjectTwitterUserID string

	Action      string
	SubjectID   string
	SubjectType string

	CreatedAt time.Time
	UpdatedAt time.Time
}
