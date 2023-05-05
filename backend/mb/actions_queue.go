package mb

import (
	"fmt"
)

var ActionsQueueName = fmt.Sprintf("%s.tweemote.actions", QueueVersion)

const (
	ACTION_LIKE = "like"
)

type ActionMessage struct {
	Action      string `json:"action"`
	SubjectID   uint   `json:"subject_id"`
	SubjectType string `json:"subject_type"`
	KeywordID   uint   `json:"keyword_id"`
}
