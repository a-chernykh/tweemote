package commands

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"

	"bitbucket.org/andreychernih/tweemote/errors"
	"bitbucket.org/andreychernih/tweemote/mb"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/worker"
)

const ImpressDelay = 60000 * time.Millisecond
const MaximumUserQueueBacklog = 100

var impressStmt *sql.Stmt
var isEngagedStmt *sql.Stmt
var impressThreads int
var fs *flag.FlagSet

type ImpressCommand struct {
	Meta
}

func (cmd ImpressCommand) Help() string {
	return cmd.Synopsis()
}

func (cmd ImpressCommand) Synopsis() string {
	return "worker which decides what to do with the subject (like / follow / retweet)"
}

func (cmd ImpressCommand) Run(args []string) int {
	fs = flag.NewFlagSet("impress", flag.ExitOnError)
	fs.IntVar(&impressThreads, "threads", 1, "Number of goroutines to use which will be saving tweets")
	fs.Parse(os.Args[2:])

	broker, err := mb.NewConnection(persistThreads)
	errors.Check(err)
	defer broker.Disconnect()

	dFunc, err := CreateDispatchFunc(broker, mb.ImpressQueueName, persistThreads)
	errors.Check(err)
	wFunc := func(w *worker.Worker, j *worker.Job, quitChan chan int) {
		d := j.Object.(mb.MessageDelivery)
		m := mb.ImpressMessage{}
		d.UnmarshalTo(&m)
		switch m.SubjectType {
		case mb.SUBJECT_TYPE_TWEET:
			glog.V(2).Info("Got tweet")

			db := models.Connect()

			var t models.Tweet
			if err := db.First(&t, m.SubjectID).Error; err != nil {
				glog.Errorf("Unable to locate Tweet with ID %d: %s", m.SubjectID, err)
				d.Nack()
				return
			}

			if t.RetweetedID != nil {
				glog.V(2).Info("Skiping retweet")
				d.Ack()
				return
			}

			var kws []models.Keyword
			db.Preload("Campaign").Model(&t).Related(&kws, "Keywords")
			for _, kw := range kws {
				var ta models.TwitterAccount
				db.Model(&kw.Campaign).Related(&ta)
				if IsAlreadyEngaged(ta.TwitterUserID, t.UserID) {
					glog.V(2).Info("Already engaged with user, moving on")
					continue
				}

				queueName := actionQueueName(ta.ID)

				queueInfo, err := broker.QueueInspect(queueName)
				if err != nil {
					errors.Check(err)
				}

				if queueInfo.Messages < MaximumUserQueueBacklog {
					glog.V(2).Infof("Queueing message for account %d", t.ID)

					m := mb.ActionMessage{
						Action:      mb.ACTION_LIKE,
						SubjectID:   t.ID,
						SubjectType: mb.SUBJECT_TYPE_TWEET,
						KeywordID:   kw.ID,
					}
					broker.Publish(queueName, mb.NewJsonMessage(m))
				} else {
					glog.V(2).Infof("Dropping message for account %d (%d queued)", t.ID, queueInfo.Messages)
				}
			}

			d.Ack()
		}
	}

	wPool := worker.NewWorkerPool("impress", impressThreads, dFunc, wFunc)
	wPool.Start()
	wPool.Wait()

	return 0
}

func actionQueueName(taId uint) string {
	return fmt.Sprintf("%s-%d", mb.ActionsQueueName, taId)
}

func IsAlreadyEngaged(userId string, tweetUserId string) bool {
	db := models.Connect()
	var count int
	db.Model(&models.Impression{}).Where("actor_twitter_user_id = ? AND subject_twitter_user_id = ?", userId, tweetUserId).Count(&count)
	return (count > 0)
}
