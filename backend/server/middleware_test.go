package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/andreychernih/tweemote/ctx"
	"bitbucket.org/andreychernih/tweemote/lib"
	"bitbucket.org/andreychernih/tweemote/tests"
)

type TestErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func TestAuthMiddlewareNoAuth(t *testing.T) {
	tests.Before()
	defer tests.After()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/blah", nil)

	dummyHandler := func(http.ResponseWriter, *http.Request) {
		t.Fatal("I should not be called")
	}
	h := authMiddleware(dummyHandler)

	// Test without Authorization header
	h(w, r)

	tests.AssertResponseCode(w, t, 401)

	var er TestErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &er)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if er.Error != "invalid_request" {
		t.Errorf("Expected error to be invalid_request. Got %s.", er.Error)
	}
}

func TestAuthMiddlewareInvalid(t *testing.T) {
	tests.Before()
	defer tests.After()

	dummyHandler := func(http.ResponseWriter, *http.Request) {
		t.Fatal("I should not be called")
	}
	h := authMiddleware(dummyHandler)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/blah", nil)
	r.Header.Set("Authorization", "bearer tok")
	h(w, r)

	tests.AssertResponseCode(w, t, 401)
}

func TestAuthMiddlewareExpired(t *testing.T) {
	tests.Before()
	defer tests.After()

	dummyHandler := func(http.ResponseWriter, *http.Request) {
		t.Fatal("I should not be called")
	}
	h := authMiddleware(dummyHandler)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/blah", nil)

	user := tests.CreateUser("user@example.com")
	tok := tests.AddExpiredUserToken(lib.GetOsinServer(), lib.GetDefaultOsinClient(), user)
	r.Header.Set("Authorization", "bearer "+tok.AccessToken)

	h(w, r)

	tests.AssertResponseCode(w, t, 401)

	var er TestErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &er)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if er.Error != "invalid_grant" {
		t.Errorf("Expected error to be invalid_grant. Got %s.", er.Error)
	}
}

func TestAuthMiddlewareValidToken(t *testing.T) {
	tests.Before()
	defer tests.After()

	user := tests.CreateUser("user@example.com")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/blah", nil)

	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		u := ctx.UserFromContext(r.Context())
		w.Write([]byte(u.Email))
	}
	h := authMiddleware(dummyHandler)

	tok := tests.AddValidUserToken(lib.GetOsinServer(), lib.GetDefaultOsinClient(), user)
	r.Header.Set("Authorization", "bearer "+tok.AccessToken)

	w = httptest.NewRecorder()
	h(w, r)

	tests.AssertResponseCode(w, t, 200)
	if string(w.Body.Bytes()) != user.Email {
		t.Error("Response should be hi")
	}
}
