package services

import (
	"bitbucket.org/andreychernih/tweemote/forms"
	"bitbucket.org/andreychernih/tweemote/models"
	"github.com/asaskevich/govalidator"
)

func AddKeyword(campaignId uint, f forms.Keyword) (*models.Keyword, error) {
	_, err := govalidator.ValidateStruct(f)
	if err != nil {
		return nil, err
	}

	db := models.Connect()
	k := models.Keyword{CampaignID: campaignId, Keyword: f.Keyword}
	if err := db.Create(&k).Error; err != nil {
		return &k, err
	}

	return &k, nil
}
