package models

import "time"

type Follower struct {
	ID uint `gorm:"primary_key"`

	TwitterUserID     string
	TwitterFollowerID string

	FollowedAt time.Time
}
