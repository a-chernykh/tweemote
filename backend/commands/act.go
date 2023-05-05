package commands

import (
	"fmt"
	"time"

	"github.com/golang/glog"

	"bitbucket.org/andreychernih/tweemote/errors"
	"bitbucket.org/andreychernih/tweemote/mb"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/twitter"
	"bitbucket.org/andreychernih/tweemote/worker"
)

type ActCommand struct {
	Meta
}

func (cmd ActCommand) Run(args []string) int {
	broker, err := mb.NewConnection(persistThreads)
	errors.Check(err)
	defer broker.Disconnect()

	db := models.Connect()

	dFunc, err := CreateDispatchFunc(broker, mb.ActionsQueueName, persistThreads)
	errors.Check(err)
	wFunc := func(w *worker.Worker, j *worker.Job, quitChan chan int) {
		d := j.Object.(mb.MessageDelivery)
		m := mb.ActionMessage{}
		d.UnmarshalTo(&m)

		if m.SubjectType == mb.SUBJECT_TYPE_TWEET {
			var t models.Tweet
			if err := db.First(&t, m.SubjectID).Error; err != nil {
				errStr := fmt.Sprintf("Unable to locate Tweet with ID %d: %s", m.SubjectID, err)
				glog.Error(errStr)
				broker.Nack(d, errStr)
				return
			}

			var k models.Keyword
			if err := db.First(&k, m.KeywordID).Error; err != nil {
				glog.V(2).Infof("Unable to locate Keyword with ID %d: %s. Keyword was deleted?", m.KeywordID, err)
				d.Ack()
				return
			}

			var c models.Campaign
			db.Model(&k).Related(&c)
			var ta models.TwitterAccount
			db.Model(&c).Related(&ta)

			lit := ta.LastImpressionTime()
			if lit != nil && time.Now().Before(lit.Add(ProcessUserQueuesDelay).Add(-1*time.Second)) {
				glog.Info("Too early, pushing tweet back to the user queue")
				broker.Publish(actionQueueName(ta.ID), d.GetMessage())
				d.Ack()
				return
			}

			if IsAlreadyEngaged(ta.TwitterUserID, t.UserID) {
				glog.V(2).Info("Already engaged with user, moving on")
				d.Ack()
				return
			}

			api := twitter.NewTwitterAccountAPI(&ta)

			err, liked := api.Favorite(t.TweetID)
			if err != nil {
				errStr := fmt.Sprintf("Error liking tweet: %s", err)
				glog.Errorf(errStr)
				broker.Nack(d, errStr)
				return
			}

			if liked {
				imp := models.Impression{
					CampaignID:           c.ID,
					ActorTwitterUserID:   ta.TwitterUserID,
					SubjectTwitterUserID: t.UserID,
					Action:               "like",
					SubjectID:            fmt.Sprintf("%d", t.ID),
					SubjectType:          mb.SUBJECT_TYPE_TWEET,
				}
				if err := db.Create(&imp).Error; err != nil {
					errStr := fmt.Sprintf("Error saving impression: %s", err)
					glog.Error(errStr)
					broker.Nack(d, errStr)
					return
				}
			}
		} else {
			errStr := fmt.Sprintf("Unknown subject type: %s", m.SubjectType)
			glog.Errorf(errStr)
			broker.Nack(d, errStr)
			return
		}

		d.Ack()
	}

	wPool := worker.NewWorkerPool("impress", impressThreads, dFunc, wFunc)
	wPool.Start()
	wPool.Wait()

	return 0
}

func (cmd ActCommand) Synopsis() string {
	return "Act tweets using Twitter API"
}

func (cmd ActCommand) Help() string {
	return cmd.Synopsis()
}
