package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"bitbucket.org/andreychernih/tweemote/errors"
	"bitbucket.org/andreychernih/tweemote/mb"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/utils"
	"bitbucket.org/andreychernih/tweemote/worker"

	"github.com/golang/glog"
)

var persistThreads int
var verbosity int

type PersistCommand struct {
	Meta
}

func (cmd PersistCommand) Help() string {
	return "Saves tweets to the DB"
}

func (cmd PersistCommand) Synopsis() string {
	return "(Worker) saves tweets to the database"
}

func (cmd PersistCommand) Run(args []string) int {
	fs := flag.NewFlagSet("persist", flag.ExitOnError)
	fs.IntVar(&persistThreads, "threads", 1, "Number of goroutines to use which will be saving tweets")
	fs.Parse(os.Args[2:])

	broker, err := mb.NewConnection(persistThreads)
	errors.Check(err)
	defer broker.Disconnect()

	dFunc, err := CreateDispatchFunc(broker, mb.TweetsQueueName, persistThreads)
	errors.Check(err)

	wFunc := func(w *worker.Worker, j *worker.Job, quitChan chan int) {
		d := j.Object.(mb.MessageDelivery)
		//glog.Info("PERSIST: %s", d, d, string(d.GetMessage().Serialize()))

		m := mb.TweetMessage{}
		d.UnmarshalTo(&m)

		words := utils.TokenizeTweet(m.TweetText)
		glog.V(2).Infof("Extracted words: %v", words)

		wm := make(map[string]bool)
		for _, w := range words {
			wm[w] = true
		}

		for _, stopWord := range cmd.Meta.Config.StopWords {
			if _, ok := wm[stopWord]; ok {
				glog.Infof("Found stop word '%s', skipping", stopWord)
				d.Ack()
				return
			}
		}

		glog.V(2).Infof("Saving tweet: '%s'", utils.FormatTweet(m.TweetText))

		kws, err := models.GetCampaignKeywords(m.CampaignIds)
		if err != nil {
			glog.Errorf("Error processing tweet: %s", err)
		}

		t := models.Tweet{
			TweetID:         m.TweetID,
			UserID:          m.UserID,
			TweetText:       m.TweetText,
			RetweetedID:     m.RetweetedID,
			RetweetedUserID: m.RetweetedUserID,
			LikesCount:      m.LikesCount,
			RetweetCount:    m.RetweetCount,
		}

		db := models.Connect()
		tx := db.Begin()

		if err := tx.Create(&t).Error; err != nil {
			tx.Rollback()

			glog.Errorf("Error saving tweet: %s", err)
			err = d.Nack()
			errors.Check(err)
			return
		}

		any := false
		for k, ids := range kws {
			glog.V(3).Infof("Looking for a '%s' keyword", k)

			lk := strings.ToLower(k)
			lkws := strings.Fields(lk)

			found := true
			for _, lkw := range lkws {
				if _, ok := wm[lkw]; !ok {
					found = false
					break
				}
			}

			if found {
				glog.V(2).Infof("Found a '%s' keyword", k)

				for _, id := range ids {
					kt := models.KeywordTweet{
						KeywordID: id,
						TweetID:   t.ID,
					}

					if err := tx.Create(&kt).Error; err != nil {
						tx.Rollback()

						glog.Errorf("Error saving KeywordTweet: %s", err)
						err = d.Nack()
						errors.Check(err)
						return
					}

					any = true
				}
			}
		}

		if any {
			tx.Commit()
			d.Ack()

			m := mb.ImpressMessage{
				SubjectID:   fmt.Sprintf("%d", t.ID),
				SubjectType: mb.SUBJECT_TYPE_TWEET,
			}
			broker.Publish(mb.ImpressQueueName, mb.NewJsonMessage(m))
		} else {
			glog.V(2).Info("No keywords were found, moving on")
			tx.Rollback()
			d.Ack()
			return
		}
	}

	// Processing pool - goroutines which will be saving tweets to database
	wPool := worker.NewWorkerPool("persist", persistThreads, dFunc, wFunc)
	wPool.Start()
	wPool.Wait()

	return 0
}
