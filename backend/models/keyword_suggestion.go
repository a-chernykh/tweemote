package models

type KeywordSuggestion struct {
	ID                   uint `gorm:"primary_key"`
	CampaignID           uint
	Keyword              string
	PotentialImpressions uint
}
