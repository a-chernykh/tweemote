package server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/andreychernih/tweemote/ctx"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/services"
	"bitbucket.org/andreychernih/tweemote/twitter"
)

type TwitterAccountsCollectionResponse struct {
	TwitterAccounts []TwitterAccountResponse `json:"twitter_accounts"`
}

type TwitterAccountResponse struct {
	ID              uint   `json:"id"`
	TwitterUserID   string `json:"twitter_user_id"`
	TwitterUsername string `json:"twitter_username"`
}

type LinkTwitterAccountResponse struct {
	RedirectUrl string `json:"redirect_url"`
}

func buildTwitterAccountResponse(a models.TwitterAccount) TwitterAccountResponse {
	return TwitterAccountResponse{
		ID:              a.ID,
		TwitterUserID:   a.TwitterUserID,
		TwitterUsername: a.TwitterUsername,
	}
}

func GetTwitterAccountsHandler(w http.ResponseWriter, r *http.Request) {
	user := ctx.UserFromContext(r.Context())

	accounts := make([]models.TwitterAccount, 0)
	db := models.Connect()
	db.Model(&user).Related(&accounts)

	var cr TwitterAccountsCollectionResponse
	rr := make([]TwitterAccountResponse, 0, len(accounts))
	for _, a := range accounts {
		rr = append(rr, buildTwitterAccountResponse(a))
	}
	cr.TwitterAccounts = rr

	renderJson(cr, w)
}

func CreateLinkTwitterAccountHandler(p twitter.TwitterClientProvider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := ctx.UserFromContext(r.Context())

		params := HTTPParamsGetter{r: r}
		cb := params.GetRequired("callback")
		if params.Error() != nil {
			renderError(params.Error(), w)
			return
		}

		authURL, err := services.RequestTwitterToken(cb, user, p)
		if err != nil {
			renderError(err, w)
			return
		}

		var re LinkTwitterAccountResponse
		re.RedirectUrl = authURL

		renderJson(re, w)
	}
}

func CreateLinkTwitterAccountCallbackHandler(p twitter.TwitterClientProvider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := HTTPParamsGetter{r: r}
		appId := params.GetRequired("twitter_application_id")
		oauthToken := params.GetRequired("oauth_token")
		oauthVerifier := params.GetRequired("oauth_verifier")
		if params.Error() != nil {
			renderError(params.Error(), w)
			return
		}

		ta, err := services.AuthorizeTwitterToken(p, appId, oauthToken, oauthVerifier)
		if err != nil {
			renderError(err, w)
			return
		}

		tar := buildTwitterAccountResponse(*ta)
		renderJson(tar, w)
	}
}

func UnlinkTwitterAccountHandler(w http.ResponseWriter, r *http.Request) {
	user := ctx.UserFromContext(r.Context())

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		renderError(errors.New("id must be present"), w)
		return
	}

	db := models.Connect()
	deleted := db.Where("user_id = ? AND id = ?", user.ID, id).Delete(models.TwitterAccount{}).RowsAffected
	if deleted == 0 {
		renderNotFound(w)
	} else {
		renderSuccess(w)
	}
}
