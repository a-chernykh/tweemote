package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/andreychernih/tweemote/ctx"
	"bitbucket.org/andreychernih/tweemote/forms"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/services"
)

type KeywordsCollectionResponse struct {
	Keywords []KeywordResponse `json:"keywords"`
}

type KeywordResponse struct {
	ID               uint   `json:"id"`
	CampaignID       uint   `json:"campaign_id"`
	Keyword          string `json:"keyword"`
	ImpressionsCount uint   `json:"impressions_count"`
	FollowersCount   uint   `json:"followers_count"`
}

type SingleKeywordResponse struct {
	Keyword KeywordResponse `json:"keyword"`
}

func buildKeywordResponse(c *models.Keyword) KeywordResponse {
	return KeywordResponse{
		ID:               c.ID,
		CampaignID:       c.CampaignID,
		Keyword:          c.Keyword,
		ImpressionsCount: c.ImpressionsCount,
		FollowersCount:   c.FollowersCount,
	}
}

func getCampaign(r *http.Request) (*models.Campaign, error) {
	user := ctx.UserFromContext(r.Context())

	vars := mux.Vars(r)
	id := vars["campaignId"]
	if id == "" {
		return nil, errors.New("campaignId must be present")
	}

	db := models.Connect()

	var campaign models.Campaign
	var account models.TwitterAccount
	notFoundErr := errors.New("Unable to find campaign")

	if err := db.First(&campaign, id).Error; err != nil {
		return nil, notFoundErr
	}
	if err := db.Model(&campaign).Related(&account).Error; err != nil {
		return nil, notFoundErr
	}

	if account.UserID != user.ID {
		return nil, notFoundErr
	}

	return &campaign, nil
}

func GetKeywordsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := getCampaign(r)
	if err != nil {
		renderError(err, w)
	}

	keywords := make([]models.Keyword, 0)
	db := models.Connect()
	db.Order("keywords.keyword ASC").Model(&c).Related(&keywords)

	var kcr KeywordsCollectionResponse
	kr := make([]KeywordResponse, 0, len(keywords))
	for _, k := range keywords {
		kr = append(kr, buildKeywordResponse(&k))
	}
	kcr.Keywords = kr

	renderJson(kcr, w)
}

func CreateKeywordHandler(w http.ResponseWriter, r *http.Request) {
	c, err := getCampaign(r)
	if err != nil {
		renderError(err, w)
	}

	var kf forms.Keyword

	err = json.NewDecoder(r.Body).Decode(&kf)
	if err != nil {
		http.Error(w, errorJson(err), 400)
		return
	}

	k, err := services.AddKeyword(c.ID, kf)
	if err != nil {
		renderError(err, w)
		return
	}

	kr := buildKeywordResponse(k)
	skr := SingleKeywordResponse{Keyword: kr}
	renderJson(skr, w)
}

func DeleteKeywordHandler(w http.ResponseWriter, r *http.Request) {
	c, err := getCampaign(r)
	if err != nil {
		renderError(err, w)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var k models.Keyword

	db := models.Connect()
	if err := db.First(&k, id).Error; err != nil {
		renderError(err, w)
		return
	}

	if k.CampaignID != c.ID {
		renderError(errors.New("Keyword was not found"), w)
		return
	}

	err = services.DeleteKeyword(k.ID)
	if err != nil {
		renderError(err, w)
		return
	}

	renderSuccess(w)
}
