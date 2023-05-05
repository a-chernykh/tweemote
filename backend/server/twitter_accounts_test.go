package server

import (
	"fmt"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"

	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/tests"
)

func TestGetTwitterAccountsHandler(t *testing.T) {
	tests.Before()
	defer tests.After()

	user1 := tests.CreateUser("user1@example.com")
	user2 := tests.CreateUser("user2@example.com")

	acc1 := tests.CreateTwitterAccount(user1)
	acc2 := tests.CreateTwitterAccount(user1)
	tests.CreateTwitterAccount(user2)

	req := httptest.NewRequest("GET", "/api/v1/twitter_accounts", nil)
	req = tests.AddUser(req, user1)

	w := httptest.NewRecorder()

	GetTwitterAccountsHandler(w, req)

	var rr TwitterAccountsCollectionResponse
	tests.UnmarshalResponse(w, t, &rr)

	rrIds := make([]string, 0, len(rr.TwitterAccounts))
	for _, r := range rr.TwitterAccounts {
		rrIds = append(rrIds, r.TwitterUserID)
	}
	sort.Strings(rrIds)

	user1Accts := []string{acc1.TwitterUserID, acc2.TwitterUserID}
	sort.Strings(user1Accts)

	if !reflect.DeepEqual(rrIds, user1Accts) {
		t.Error("Expected response to only include user1 accounts:", user1Accts, " Got:", rrIds)
	}
}

func TestLinkTwitterAccountHandler(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := tests.CreateUser("user@example.com")

	req := httptest.NewRequest("GET", "/api/v1/twitter_accounts/link?callback=http://example.com", nil)
	req = tests.AddUser(req, user)

	w := httptest.NewRecorder()

	h := CreateLinkTwitterAccountHandler(&tests.TestTwitterClientProvider{})
	h(w, req)
	tests.AssertResponseCode(w, t, 200)

	var r LinkTwitterAccountResponse
	tests.UnmarshalResponse(w, t, &r)

	if r.RedirectUrl != "http://example.com/bla" {
		t.Error("Invalid redirect URL")
	}

	db := models.Connect()
	var tokens []models.TwitterRequestToken
	if err := db.Model(&user).Related(&tokens).Error; err != nil {
		t.Fatal(err)
	}
	if len(tokens) == 0 {
		t.Fatal("Expected to create request token")
	}

	tok := tokens[0]
	if tok.UserID != user.ID {
		t.Fatal("Expected to link request token to user")
	}
}

func TestLinkTwitterAccountCallback(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := tests.CreateUser("user@example.com")
	app := tests.CreateTwitterApplication("app")
	tests.CreateTwitterRequestToken(app, user, "tok", "secret")

	url := fmt.Sprintf("/api/v1/twitter_accounts/callback?twitter_application_id=%d&oauth_token=%s&oauth_verifier=%s", app.ID, "tok", "verifier")
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()

	h := CreateLinkTwitterAccountCallbackHandler(&tests.TestTwitterClientProvider{})
	h(w, req)
	tests.AssertResponseCode(w, t, 200)

	var accts []models.TwitterAccount
	db := models.Connect()
	if err := db.Model(&user).Related(&accts).Error; err != nil {
		t.Fatal(err)
	}

	if len(accts) == 0 {
		t.Error("Expected to create new twitter account")
	}

	acc := accts[0]
	if acc.TwitterUserID != "123" {
		t.Errorf("Expected user id to be 123, got %s", acc.TwitterUserID)
	}
	if acc.TwitterUsername != "test" {
		t.Errorf("Expected user name to be test, got %s", acc.TwitterUsername)
	}
}

func TestUnlinkTwitterAccount(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := tests.CreateUser("user@example.com")
	acc1 := tests.CreateTwitterAccount(user)
	acc2 := tests.CreateTwitterAccount(user)

	url := fmt.Sprintf("/api/v1/twitter_accounts/%d", acc1.ID)
	req := httptest.NewRequest("DELETE", url, nil)
	tests.StubAuth(req, user)

	w := httptest.NewRecorder()
	r := createRouter()

	r.ServeHTTP(w, req)
	tests.AssertResponseCode(w, t, 204)

	db := models.Connect()

	if err := db.First(&acc1, acc1.ID).Error; err == nil {
		t.Error("Expected to delete account")
	}

	if err := db.First(&acc2, acc2.ID).Error; err != nil {
		t.Errorf("Expected not to delete another account: %s", err)
	}
}

func TestUnlinkTwitterAccountAnotherUser(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := tests.CreateUser("user@example.com")
	user2 := tests.CreateUser("user2@example.com")

	acc := tests.CreateTwitterAccount(user2)

	url := fmt.Sprintf("/api/v1/twitter_accounts/%d", acc.ID)
	req := httptest.NewRequest("DELETE", url, nil)
	tests.StubAuth(req, user)

	w := httptest.NewRecorder()
	r := createRouter()

	r.ServeHTTP(w, req)
	tests.AssertResponseCode(w, t, 404)

	db := models.Connect()

	if err := db.First(&acc, acc.ID).Error; err != nil {
		t.Error("Expected not to delete another user account")
	}
}

func TestUnlinkTwitterAccountNotFound(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := tests.CreateUser("user@example.com")

	url := "/api/v1/twitter_accounts/123456"
	req := httptest.NewRequest("DELETE", url, nil)
	tests.StubAuth(req, user)

	w := httptest.NewRecorder()
	r := createRouter()

	r.ServeHTTP(w, req)
	tests.AssertResponseCode(w, t, 404)
}
