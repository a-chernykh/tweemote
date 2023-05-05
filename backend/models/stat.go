package models

type Stat struct {
	ID uint `gorm:"primary_key"`

	CampaignID  uint
	Day         string
	Impressions uint
	Followers   uint
}
