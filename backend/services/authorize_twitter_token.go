package services

import (
	"errors"

	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/twitter"
	"github.com/garyburd/go-oauth/oauth"
)

func AuthorizeTwitterToken(p twitter.TwitterClientProvider, appId string, oauthToken string, oauthVerifier string) (*models.TwitterAccount, error) {
	db := models.Connect()

	var app models.TwitterApplication
	if err := db.Where("id = ?", appId).First(&app).Error; err != nil {
		return nil, err
	}

	var toks []models.TwitterRequestToken
	if err := db.Model(&app).Where("oauth_token = ?", oauthToken).Related(&toks).Error; err != nil {
		return nil, err
	}

	if len(toks) == 0 {
		return nil, errors.New("Token not found")
	}

	tok := toks[0]

	tc := p.GetClientForApp(app)
	at, err := tc.CreateAccessToken(oauth.Credentials{Token: tok.OauthToken, Secret: tok.OauthSecret}, oauthVerifier)
	if err != nil {
		return nil, err
	}

	api := tc.API(at)
	twUser, err := api.GetSelf()
	if err != nil {
		return nil, err
	}

	var account models.TwitterAccount
	if nf := db.Where("twitter_user_id = ?", twUser.UserID).First(&account).RecordNotFound(); nf {
		account = models.TwitterAccount{
			UserID:               tok.UserID,
			TwitterApplicationID: app.ID,
			TwitterUserID:        twUser.UserID,
			TwitterUsername:      twUser.Username,
			AccessToken:          at.Token,
			AccessTokenSecret:    at.Secret,
			State:                "active",
		}
		if err := db.Create(&account).Error; err != nil {
			return nil, err
		}

		campaign := models.Campaign{TwitterAccount: account, Name: "Default campaign"}
		if err := db.Create(&campaign).Error; err != nil {
			return nil, err
		}
	} else {
		if account.UserID != tok.UserID {
			return nil, errors.New("Twitter account is already linked to another user")
		} else {
			account.TwitterApplicationID = app.ID
			account.TwitterUserID = twUser.UserID
			account.TwitterUsername = twUser.Username
			account.AccessToken = at.Token
			account.AccessTokenSecret = at.Secret
			account.State = "active"
			if err := db.Save(&account).Error; err != nil {
				return nil, err
			}
		}
	}

	return &account, nil
}
