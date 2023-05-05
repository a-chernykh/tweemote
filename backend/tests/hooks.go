package tests

import (
	"fmt"
	"os"

	"bitbucket.org/andreychernih/tweemote/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func BeforeSuite() {
	//log.Println("BeforeSuite()")
	prepareDb()
}

func AfterSuite() {
	//log.Println("BeforeSuite()")
}

func Before() {
	//log.Println("Before()")
	prepareDb()
	tables := []string{"users", "access", "refresh", "twitter_accounts", "twitter_applications"}
	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)).Error; err != nil {
			panic(err)
		}
	}
}

func After() {
	//log.Println("After()")
}

func prepareDb() {
	os.Setenv("DATABASE_URL", "postgres://postgres:w0ntt3lly0u@localhost/tweemote_test?sslmode=disable")
	if db == nil {
		db = models.Connect()
		//db.LogMode(true)
	}
}
