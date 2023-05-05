package mb

import "fmt"

var TweetsQueueName = fmt.Sprintf("%s.tweemote.tweets", QueueVersion)

type TweetMessage struct {
	TweetID         string  `json:"tweet_id"`
	UserID          string  `json:"user_id"`
	RetweetedID     *string `json:"retweeted_id"`      // ID of the tweet which was retweeted (original)
	RetweetedUserID *string `json:"retweeted_user_id"` // ID of the user whose tweet was retweeted (original user)
	TweetText       string  `json:"tweet_text"`
	LikesCount      uint    `json:"likes_count"`
	RetweetCount    uint    `json:"retweet_count"`
	CampaignIds     []uint  `json:"campaign_ids"` // A hint which keywords to look for
}
