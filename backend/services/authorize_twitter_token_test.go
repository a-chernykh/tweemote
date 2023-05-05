package services

import (
	"strconv"
	"testing"

	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/tests"
)

func TestAuthorizeTwitterTokenCreatesCampaign(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := tests.CreateUser("user@example.com")
	app := tests.CreateTwitterApplication("app")
	tests.CreateTwitterRequestToken(app, user, "tok", "secret")

	_, err := AuthorizeTwitterToken(&tests.TestTwitterClientProvider{}, strconv.Itoa(int(app.ID)), "tok", "verifier")
	if err != nil {
		t.Error(err)
	}

	var accts []models.TwitterAccount
	db := models.Connect()
	if err := db.Model(&user).Related(&accts).Error; err != nil {
		t.Fatal(err)
	}

	var campaigns []models.Campaign
	if err := db.Model(&accts[0]).Related(&campaigns).Error; err != nil {
		t.Fatal(err)
	}

	if len(campaigns) != 1 {
		t.Error("Expected to create 1 default campaign")
	}
}
