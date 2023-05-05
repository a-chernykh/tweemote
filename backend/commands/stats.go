package commands

import (
	"flag"
	"fmt"
	"os"
	"time"

	"bitbucket.org/andreychernih/tweemote/errors"
	"bitbucket.org/andreychernih/tweemote/lib"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/services"
	"bitbucket.org/andreychernih/tweemote/twitter"
	"bitbucket.org/andreychernih/tweemote/worker"
	"cloud.google.com/go/civil"
	"github.com/golang/glog"
)

type StatsCommand struct {
	Meta
}

var periodSeconds int

func (cmd StatsCommand) Help() string {
	return cmd.Synopsis()
}

func (cmd StatsCommand) Synopsis() string {
	return "calculates stats (number of impressions and number of followers)"
}

func (cmd StatsCommand) Run(args []string) int {
	fs = flag.NewFlagSet("stats", flag.ExitOnError)
	fs.IntVar(&periodSeconds, "period-seconds", 3600, "Recalculate stats every seconds")
	fs.Parse(os.Args[2:])

	err := calculateKeywordStats()
	errors.Check(err)

	dFunc := func(jobsChan chan *worker.Job, quitChan chan int) {
		ticker := time.NewTicker(time.Duration(periodSeconds) * time.Second)

		recalcFunc := func() {
			glog.Info("Re-calculating stats for all Twitter accounts")
			var accounts []models.TwitterAccount
			models.GetAllActiveTwitterAccounts(&accounts)

			for _, account := range accounts {
				j := worker.NewJob(fmt.Sprintf("twitter-account-%s", account.TwitterUsername), account)
				jobsChan <- j
			}
		}
		recalcFunc()

		for {
			select {
			case <-ticker.C:
				recalcFunc()

			case <-quitChan:
				ticker.Stop()
				return
			}
		}
	}
	wFunc := func(w *worker.Worker, j *worker.Job, quitChan chan int) {
		ta := j.Object.(models.TwitterAccount)
		glog.Infof("[%s] Collecting stats for @%s", w.GetName(), ta.TwitterUsername)

		db := models.Connect()

		var campaigns []models.Campaign
		db.Model(&ta).Related(&campaigns)

		//	err := CollectFollowers(&ta)
		//	errors.Check(err)

		for _, c := range campaigns {
			glog.Infof("[%s] Processing campaign %d", w.GetName(), c.ID)

			err := calculateStats(&c)
			if err != nil {
				panic(errors.Wrap(err, fmt.Sprintf("Error calculating stats for campaign %d", c.ID)))
			}

			err = calculateKeywordSuggestions(&c)
			if err != nil {
				panic(errors.Wrap(err, fmt.Sprintf("Error calculating keyword suggestions for campaign %d", c.ID)))
			}
		}
	}

	wPool := worker.NewWorkerPool("stats", 4, dFunc, wFunc)
	wPool.Start()
	wPool.Wait()

	return 0
}

func CollectFollowers(ta *models.TwitterAccount) error {
	db := models.Connect()

	api := twitter.NewTwitterAccountAPI(ta)

	followers, err := api.GetFollowerIds(ta.TwitterUserID)
	if err != nil {
		return err
	}

	for _, followerId := range followers {
		// https://github.com/jinzhu/gorm/issues/1623
		if err := db.Exec("INSERT INTO followers (twitter_user_id, twitter_follower_id, followed_at) VALUES(?, ?, ?) ON CONFLICT (twitter_user_id, twitter_follower_id) DO NOTHING", ta.TwitterUserID, followerId, time.Now()).Error; err != nil {
			return err
		}
	}

	return nil
}

func calculateKeywordSuggestions(c *models.Campaign) error {
	services.SuggestKeywords(c, 10)
	return nil
}

func calculateStats(c *models.Campaign) error {
	db := models.Connect()

	var ta models.TwitterAccount
	db.Model(&c).Related(&ta)

	impressions := make(map[civil.Date]uint)
	followers := make(map[civil.Date]uint)

	startDate := lib.GetCurrentDate()
	endDate := lib.GetCurrentDate()

	tx := db.Begin()

	tx.Delete(models.Stat{}, "campaign_id = ?", c.ID)

	rows, err := tx.Table("impressions").Select("DATE(created_at)::VARCHAR as date, COUNT(*) AS impressions").Where("campaign_id = ?", c.ID).Group("DATE(created_at)").Rows()
	if err != nil {
		tx.Rollback()
		return err
	}

	for rows.Next() {
		var dateStr string
		var count uint
		rows.Scan(&dateStr, &count)

		date, err := civil.ParseDate(dateStr)
		if err != nil {
			tx.Rollback()
			return err
		}

		impressions[date] = count

		if date.Before(startDate) {
			startDate = date
		}
	}

	rows, err = tx.Table("followers").Select("DATE(followed_at)::VARCHAR as date, COUNT(*) AS cnt").Where("twitter_user_id = ?", ta.TwitterUserID).Group("DATE(followed_at)").Rows()
	if err != nil {
		tx.Rollback()
		return err
	}

	for rows.Next() {
		var dateStr string
		var count uint
		rows.Scan(&dateStr, &count)

		date, err := civil.ParseDate(dateStr)
		if err != nil {
			tx.Rollback()
			return err
		}

		followers[date] = count
	}

	currentDate := startDate
	for currentDate.Before(endDate.AddDays(1)) {
		stat := models.Stat{
			CampaignID:  c.ID,
			Day:         currentDate.String(),
			Impressions: impressions[currentDate],
			Followers:   followers[currentDate],
		}
		if err := tx.Create(&stat).Error; err != nil {
			tx.Rollback()
			return err
		}

		currentDate = currentDate.AddDays(1)
	}

	tx.Commit()

	return nil
}

func calculateKeywordStats() error {
	db := models.Connect()

	tx := db.Begin()
	if err := db.Exec("DELETE FROM keyword_stats").Error; err != nil {
		tx.Rollback()
		return err
	}

	sql := `
INSERT INTO keyword_stats (keyword_id, day, impressions_count, followers_count) (
	SELECT keywords.id, DATE(impressions.created_at) AS day, COUNT(impressions.id) AS impressions_count, COUNT(followers.id) AS followers_count
	FROM keywords
	LEFT JOIN keyword_tweets ON keywords.id = keyword_tweets.keyword_id
	LEFT JOIN impressions ON keyword_tweets.tweet_id = impressions.subject_id::integer
	LEFT JOIN followers on followers.twitter_follower_id = impressions.subject_twitter_user_id
	WHERE impressions.id IS NOT NULL
	GROUP BY keywords.id, DATE(impressions.created_at)
);
	`
	if err := db.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}

	sql = `
UPDATE keywords
SET impressions_count=subquery.impressions_count, followers_count=subquery.followers_count
FROM (
	SELECT keyword_id, SUM(impressions_count) AS impressions_count, SUM(followers_count) AS followers_count
	FROM keyword_stats
	GROUP BY keyword_stats.keyword_id
) AS subquery
WHERE keywords.id = subquery.keyword_id;
	`
	if err := db.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
