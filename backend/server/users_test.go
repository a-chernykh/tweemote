package server

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/tests"
)

var validRequest = `{ "email": "andrey@example.com", "password": "secretpass", "password_confirmation": "secretpass" }`
var invalidRequest = `{ "email": "andrey@example.com", "password": "secretpass", "password_confirmation": "anotherpass" }`

func sendCreateUserRequest(json string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", "/api/v1/users", strings.NewReader(json))
	w := httptest.NewRecorder()
	CreateUserHandler(w, req)

	return w
}

func TestMain(m *testing.M) {
	tests.BeforeSuite()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateUserCreatesAUser(t *testing.T) {
	tests.Before()
	defer tests.After()

	w := sendCreateUserRequest(validRequest)

	tests.AssertResponseCode(w, t, 201)

	var user models.User
	db := models.Connect()

	if err := db.Where("email = ?", "andrey@example.com").First(&user).Error; err != nil {
		t.Error(err)
	}
}

func TestCreateUserValidation(t *testing.T) {
	tests.Before()
	defer tests.After()

	w := sendCreateUserRequest(invalidRequest)

	tests.AssertResponseCode(w, t, 422)
}

func TestGetCurrentUserAuthenticated(t *testing.T) {
	tests.Before()
	defer tests.After()

	req := httptest.NewRequest("GET", "/api/v1/users/me", nil)

	user := tests.CreateUser("user@example.com")
	req = tests.AddUser(req, user)

	w := httptest.NewRecorder()
	GetCurrentUserHandler(w, req)

	tests.AssertResponseCode(w, t, 200)

	var r CurrentUserResponse
	err := json.Unmarshal(w.Body.Bytes(), &r)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if r.Email != user.Email {
		t.Errorf("Expected e-amil to be %s. Got %s.", user.Email, r.Email)
	}
}
