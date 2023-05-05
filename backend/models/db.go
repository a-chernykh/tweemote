package models

import (
	"os"

	"bitbucket.org/andreychernih/tweemote/errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbInstance *gorm.DB

func establishConnection() (*gorm.DB, error) {
	var err error

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, nil
}

func Connect() *gorm.DB {
	if dbInstance == nil {
		var err error

		dbInstance, err = establishConnection()
		if err != nil {
			errors.Check(err)
		}
	}

	return dbInstance
}

func TryConnect() (*gorm.DB, error) {
	if dbInstance == nil {
		var err error

		dbInstance, err = establishConnection()
		if err != nil {
			return nil, err
		}
	}

	return dbInstance, nil
}
