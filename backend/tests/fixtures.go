package tests

import (
	"fmt"

	"bitbucket.org/andreychernih/tweemote/models"
)

func CreateUser(email string) *models.User {
	user := models.User{Email: email, PasswordHash: "$2a$10$VAzY/UAKXmdWYJ7G03SSN.LOvX.HSr8i6hsoSNZ7AGWiR4yzI5a2q"} // "secretpass"
	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}
	return &user
}

var twitterUserIdCounter = 1

func CreateTwitterAccount(u *models.User) *models.TwitterAccount {
	db := models.Connect()

	acc := models.TwitterAccount{UserID: u.ID,
		TwitterUserID:     fmt.Sprint(twitterUserIdCounter),
		TwitterUsername:   "username",
		AccessToken:       "tok",
		AccessTokenSecret: "sec"}

	if err := db.Create(&acc).Error; err != nil {
		panic(err)
	}

	twitterUserIdCounter++

	return &acc
}

func CreateTwitterApplication(name string) *models.TwitterApplication {
	db := models.Connect()
	app := models.TwitterApplication{Name: name,
		ConsumerKey:    fmt.Sprintf("consumer_key-%s", name),
		ConsumerSecret: fmt.Sprintf("consumer_secret-%s", name)}

	if err := db.Create(&app).Error; err != nil {
		panic(err)
	}

	return &app
}

func CreateTwitterRequestToken(app *models.TwitterApplication, user *models.User, token string, secret string) *models.TwitterRequestToken {
	db := models.Connect()
	t := models.TwitterRequestToken{TwitterApplicationID: app.ID,
		UserID:      user.ID,
		OauthToken:  token,
		OauthSecret: secret}

	if err := db.Create(&t).Error; err != nil {
		panic(err)
	}

	return &t
}
