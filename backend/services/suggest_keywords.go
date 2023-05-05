package services

import (
	"fmt"
	"reflect"
	"sort"

	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/utils"
	"github.com/golang/glog"
)

type tweetStat struct {
	tweet_text        string
	impressions_count uint
	followers_count   uint
}

type WordFrequency struct {
	Word      string
	Frequency uint
}

type WordFrequencyList []WordFrequency

func (wfl WordFrequencyList) Len() int           { return len(wfl) }
func (wfl WordFrequencyList) Less(i, j int) bool { return wfl[i].Frequency < wfl[j].Frequency }
func (wfl WordFrequencyList) Swap(i, j int)      { wfl[i], wfl[j] = wfl[j], wfl[i] }

func SuggestKeywords(campaign *models.Campaign, max int) error {
	db := models.Connect()

	var keywords []models.Keyword
	if err := db.Model(&campaign).Related(&keywords).Error; err != nil {
		return err
	}

	var existingWords []map[string]bool
	for _, kw := range keywords {
		wordsMap := make(map[string]bool)
		for _, word := range utils.TokenizeKeyword(kw.Keyword) {
			wordsMap[word] = true
		}
		existingWords = append(existingWords, wordsMap)
	}

	rows, err := db.Table("tweets").
		Joins("left join keyword_tweets on tweets.id = keyword_tweets.tweet_id left join keywords on keywords.id = keyword_tweets.keyword_id").
		Where("keywords.campaign_id = ? AND tweets.retweeted_id IS NULL", campaign.ID).
		Select("tweet_text, keywords.impressions_count AS impressions_count, keywords.followers_count AS followers_count").
		Rows()
	if err != nil {
		return err
	}

	//ngram1 := make(map[string]uint)
	ngram2 := make(map[string]uint)

	for rows.Next() {
		var tweet tweetStat
		err = db.ScanRows(rows, &tweet)
		if err != nil {
			return err
		}
		glog.Info(tweet)

		text := tweet.tweet_text

		words := utils.DeleteFrequentWords(utils.TokenizeTweet(utils.CleanTweet(text)))

		for i, _ := range words {
			// ngram2
			if i > 0 {
				ngramMap := make(map[string]bool)
				ngramMap[words[i-1]] = true
				ngramMap[words[i]] = true

				found := false
				for _, kwMap := range existingWords {
					if reflect.DeepEqual(kwMap, ngramMap) {
						found = true
						break
					}
				}

				if !found {
					ngram := fmt.Sprintf("%s %s", words[i-1], words[i])
					if _, ok := ngram2[ngram]; !ok {
						ngram2[ngram] = 1
					} else {
						ngram2[ngram] += 1
					}
				}
			}

			// ngram1
			// if _, ok := ngram1[word]; !ok {
			// 	ngram1[word] = 1
			// } else {
			// 	ngram1[word] += 1
			// }
		}
	}

	wfl := make(WordFrequencyList, len(ngram2))
	i := 0
	for w, f := range ngram2 {
		wfl[i] = WordFrequency{
			Word:      w,
			Frequency: f,
		}
		i++
	}
	sort.Sort(sort.Reverse(wfl))

	len := len(wfl)
	if len > max {
		len = max
	}

	tx := db.Begin()
	tx.Delete(models.KeywordSuggestion{}, "campaign_id = ?", campaign.ID)

	for _, ngram := range wfl[:len] {
		ks := models.KeywordSuggestion{
			CampaignID:           campaign.ID,
			Keyword:              ngram.Word,
			PotentialImpressions: ngram.Frequency,
		}
		if err := tx.Create(&ks).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
