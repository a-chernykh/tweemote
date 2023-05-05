package twitter

import (
	"net/http"

	"bitbucket.org/andreychernih/tweemote/models"
	"github.com/garyburd/go-oauth/oauth"
)

type Client struct {
	oauthClient oauth.Client
	app         models.TwitterApplication
}

type TwitterClient interface {
	GetApp() models.TwitterApplication
	GetAuthURL(callback string) (string, *oauth.Credentials, error)
	CreateAccessToken(oauth.Credentials, string) (*oauth.Credentials, error)

	API(*oauth.Credentials) TwitterAPI
}

type TwitterClientProvider interface {
	GetClient() TwitterClient
	GetClientForApp(models.TwitterApplication) TwitterClient
}

type DefaultTwitterClientProvider struct{}

func (p DefaultTwitterClientProvider) GetClient() TwitterClient {
	db := models.Connect()

	var app models.TwitterApplication
	if err := db.Order("RANDOM()").Limit(1).First(&app).Error; err != nil {
		return nil
	}

	return NewClient(&app)
}
func (p DefaultTwitterClientProvider) GetClientForApp(app models.TwitterApplication) TwitterClient {
	return NewClient(&app)
}

func NewClient(app *models.TwitterApplication) *Client {
	var oauthClient = oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authenticate",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
	}

	oauthClient.Credentials.Token = app.ConsumerKey
	oauthClient.Credentials.Secret = app.ConsumerSecret

	return &Client{oauthClient: oauthClient, app: *app}
}

func (c Client) GetApp() models.TwitterApplication {
	return c.app
}

// TODO This will be buggy in multi-threaded env because of
// https://github.com/ChimeraCoder/anaconda/issues/101
func (c Client) GetAuthURL(callback string) (string, *oauth.Credentials, error) {
	tempCred, err := c.oauthClient.RequestTemporaryCredentials(http.DefaultClient, callback, nil)
	if err != nil {
		return "", nil, err
	}
	return c.oauthClient.AuthorizationURL(tempCred, nil), tempCred, nil
}

func (c Client) CreateAccessToken(tc oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	tok, _, err := c.oauthClient.RequestToken(nil, &tc, oauthVerifier)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func (c Client) API(creds *oauth.Credentials) TwitterAPI {
	// TODO this is not thread safe https://github.com/ChimeraCoder/anaconda/issues/101
	return NewTwitterAPI(c.app.ConsumerKey, c.app.ConsumerSecret,
		creds.Token, creds.Secret)
}
