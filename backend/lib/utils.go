package lib

import (
	"time"

	"cloud.google.com/go/civil"
)

func GetCurrentDate() civil.Date {
	return civil.DateOf(time.Now())
}

func ToDate(t time.Time) string {
	return t.Format("2006-01-02")
}
