package mb

import "fmt"

var ImpressQueueName = fmt.Sprintf("%s.tweemote.impress", QueueVersion)

const (
	SUBJECT_TYPE_TWEET        = "Tweet"
	SUBJECT_TYPE_TWITTER_USER = "TwitterUser"
)

type ImpressMessage struct {
	SubjectID   string `json:"subject_id"`
	SubjectType string `json:"subject_type"`
}
