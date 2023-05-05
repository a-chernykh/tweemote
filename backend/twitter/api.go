package twitter

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"bitbucket.org/andreychernih/tweemote/models"

	"github.com/ChimeraCoder/anaconda"
	"github.com/golang/glog"
)

const maxSearchResultsPerPage = 100

type TwitterAPI interface {
	GetSelf() (*User, error)
}

type API struct {
	twitterApi *anaconda.TwitterApi
}

func NewTwitterAPI(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *API {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	api.SetLogger(anaconda.BasicLogger)

	return &API{twitterApi: api}
}

func NewTwitterAccountAPI(ta *models.TwitterAccount) *API {
	db := models.Connect()
	var tapp models.TwitterApplication
	db.Model(&ta).Related(&tapp)
	return NewTwitterAPI(tapp.ConsumerKey, tapp.ConsumerSecret, ta.AccessToken, ta.AccessTokenSecret)
}

func (api *API) GetSelf() (*User, error) {
	u, err := api.twitterApi.GetSelf(url.Values{})
	if err != nil {
		return nil, err
	}
	return &User{UserID: u.IdStr, Username: u.ScreenName}, nil
}

func (api *API) TwitterUserByScreenName(screenName string) (*User, error) {
	u, err := api.twitterApi.GetUsersShow(screenName, url.Values{})
	if err != nil {
		return nil, err
	}
	return &User{UserID: u.IdStr}, nil
}

func (api *API) GetFollowerIds(userId string) ([]string, error) {
	v := url.Values{}
	v.Set("user_id", userId)

	followersChan := api.twitterApi.GetFollowersIdsAll(v)
	followers := make([]string, 0)

	for page := range followersChan {
		if page.Error != nil {
			return nil, page.Error
		}

		for _, followerId := range page.Ids {
			followers = append(followers, strconv.FormatInt(followerId, 10))
		}
	}

	return followers, nil
}

func (api *API) FollowUserId(userId string) error {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}

	v := url.Values{}
	_, err = api.twitterApi.FollowUserId(int64(userIdInt), v)

	return err
}

func (api *API) Favorite(tweetId string) (error, bool) {
	tweetIdInt, err := strconv.Atoi(tweetId)
	if err != nil {
		return err, false
	}

	_, err = api.twitterApi.Favorite(int64(tweetIdInt))

	if err == nil {
		return nil, true
	} else {
		switch err.(type) {
		case *anaconda.ApiError:
			twitterError := err.(*anaconda.ApiError)
			glog.Errorf("Got error from Twitter: %s", twitterError)

			if len(twitterError.Decoded.Errors) > 0 {
				errMessage := twitterError.Decoded.Error()
				errCode := twitterError.Decoded.Errors[0].Code

				switch {
				case errMessage == "You have already favorited this status.":
					glog.Info("Already liked, skipping")
					return nil, false
				case errMessage == "No status found with that ID.":
					glog.Info("Status was not found. Deleted? Moving on.")
					return nil, false
				case errCode == 142:
					glog.Info("Can't like tweet by protected user. Moving on.")
					return nil, false
				}
			}
		}

		return err, false
	}
}

func (api *API) KeywordsStream(keywords []string) *anaconda.Stream {
	v := url.Values{}
	v.Set("track", strings.Join(keywords, ","))
	glog.Infof("Subscribing to Twitter channel: %v", v)
	return api.twitterApi.PublicStreamFilter(v)
}

func (api *API) Search(query string, maxResults int) ([]anaconda.Tweet, error) {
	tweets := make([]anaconda.Tweet, 0, maxResults)

	v := url.Values{}
	if maxResults > maxSearchResultsPerPage {
		v.Set("count", "100")
	} else {
		v.Set("count", fmt.Sprintf("%d", maxResults))
	}

	result, err := api.twitterApi.GetSearch(query, v)
	if err != nil {
		return nil, err
	}

OuterLoop:
	for true {
		if len(result.Statuses) == 0 {
			break
		}

		for _, tweet := range result.Statuses {
			tweets = append(tweets, tweet)
			if len(tweets) >= maxResults {
				break OuterLoop
			}
		}

		result, err = result.GetNext(api.twitterApi)
		if err != nil {
			return nil, err
		}
	}

	return tweets, nil
}
