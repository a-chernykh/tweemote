package lib

import (
	"database/sql"
	"log"

	"github.com/RangelReale/osin"
	"github.com/ory/osin-storage/storage/postgres"
)

var WEB_CLIENT_ID = "web"
var WEB_CLIENT_SECRET = "redacted"

func GetDefaultOsinClient() osin.Client {
	return &osin.DefaultClient{Id: WEB_CLIENT_ID,
		Secret:      WEB_CLIENT_SECRET,
		RedirectUri: "/auth",
		UserData:    nil}
}

func CreateOsinStorage(db *sql.DB) (osin.Storage, error) {
	store := postgres.New(db)
	err := store.CreateSchemas()
	if err != nil {
		return nil, err
	}

	err = store.RemoveClient(WEB_CLIENT_ID)
	if err != nil {
		return nil, err
	}

	err = store.CreateClient(GetDefaultOsinClient())
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}

	return store, nil
}
