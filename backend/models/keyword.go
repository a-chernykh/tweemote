package models

import (
	"strconv"
	"strings"
	"time"

	"bitbucket.org/andreychernih/tweemote/errors"
)

type Keyword struct {
	ID         uint `gorm:"primary_key"`
	Keyword    string
	CampaignID uint
	WorkerID   *uint

	Campaign Campaign
	Tweets   []Tweet `gorm:"many2many:keyword_tweets;"`

	ImpressionsCount uint
	FollowersCount   uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Returns up to [count] keywords and marks them as being processed by worker [worker]
func GetAndLockKeywords(workerID uint, count int) ([]Keyword, error) {
	db := Connect()
	db.Exec("UPDATE keywords SET worker_id = ? WHERE id IN (SELECT id FROM keywords WHERE worker_id IS NULL ORDER BY campaign_id LIMIT ? FOR UPDATE)", workerID, count)

	var keywords []Keyword
	if err := db.Where("worker_id = ?", workerID).Find(&keywords).Error; err != nil {
		return nil, err
	}

	return keywords, nil
}

func GetCampaignKeywords(campaignIds []uint) (map[string][]uint, error) {
	db, err := TryConnect()
	if err != nil {
		return nil, err
	}

	rows, err := db.Raw("SELECT keyword, string_agg(id::varchar, ',') AS ids FROM keywords WHERE campaign_id IN (?) GROUP BY keyword", campaignIds).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := make(map[string][]uint, 0)

	var keyword string
	var ids string

	for rows.Next() {
		rows.Scan(&keyword, &ids)

		idsArr := strings.Split(ids, ",")
		idsUint := make([]uint, 0, len(idsArr))
		for _, id := range idsArr {
			idUint, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				errors.Check(err)
			}
			idsUint = append(idsUint, uint(idUint))
		}
		r[keyword] = idsUint
	}

	return r, nil
}
