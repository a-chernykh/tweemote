package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func UnmarshalResponse(w *httptest.ResponseRecorder, t *testing.T, o interface{}) {
	err := json.Unmarshal(w.Body.Bytes(), o)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
}

func AssertResponseCode(w *httptest.ResponseRecorder, t *testing.T, code int) {
	if w.Code != code {
		t.Fatalf("Expected response code to be %d, got: %d. Body: %s", code, w.Code, w.Body)
	}
}
