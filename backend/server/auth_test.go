package server

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"bitbucket.org/andreychernih/tweemote/lib"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/services"
	"bitbucket.org/andreychernih/tweemote/tests"
)

type AccessResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

var userPassword = "password"

func createUser() *models.User {
	user, err := services.RegisterUser("andrey@example.com", userPassword)
	if err != nil {
		panic(err)
	}
	return user
}

func sendTokenRequest(params url.Values) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", "/api/v1/auth", strings.NewReader(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.SetBasicAuth(lib.WEB_CLIENT_ID, lib.WEB_CLIENT_SECRET)

	w := httptest.NewRecorder()
	CreateAccessToken(w, req)

	return w
}

func assertValidAccessResponse(t *testing.T, rr *httptest.ResponseRecorder, user *models.User) {
	var resp AccessResponse
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Error(err)
	}

	if resp.ExpiresIn != 3600 {
		t.Error("Token should expire in 3600 seconds")
	}

	if resp.AccessToken == "" {
		t.Errorf("access_token should be present, got: %s", resp.AccessToken)
	}

	if resp.RefreshToken == "" {
		t.Errorf("refresh_token should be present, got: %s", resp.RefreshToken)
	}

	s := lib.GetOsinServer()
	d, err := s.Storage.LoadAccess(resp.AccessToken)
	if err != nil {
		t.Fatalf("Expected to load access token, got error: %v", err)
	}
	if d == nil {
		t.Fatal("Token was not found")
	}

	if d.UserData.(string) != fmt.Sprintf("%d", user.ID) {
		t.Fatal("Wrong user. Expected: %d. Got: %d.", user.ID, d.UserData.(string))
	}
}

func TestCreateAccessToken(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := createUser()

	params := url.Values{}
	params.Set("grant_type", "password")
	params.Set("username", user.Email)
	params.Set("password", userPassword)
	params.Set("scope", "write")
	w := sendTokenRequest(params)

	if w.Code != 201 {
		t.Fatalf("Expected response code to be 201, got: %d. Body: %s", w.Code, w.Body)
	}

	assertValidAccessResponse(t, w, user)
}

func TestCreateAccessTokenInvalidUser(t *testing.T) {
	tests.Before()
	defer tests.After()

	params := url.Values{}
	params.Set("grant_type", "password")
	params.Set("username", "andrey")
	params.Set("password", "password")
	params.Set("scope", "write")
	w := sendTokenRequest(params)

	if w.Code != 403 {
		t.Fatalf("Expected response code to be 403, got: %d. Body: %s", w.Code, w.Body)
	}
}

func TestRefreshAccessToken(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := createUser()
	tok := tests.AddValidUserToken(lib.GetOsinServer(), lib.GetDefaultOsinClient(), user)

	params := url.Values{}
	params.Set("grant_type", "refresh_token")
	params.Set("refresh_token", tok.RefreshToken)
	w := sendTokenRequest(params)

	assertValidAccessResponse(t, w, user)
}

func TestInvalidRefreshToken(t *testing.T) {
	tests.Before()
	defer tests.After()

	params := url.Values{}
	params.Set("grant_type", "refresh_token")
	params.Set("refresh_token", "bla bla bla")
	w := sendTokenRequest(params)

	if w.Code != 403 {
		t.Fatalf("Expected response code to be 403, got: %d. Body: %s", w.Code, w.Body)
	}
}
