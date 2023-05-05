package services

import (
	"errors"
	"fmt"

	"github.com/garyburd/go-oauth/oauth"

	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/twitter"
)

type AuthURLProvider interface {
	GetAuthURL(string, *models.TwitterApplication) (string, *oauth.Credentials, error)
}

func RequestTwitterToken(cb string, user *models.User, p twitter.TwitterClientProvider) (string, error) {
	tc := p.GetClient()
	if tc == nil {
		return "", errors.New("Unable to find twitter client")
	}

	cbUrl := fmt.Sprintf("%s?twitter_application_id=%d", cb, tc.GetApp().ID)
	authUrl, tempCreds, err := tc.GetAuthURL(cbUrl)
	if err != nil {
		return "", err
	}

	db := models.Connect()
	rt := models.TwitterRequestToken{
		TwitterApplicationID: tc.GetApp().ID,
		UserID:               user.ID,
		OauthToken:           tempCreds.Token,
		OauthSecret:          tempCreds.Secret,
	}
	if err := db.Create(&rt).Error; err != nil {
		return "", err
	}

	return authUrl, nil
}
