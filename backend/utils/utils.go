package utils

import (
	"fmt"
	"regexp"
	"time"

	"github.com/golang/glog"
)

func FormatTweet(tweet string) string {
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(tweet, " ")
}

func Retry(attempts int, sleep time.Duration, callback func() error) error {
	var err error

	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return nil
		}
		if i > attempts {
			break
		}

		glog.Warningf("Got error: %s. Re-trying in %s (%d / %d)", err, sleep, i, attempts)
		time.Sleep(sleep)
	}
	return fmt.Errorf("Giving up after %d attempts. Last error was: %s", attempts, err)
}
