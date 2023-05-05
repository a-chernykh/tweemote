package services

import (
	"testing"

	"bitbucket.org/andreychernih/tweemote/tests"
	"golang.org/x/crypto/bcrypt"
)

func TestRegistersUser(t *testing.T) {
	tests.Before()
	defer tests.After()

	user, err := RegisterUser("andrey@example.com", "secretpass")
	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.Fatal("Expected User not to be nil")
	}

	if user.Email != "andrey@example.com" {
		t.Error("Expected email to be andrey@example.com")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte("secretpass"))
	if err != nil {
		t.Errorf("Expected password to be cached with bcrypt. Error: %s", err)
	}
}
