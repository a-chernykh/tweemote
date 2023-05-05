package tests

import (
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/twitter"
	"github.com/garyburd/go-oauth/oauth"
)

type TestTwitterClientProvider struct{}
type TestTwitterClient struct {
	app models.TwitterApplication
}
type TestTwitterAPI struct{}

func (p *TestTwitterClientProvider) GetClient() twitter.TwitterClient {
	return &TestTwitterClient{app: *CreateTwitterApplication("app")}
}
func (p *TestTwitterClientProvider) GetClientForApp(app models.TwitterApplication) twitter.TwitterClient {
	return &TestTwitterClient{app: app}
}

func (c *TestTwitterClient) GetApp() models.TwitterApplication {
	return c.app
}
func (c *TestTwitterClient) GetAuthURL(callback string) (string, *oauth.Credentials, error) {
	return "http://example.com/bla", &oauth.Credentials{Token: "tok", Secret: "sec"}, nil
}
func (c *TestTwitterClient) CreateAccessToken(tc oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	return &oauth.Credentials{Token: "accesstok", Secret: "accesssec"}, nil
}
func (c *TestTwitterClient) API(creds *oauth.Credentials) twitter.TwitterAPI {
	return TestTwitterAPI{}
}

func (a TestTwitterAPI) GetSelf() (*twitter.User, error) {
	return &twitter.User{UserID: "123", Username: "test"}, nil
}
