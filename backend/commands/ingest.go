package commands

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/golang/glog"

	"bitbucket.org/andreychernih/tweemote/errors"
	"bitbucket.org/andreychernih/tweemote/mb"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/twitter"
	"bitbucket.org/andreychernih/tweemote/worker"
)

var threads int

type IngestCommand struct {
	Meta
}

func (cmd IngestCommand) Help() string {
	return "Uses Twitter Stream API to subscribe for keywords and push them to the queue for further processing"
}

func (cmd IngestCommand) Synopsis() string {
	return "(Worker) hooks into twitter stream and search for tweets"
}

func (cmd IngestCommand) Run(args []string) int {
	flags := flag.NewFlagSet("persist", flag.ExitOnError)
	flags.IntVar(&threads, "threads", 4, "Number of goroutines")
	flags.Parse(os.Args[2:])

	dFunc := func(jobsChan chan *worker.Job, quitChan chan int) {
		for i := 0; i < threads; i++ {
			j := worker.NewJob(fmt.Sprintf("keywords-%d", i), nil)
			jobsChan <- j
		}
	}

	wFunc := func(w *worker.Worker, j *worker.Job, quitChan chan int) {
		ingest(w, j, quitChan)
	}

	// Start worker pool
	sPool := worker.NewWorkerPool("ingest", threads, dFunc, wFunc)
	sPool.Start()
	sPool.Wait()

	return 0
}

func ingest(w *worker.Worker, j *worker.Job, quitChan chan int) {
	tick := time.Tick(10 * time.Second)

	for {
		select {
		case <-quitChan:
			return
		case <-tick:
			keywords, err := models.GetAndLockKeywords(w.GetID(), 400)
			if err != nil {
				errors.Check(err)
			}

			if len(keywords) > 0 {
				ta, err := models.GetAndLockTwitterAccount(w.GetID())
				if err != nil {
					errors.Check(err)
				}

				if ta != nil {
					app := ta.TwitterApplication

					campaignIdsMap := make(map[uint]struct{})
					track := make([]string, 0)
					for _, k := range keywords {
						track = append(track, k.Keyword)
						campaignIdsMap[k.CampaignID] = struct{}{}
					}
					campaignIds := make([]uint, 0, len(campaignIdsMap))
					for cid := range campaignIdsMap {
						campaignIds = append(campaignIds, cid)
					}

					glog.Infof("%s: Keywords %v", w.GetName(), track)
					glog.Infof("%s: Campaigns %v", w.GetName(), campaignIds)

					twitterApi := twitter.NewTwitterAPI(app.ConsumerKey, app.ConsumerSecret, ta.AccessToken, ta.AccessTokenSecret)
					stream := twitterApi.KeywordsStream(track)

					broker, err := mb.NewConnection(1)
					errors.Check(err)
					defer broker.Disconnect()

					nils := 0

					for {
						select {
						case <-quitChan:
							stream.Stop()
							return
						case o := <-stream.C:
							if o != nil {
								nils = 0

								switch v := o.(type) {
								case anaconda.Tweet:
									glog.Infof("Tweet: %s %s", v.User.ScreenName, v.Text)

									m := buildMessage(&v)
									m.CampaignIds = campaignIds
									broker.Publish(mb.TweetsQueueName, mb.NewJsonMessage(m))

								default:
									glog.Errorf("Received non-tweet: %v", v)
								}
							} else {
								nils = nils + 1
								if nils > 10 {
									panic("Received nil from stream")
								}
							}
						}
					}
				}
			}
		}
	}
}

func buildMessage(t *anaconda.Tweet) *mb.TweetMessage {
	var retweetedId *string
	var retweetedUserId *string

	if t.RetweetedStatus != nil {
		rId := strconv.FormatInt(t.RetweetedStatus.Id, 10)
		rUid := strconv.FormatInt(t.RetweetedStatus.User.Id, 10)
		retweetedId = &rId
		retweetedUserId = &rUid
	}

	message := mb.TweetMessage{}
	message.TweetID = strconv.FormatInt(t.Id, 10)
	message.UserID = strconv.FormatInt(t.User.Id, 10)
	message.RetweetedID = retweetedId
	message.RetweetedUserID = retweetedUserId
	message.TweetText = t.Text
	message.LikesCount = uint(t.FavoriteCount)
	message.RetweetCount = uint(t.RetweetCount)

	return &message
}
