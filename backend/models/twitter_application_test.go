package models_test

import (
	"testing"

	. "bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/tests"
)

func TestGetRandomTwitterApplicationNoApps(t *testing.T) {
	tests.Before()
	defer tests.After()

	_, err := GetRandomTwitterApplication()
	if err == nil {
		t.Error("Expected to return an error when no apps available")
	}
}

func TestGetRandomTwitterApplication(t *testing.T) {
	tests.Before()
	defer tests.After()

	tests.CreateTwitterApplication("app1")
	tests.CreateTwitterApplication("app2")

	app, err := GetRandomTwitterApplication()
	if err != nil {
		t.Fatalf("Expected not to return error, got: %s", err)
	}
	if app == nil {
		t.Fatal("Expected app not to be nil")
	}

	if app.Name != "app1" && app.Name != "app2" {
		t.Error("Expected to return either app1 or app2")
	}
}
