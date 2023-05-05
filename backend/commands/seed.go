package commands

import (
	"log"

	"bitbucket.org/andreychernih/tweemote/models"
)

type SeedCommand struct {
	Meta
}

func (cmd SeedCommand) Help() string {
	return ""
}

func (cmd SeedCommand) Synopsis() string {
	return "seed development database"
}

func (cmd SeedCommand) Run(args []string) int {
	log.Printf("Seeding Postgres database")

	db := models.Connect()

	app := models.TwitterApplication{Name: "qainstructor",
		ConsumerKey:    "redacted",
		ConsumerSecret: "redacted"}

	if err := db.Create(&app).Error; err != nil {
		panic(err)
	}

	log.Printf("Done")
	return 0
}
