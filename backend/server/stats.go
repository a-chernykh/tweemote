package server

import (
	"net/http"

	"bitbucket.org/andreychernih/tweemote/models"
)

type StatsCollectionResponse struct {
	Stats []StatResponse `json:"stats"`
}

type StatResponse struct {
	ID          uint   `json:"id"`
	Day         string `json:"day"`
	Impressions uint   `json:"impressions"`
	Followers   uint   `json:"followers"`
}

func buildStatResponse(s models.Stat) StatResponse {
	return StatResponse{
		ID:          s.ID,
		Day:         s.Day,
		Impressions: s.Impressions,
		Followers:   s.Followers,
	}
}

func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := getCampaign(r)
	if err != nil {
		renderError(err, w)
	}

	db := models.Connect()

	var stats []models.Stat
	db.Model(&c).Related(&stats)

	var cr StatsCollectionResponse
	rr := make([]StatResponse, 0, len(stats))
	for _, a := range stats {
		rr = append(rr, buildStatResponse(a))
	}
	cr.Stats = rr

	renderJson(cr, w)
}
