package models

import "time"

type TwitterAccount struct {
	ID                   uint `gorm:"primary_key"`
	TwitterApplicationID uint
	UserID               uint
	TwitterUserID        string
	TwitterUsername      string
	AccessToken          string
	AccessTokenSecret    string
	State                string

	TwitterApplication TwitterApplication
	Campaigns          []Campaign

	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllActiveTwitterAccounts(accounts *[]TwitterAccount) {
	db := Connect()
	db.Where("state = 'active'").Find(&accounts)
}

func GetAndLockTwitterAccount(workerID uint) (*TwitterAccount, error) {
	db := Connect()
	if err := db.Exec("UPDATE twitter_accounts SET worker_id = ? WHERE worker_id IS NULL AND state = 'active'", workerID).Error; err != nil {
		return nil, err
	}

	var ta TwitterAccount
	res := db.Preload("TwitterApplication").Where("worker_id = ?", workerID).First(&ta)
	if res.RecordNotFound() {
		return nil, nil
	}

	if err := res.Error; err != nil {
		return nil, err
	}

	return &ta, nil
}

func (ta *TwitterAccount) LastImpressionTime() *time.Time {
	db := Connect()

	var times []time.Time
	db.Model(&Impression{}).Where("actor_twitter_user_id = ?", ta.TwitterUserID).Order("created_at DESC").Limit(1).Pluck("created_at", &times)

	if len(times) == 0 {
		return nil
	} else {
		return &times[0]
	}
}
