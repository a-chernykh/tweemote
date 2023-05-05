package services

import (
	"testing"

	"bitbucket.org/andreychernih/tweemote/tests"
)

func TestAuthorizeUser(t *testing.T) {
	tests.Before()
	defer tests.After()

	var email = "andrey@example.com"
	var password = "secretpass"

	RegisterUser(email, password)

	user, err := AuthorizeUser(email, password)
	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.Fatal("Expected User not to be nil")
	}

	if user.Email != "andrey@example.com" {
		t.Error("Expected email to be andrey@example.com")
	}
}

func TestAuthorizeUserDoesNotExist(t *testing.T) {
	tests.Before()
	defer tests.After()

	var email = "andrey@example.com"
	var password = "secretpass"

	user, err := AuthorizeUser(email, password)
	if err == nil {
		t.Error("Expected to return error")
	}

	if user != nil {
		t.Errorf("Should not return user, got %v", user)
	}
}

func TestAuthorizeUserInvalidPassword(t *testing.T) {
	tests.Before()
	defer tests.After()

	var email = "andrey@example.com"
	var password = "secretpass"

	RegisterUser(email, password)

	user, err := AuthorizeUser(email, "invalidpassword")
	if err == nil {
		t.Error("Expected to return error")
	}

	if user != nil {
		t.Errorf("Should not return user, got %v", user)
	}
}
