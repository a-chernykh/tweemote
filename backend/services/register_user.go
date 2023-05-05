package services

import (
	"bitbucket.org/andreychernih/tweemote/models"
	"golang.org/x/crypto/bcrypt"
)

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func RegisterUser(email string, password string) (*models.User, error) {
	db := models.Connect()
	passwordHash, err := encryptPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{Email: email, PasswordHash: passwordHash}
	if err := db.Create(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil
}
