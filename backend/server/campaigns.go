package server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/andreychernih/tweemote/ctx"
	"bitbucket.org/andreychernih/tweemote/models"
)

type CampaignsCollectionResponse struct {
	Campaigns []CampaignResponse `json:"campaigns"`
}

type CampaignResponse struct {
	ID               uint   `json:"id"`
	TwitterAccountID uint   `json:"twitter_account_id"`
	Name             string `json:"name"`
}

func buildCampaignResponse(c models.Campaign) CampaignResponse {
	return CampaignResponse{
		ID:               c.ID,
		TwitterAccountID: c.TwitterAccountID,
		Name:             c.Name,
	}
}

func getTwitterAccount(r *http.Request) (*models.TwitterAccount, error) {
	user := ctx.UserFromContext(r.Context())

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		return nil, errors.New("id must be present")
	}

	db := models.Connect()

	var accts []models.TwitterAccount
	if err := db.Model(&user).Where("id = ?", id).Related(&accts).Error; err != nil {
		return nil, err
	}

	if len(accts) == 0 {
		return nil, errors.New("Twitter Account not found")
	}

	return &accts[0], nil
}

func GetCampaignsHandler(w http.ResponseWriter, r *http.Request) {
	a, err := getTwitterAccount(r)
	if err != nil {
		renderError(err, w)
	}

	campaigns := make([]models.Campaign, 0)
	db := models.Connect()
	db.Model(&a).Related(&campaigns)

	var cr CampaignsCollectionResponse
	rr := make([]CampaignResponse, 0, len(campaigns))
	for _, a := range campaigns {
		rr = append(rr, buildCampaignResponse(a))
	}
	cr.Campaigns = rr

	renderJson(cr, w)
}
