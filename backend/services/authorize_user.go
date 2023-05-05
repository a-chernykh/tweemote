package services

import (
	"errors"

	"bitbucket.org/andreychernih/tweemote/models"
	"golang.org/x/crypto/bcrypt"
)

func AuthorizeUser(email string, password string) (*models.User, error) {
	var user models.User
	var unauthorizedErr = errors.New("E-mail or password is invalid")

	db := models.Connect()
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, unauthorizedErr
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err == nil {
		return &user, nil
	} else {
		return nil, unauthorizedErr
	}
}
